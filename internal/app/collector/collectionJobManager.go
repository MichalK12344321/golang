package collector

import (
	"context"
	"fmt"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"

	"github.com/google/uuid"
)

type CollectionJobManager struct {
	logger             logging.Logger
	jobs               map[uuid.UUID]Job
	sshFactoryProvider ssh.SSHFactoryProvider
	ctx                context.Context
	sinkFactory        CollectionSinkFactory
	publisher          *Publisher
	repo               Repository
	runners            []CollectionRunner
}

func NewCollectionJobManager(
	ctx context.Context,
	logger logging.Logger,
	sshFactoryProvider ssh.SSHFactoryProvider,
	sinkFactory CollectionSinkFactory,
	publisher *Publisher,
	dataContext Repository,
	runners []CollectionRunner,
) *CollectionJobManager {
	return &CollectionJobManager{
		logger:             logger,
		jobs:               make(map[uuid.UUID]Job),
		sshFactoryProvider: sshFactoryProvider,
		ctx:                ctx,
		sinkFactory:        sinkFactory,
		publisher:          publisher,
		repo:               dataContext,
		runners:            runners,
	}
}

func (manager *CollectionJobManager) Start(payload broker.EventPayload) error {
	runCreateEvent := &events.RunCreateEvent{}
	parseError := payload.ParseBody(runCreateEvent, events.FromJson)
	if parseError != nil {
		payload.Acknowledge(true, false) // dead letter
		return parseError
	}
	key := runCreateEvent.RunId

	collectionCreateEvent := &events.CollectionCreateEvent{}
	eventDb, getError := manager.repo.GetEvent(collectionCreateEvent, runCreateEvent.CollectionId)
	if getError != nil {
		payload.Acknowledge(true, true) // collection not synced yet
		return getError
	}
	parseError = events.FromJson(
		collectionCreateEvent,
		eventDb.Raw.Bytes,
		eventDb.Type,
	)

	insertError := manager.repo.SaveEvent(runCreateEvent, runCreateEvent.RunId)
	if insertError != nil {
		payload.Acknowledge(true, true) // requeue
		return insertError
	}

	payload.Acknowledge(false, false) // event synced

	if parseError != nil {
		return parseError
	}

	_, hasKey := manager.jobs[key]
	if hasKey {
		return fmt.Errorf("run '%s' already added", key)
	}

	ctx, cancel := context.WithTimeout(manager.ctx, collectionCreateEvent.GetTimeout())

	collectionJob := &CollectionJob{
		RunId:     runCreateEvent.RunId,
		jobData:   (JobData)(NewCollectionJobData()),
		ctx:       ctx,
		ctxCancel: cancel,
		logger:    manager.logger,
		publisher: manager.publisher,
		event:     collectionCreateEvent,
		runners:   manager.runners,
	}

	collectionJob.sinks = manager.sinkFactory.New(ctx, runCreateEvent, collectionJob.jobData)
	manager.jobs[key] = collectionJob

	go func() {
		runErr := collectionJob.Run()
		manager.logger.Errors(runErr...)
		delete(manager.jobs, key)
	}()

	return nil
}

func (manager *CollectionJobManager) Terminate(id uuid.UUID) error {
	collectionJob, hasKey := manager.jobs[id]
	if !hasKey {
		return fmt.Errorf("job '%s' does not exists", id)
	}
	manager.logger.Warn("job '%s' is being terminated", id)
	return collectionJob.Terminate()
}

var _ JobManager = new(CollectionJobManager)
