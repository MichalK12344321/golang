package collector

import (
	"context"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
)

type Subscriber struct {
	subscriber broker.EventSubscriber
}

func NewSubscriber(
	ctx context.Context,
	connection broker.BrokerConnection,
	logger logging.Logger,
	config *config.Config,
	publisher *Publisher,
	factoryProvider broker.SubscriberFactoryProvider,
	repo Repository,
	manager JobManager,
) *Subscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:    broker.EXCHANGE,
		Queue:       QUEUE_NAME,
		RoutingKeys: []string{broker.COLLECTION_CREATE_KEY, broker.RUN_CREATE_KEY, broker.RUN_TERMINATE_KEY},
		Job: func(c <-chan broker.EventPayload) {
			SubscriberJob(ctx, c, logger, repo, manager)
		},
	}
	return &Subscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func SubscriberJob(
	ctx context.Context,
	eventChannel <-chan broker.EventPayload,
	logger logging.Logger,
	repo Repository,
	manager JobManager,
) {
	for {
		select {
		case payload := <-eventChannel:
			switch payload.Get().Type {
			case util.GetSimpleName(events.CollectionCreateEvent{}):
				handleCollectionCreate(payload, logger, repo)
			case util.GetSimpleName(events.RunCreateEvent{}):
				handleRunCreate(manager, payload, logger)
			case util.GetSimpleName(events.RunTerminateEvent{}):
				handleRunTerminate(payload, logger, manager)
			}
		case <-ctx.Done():
			return
		}
	}
}

func handleCollectionCreate(payload broker.EventPayload, logger logging.Logger, repo Repository) {
	event := events.CollectionCreateEvent{}
	parseError := payload.ParseBody(&event, events.FromJson)
	if parseError != nil {
		logger.Errors(parseError)
		payload.Acknowledge(true, false)
	} else {
		insertError := repo.SaveEvent(event, event.CollectionId)
		logger.Errors(insertError)
		payload.Acknowledge(insertError != nil, true)
	}
}

func handleRunCreate(manager JobManager, payload broker.EventPayload, logger logging.Logger) {
	startError := manager.Start(payload)
	logger.Errors(startError)
}

func handleRunTerminate(payload broker.EventPayload, logger logging.Logger, manager JobManager) {
	event := events.RunTerminateEvent{}
	parseError := payload.ParseBody(&event, events.FromJson)
	if parseError != nil {
		logger.Errors(parseError)
		payload.Acknowledge(true, false)
	} else {
		terminateError := manager.Terminate(event.RunId)
		logger.Errors(terminateError)
		payload.Acknowledge(terminateError != nil, true)
	}
}
