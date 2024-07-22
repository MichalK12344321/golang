package broker

import (
	"context"
	"lca/internal/pkg/logging"
)

type PublisherRMQFactoryProvider struct{}

func NewPublisherRMQFactoryProvider() *PublisherRMQFactoryProvider {
	return &PublisherRMQFactoryProvider{}
}

func (p *PublisherRMQFactoryProvider) GetPublisherFactory() PublisherFactory {
	return func(ctx context.Context, connection BrokerConnection, logger logging.Logger, pubConfig *PublisherConfig) EventPublisher {
		return NewPublisherRMQ(ctx, connection.(*BrokerConnectionRMQ), logger, pubConfig)
	}
}

var _ PublisherFactoryProvider = new(PublisherRMQFactoryProvider)

type SubscriberRMQFactoryProvider struct{}

func NewSubscriberRMQFactoryProvider() *SubscriberRMQFactoryProvider {
	return &SubscriberRMQFactoryProvider{}
}

func (p *SubscriberRMQFactoryProvider) GetSubscriberFactory() SubscriberFactory {
	return func(ctx context.Context, connection BrokerConnection, logger logging.Logger, pubConfig *SubscribeConfig) EventSubscriber {
		return NewSubscriberRMQ(ctx, connection.(*BrokerConnectionRMQ), logger, pubConfig)
	}
}

var _ SubscriberFactoryProvider = new(SubscriberRMQFactoryProvider)
