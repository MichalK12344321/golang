package collector

import (
	"context"
	"fmt"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CollectionJob struct {
	RunId     uuid.UUID
	jobData   JobData
	ctx       context.Context
	ctxCancel context.CancelFunc
	sinks     []CollectionSink
	event     *events.CollectionCreateEvent
	logger    logging.Logger
	publisher *Publisher
	runners   []CollectionRunner
}

func (job *CollectionJob) Id() uuid.UUID {
	return job.RunId
}

func (job *CollectionJob) JobData() JobData {
	return job.jobData
}

func (collectionJob *CollectionJob) Terminate() error {
	err := collectionJob.publisher.PublishRunUpdateEvent(&events.RunUpdateEvent{
		RunCreateEvent: events.RunCreateEvent{
			RunId:        collectionJob.RunId,
			CollectionId: collectionJob.event.CollectionId,
		},
		Status: events.CollectionStatusTerminating,
		Error:  "",
	})
	collectionJob.ctxCancel()
	return err
}

func (collectionJob *CollectionJob) findRunner(target any) (CollectionRunner, error) {
	runnerType := reflect.TypeOf(target)
	runner, ok := lo.Find(
		collectionJob.runners,
		func(r CollectionRunner) bool {
			return reflect.TypeOf(r) == runnerType
		},
	)

	if !ok {
		return nil, fmt.Errorf("runner '%s' not registered", runnerType)
	}
	return runner, nil
}

func (collectionJob *CollectionJob) Runner() (CollectionRunner, CollectionRunnerOptions, error) {
	switch collectionJob.event.Type {
	case events.CollectionTypeSSH:
		runner, err := collectionJob.findRunner(&SSHRunner{})
		return runner, collectionJob.event.SSH, err
	case events.CollectionTypeGO:
		runner, err := collectionJob.findRunner(&GoRunner{})
		return runner, collectionJob.event.Go.Script, err
	default:
		return nil, nil, fmt.Errorf("runner for type '%s' does not exist", collectionJob.event.Type)
	}
}

func (collectionJob *CollectionJob) Run() []error {
	startedError := collectionJob.handleUpdate(events.CollectionStatusStarted)
	if startedError != nil {
		return []error{startedError}
	}

	for _, v := range collectionJob.sinks {
		err := v.Initialize()
		if err != nil {
			collectionJob.ctxCancel()
			return []error{err}
		}
	}

	runner, runnerOptions, err := collectionJob.Runner()
	if err != nil {
		return []error{err}
	}

	go runner.Run(collectionJob.ctx, collectionJob.jobData, runnerOptions)

	for {
		select {
		case data := <-collectionJob.jobData.StdoutChannel():
			lo.ForEach(collectionJob.sinks, func(sink CollectionSink, _ int) { sink.AppendStdout(data) })
		case data := <-collectionJob.jobData.StderrChannel():
			lo.ForEach(collectionJob.sinks, func(sink CollectionSink, _ int) { sink.AppendStderr(data) })
		case err := <-collectionJob.jobData.ErrorChannel():
			collectionJob.ctxCancel()
			updateError := collectionJob.handleUpdate(events.CollectionStatusFailure, err)
			return []error{err, updateError}
		case <-collectionJob.jobData.DoneChannel():
			errors := lo.Map(collectionJob.sinks, func(sink CollectionSink, _ int) error { return sink.Finalize() })
			updateError := collectionJob.handleUpdate(events.CollectionStatusSuccess)
			errors = append(errors, updateError)
			return errors
		case <-collectionJob.ctx.Done():
			updateError := collectionJob.handleUpdate(events.CollectionStatusTerminated, collectionJob.ctx.Err())
			return []error{updateError}
		}
	}
}

func (job *CollectionJob) handleUpdate(status events.CollectionStatus, errors ...error) error {
	updateEvent := &events.RunUpdateEvent{
		RunCreateEvent: events.RunCreateEvent{
			RunId:        job.RunId,
			CollectionId: job.event.CollectionId,
		},
		Status: status,
		Error:  "",
	}
	if util.AnyErrors(errors...) {
		notEmptyErrors := lo.Filter(errors, func(err error, _ int) bool { return err != nil })
		updateEvent.Error = strings.Join(lo.Map(notEmptyErrors, func(err error, i int) string { return err.Error() }), "\n")
	}
	return job.publisher.PublishRunUpdateEvent(updateEvent)
}

var _ Job = new(CollectionJob)
