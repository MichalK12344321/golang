package broker

import (
	"context"
	"lca/internal/pkg/logging"

	"github.com/wagslane/go-rabbitmq"
)

type QueueConfig struct {
	Exchange   string
	Queue      string
	RoutingKey string
}

type PublisherConfig struct {
	Exchange   string
	RoutingKey string
	Headers    map[string]any
}

type SubscribeConfig struct {
	Exchange    string
	Queue       string
	RoutingKeys []string
	Args        map[string]any
	Job         func(<-chan EventPayload)
}

type BrokerConfig struct {
	ClientName string
	Uri        string
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . EventSubscriber
type EventSubscriber interface {
	Channel() chan EventPayload
}

type SubscriberFactory func(ctx context.Context, connection BrokerConnection, logger logging.Logger, subConfig *SubscribeConfig) EventSubscriber
type PublisherFactory func(ctx context.Context, connection BrokerConnection, logger logging.Logger, pubConfig *PublisherConfig) EventPublisher

//counterfeiter:generate . SubscriberFactoryProvider
type SubscriberFactoryProvider interface {
	GetSubscriberFactory() SubscriberFactory
}

//counterfeiter:generate . PublisherFactoryProvider
type PublisherFactoryProvider interface {
	GetPublisherFactory() PublisherFactory
}

//counterfeiter:generate . EventPublisher
type EventPublisher interface {
	Publish(body any) error
}

type ParseBodyFn func(target any, bytes []byte, objType string) error

//counterfeiter:generate . EventPayload
type EventPayload interface {
	Get() *rabbitmq.Delivery
	ParseBody(any, ParseBodyFn) error
	Acknowledge(nack bool, requeue bool) error
}

//counterfeiter:generate . BrokerConnection
type BrokerConnection interface {
}
