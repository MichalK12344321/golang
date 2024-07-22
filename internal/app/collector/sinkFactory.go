package collector

import (
	"context"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type SinkFactory struct {
	brokerConnection         broker.BrokerConnection
	logger                   logging.Logger
	config                   *config.Config
	publisherFactoryProvider broker.PublisherFactoryProvider
}

func NewSinkFactory(
	brokerConnection broker.BrokerConnection,
	logger logging.Logger,
	config *config.Config,
	publisherFactoryProvider broker.PublisherFactoryProvider,
) *SinkFactory {
	return &SinkFactory{
		brokerConnection:         brokerConnection,
		logger:                   logger,
		config:                   config,
		publisherFactoryProvider: publisherFactoryProvider,
	}
}

func (factory *SinkFactory) New(ctx context.Context, event *events.RunCreateEvent, jobData JobData) []CollectionSink {
	return []CollectionSink{
		NewBrokerSink(ctx, event, factory.brokerConnection, factory.logger, jobData, factory.publisherFactoryProvider),
		NewStorageSink(factory.config, ctx, event, factory.logger, jobData),
		NewConsoleSink(factory.logger),
	}
}

var _ CollectionSinkFactory = new(SinkFactory)
