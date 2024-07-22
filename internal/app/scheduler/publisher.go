package scheduler

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type Publisher struct {
	ctx             context.Context
	connection      broker.BrokerConnection
	factoryProvider broker.PublisherFactoryProvider
	logger          logging.Logger
}

func (publisher *Publisher) PublishCollectionCreateEvent(event *events.CollectionCreateEvent) error {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.COLLECTION_CREATE_KEY,
	}
	pub := publisher.factoryProvider.GetPublisherFactory()(publisher.ctx, publisher.connection, publisher.logger, pubConfig)
	return pub.Publish(event)
}

func (publisher *Publisher) PublishRunCreateEvent(event *events.RunCreateEvent) error {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.RUN_CREATE_KEY,
	}
	pub := publisher.factoryProvider.GetPublisherFactory()(publisher.ctx, publisher.connection, publisher.logger, pubConfig)
	return pub.Publish(event)
}

func (publisher *Publisher) PublishRunTerminateEvent(event *events.RunTerminateEvent) error {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.RUN_TERMINATE_KEY,
	}
	pub := publisher.factoryProvider.GetPublisherFactory()(publisher.ctx, publisher.connection, publisher.logger, pubConfig)
	return pub.Publish(event)
}

func NewPublisher(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, factoryProvider broker.PublisherFactoryProvider) *Publisher {
	return &Publisher{ctx: ctx, connection: connection, factoryProvider: factoryProvider, logger: logger}
}
