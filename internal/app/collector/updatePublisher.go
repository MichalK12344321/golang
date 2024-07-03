package collector

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type CollectionUpdatePublisher struct {
	publisher broker.EventPublisher
}

func (publisher *CollectionUpdatePublisher) Exec(request *events.CollectionJobUpdatedEvent) error {
	return publisher.publisher.Publish(request)
}

func NewCollectionUpdatePublisher(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, factoryProvider broker.PublisherFactoryProvider) *CollectionUpdatePublisher {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.LOG_COLLECTION_JOB_CHANGE_KEY,
	}
	pub := factoryProvider.GetPublisherFactory()(ctx, connection, logger, pubConfig)
	return &CollectionUpdatePublisher{pub}
}
