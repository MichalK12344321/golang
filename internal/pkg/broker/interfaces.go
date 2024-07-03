package broker

import (
	"context"
	"lca/internal/pkg/logging"

	"github.com/wagslane/go-rabbitmq"
)

type PublisherConfig struct {
	Exchange   string
	RoutingKey string
}

type SubscribeConfig struct {
	Exchange   string
	Queue      string
	RoutingKey string
	Job        func(<-chan EventPayload)
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
	Acknowledge(err ...error) error
}

//counterfeiter:generate . BrokerConnection
type BrokerConnection interface {
}
