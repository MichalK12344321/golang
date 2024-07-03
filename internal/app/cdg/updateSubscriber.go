package cdg

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type CollectionUpdateSubscriber struct {
	subscriber broker.EventSubscriber
}

func NewCollectionUpdateSubscriber(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, repo LogCollectionRepository, factoryProvider broker.SubscriberFactoryProvider) *CollectionUpdateSubscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:   broker.EXCHANGE,
		Queue:      "cdg-collection-updates",
		RoutingKey: broker.LOG_COLLECTION_JOB_CHANGE_KEY,
		Job:        func(c <-chan broker.EventPayload) { CollectionUpdateSubscriberJob(ctx, c, logger, repo) },
	}

	return &CollectionUpdateSubscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func CollectionUpdateSubscriberJob(ctx context.Context, eventChannel <-chan broker.EventPayload, logger logging.Logger, repo LogCollectionRepository) {
	for {
		select {
		case payload := <-eventChannel:
			handleUpdatePayload(payload, logger, repo)
		case <-ctx.Done():
			return
		}
	}
}

func handleUpdatePayload(payload broker.EventPayload, logger logging.Logger, repo LogCollectionRepository) {
	event := &events.CollectionJobUpdatedEvent{}
	var parseError error = payload.ParseBody(event, events.FromJson)

	if parseError != nil {
		ackError := payload.Acknowledge(parseError)
		logger.Errors(parseError, ackError)
		return
	}

	updateError := repo.Update((*dto.CollectionDto)(event))
	ackError := payload.Acknowledge(updateError)
	logger.Errors(updateError, ackError)
}
