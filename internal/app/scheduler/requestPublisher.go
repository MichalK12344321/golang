package scheduler

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type CollectionRequestPublisher struct {
	publisher broker.EventPublisher
}

func (publisher *CollectionRequestPublisher) Publish(request *events.CollectionScheduleEvent) error {
	return publisher.publisher.Publish(request)
}

func NewCollectionRequestPublisher(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, factoryProvider broker.PublisherFactoryProvider) *CollectionRequestPublisher {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.LOG_COLLECTION_REQUEST_KEY,
	}
	pub := factoryProvider.GetPublisherFactory()(ctx, connection, logger, pubConfig)
	return &CollectionRequestPublisher{pub}
}
