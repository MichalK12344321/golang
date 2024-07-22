package collector

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

func (publisher *Publisher) PublishRunUpdateEvent(event *events.RunUpdateEvent) error {
	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.RUN_UPDATE_KEY,
	}
	pub := publisher.factoryProvider.GetPublisherFactory()(publisher.ctx, publisher.connection, publisher.logger, pubConfig)
	return pub.Publish(event)
}

func NewPublisher(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, factoryProvider broker.PublisherFactoryProvider) *Publisher {
	return &Publisher{ctx: ctx, connection: connection, factoryProvider: factoryProvider, logger: logger}
}
