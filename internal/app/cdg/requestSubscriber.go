package cdg

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type CollectionScheduleRequestSubscriber struct {
	subscriber broker.EventSubscriber
}

func NewCollectionScheduleRequestSubscriber(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, pub *CollectionInitPublisher, repo LogCollectionRepository, factoryProvider broker.SubscriberFactoryProvider) *CollectionScheduleRequestSubscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:   broker.EXCHANGE,
		Queue:      "cdg-collection-requests",
		RoutingKey: broker.LOG_COLLECTION_REQUEST_KEY,
		Job:        func(c <-chan broker.EventPayload) { CollectionScheduleRequestSubscriberJob(ctx, c, pub, logger, repo) },
	}

	return &CollectionScheduleRequestSubscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func CollectionScheduleRequestSubscriberJob(ctx context.Context, eventChannel <-chan broker.EventPayload, pub *CollectionInitPublisher, logger logging.Logger, repo LogCollectionRepository) {
	for {
		select {
		case payload := <-eventChannel:
			handleScheduleRequestPayload(payload, pub, logger, repo)
		case <-ctx.Done():
			return
		}
	}
}

func handleScheduleRequestPayload(payload broker.EventPayload, pub *CollectionInitPublisher, logger logging.Logger, repo LogCollectionRepository) {
	event := &events.CollectionScheduleEvent{}
	var parseError error = payload.ParseBody(event, events.FromJson)
	if parseError != nil {
		ackError := payload.Acknowledge(parseError)
		logger.Errors(ackError, parseError)
		return
	}

	initEvent := events.CollectionJobUpdatedEvent{
		CollectionScheduleEvent: *event,
		Status:                  events.CollectionStatusInitialized,
	}
	insertError := repo.Insert((*dto.CollectionDto)(&initEvent))
	if insertError != nil {
		ackError := payload.Acknowledge(insertError)
		logger.Errors(ackError, insertError)
		return
	}

	publishError := pub.Publish(&initEvent)
	if publishError != nil {
		ackError := payload.Acknowledge(publishError)
		logger.Errors(ackError, publishError)
		return
	}

	ackError := payload.Acknowledge(nil)
	logger.Errors(ackError)
}
