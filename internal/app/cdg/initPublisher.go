package cdg

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type CollectionInitPublisher struct {
	publisher broker.EventPublisher
}

func (publisher *CollectionInitPublisher) Publish(request *events.CollectionJobUpdatedEvent) error {
	return publisher.publisher.Publish(request)
}

func NewCollectionInitPublisher(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, factoryProvider broker.PublisherFactoryProvider) *CollectionInitPublisher {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.LOG_COLLECTION_JOB_INIT_KEY,
	}
	pub := factoryProvider.GetPublisherFactory()(ctx, connection, logger, pubConfig)
	return &CollectionInitPublisher{pub}
}
