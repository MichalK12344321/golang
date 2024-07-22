package web

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/logging"
)

type EventsSubscriber struct {
	subscriber broker.EventSubscriber
}

func NewEventsSubscriber(
	ctx context.Context,
	connection broker.BrokerConnection,
	logger logging.Logger,
	factoryProvider broker.SubscriberFactoryProvider,
	brokerData *BrokerEventData,
) *EventsSubscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:    broker.EXCHANGE,
		Queue:       QUEUE_NAME,
		RoutingKeys: []string{"lca.#"},
		Job:         func(c <-chan broker.EventPayload) { subscriberJob(ctx, c, logger, brokerData) },
		Args:        broker.StreamQueueArgs(),
	}

	return &EventsSubscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func subscriberJob(ctx context.Context, eventChannel <-chan broker.EventPayload, logger logging.Logger, brokerData *BrokerEventData) {
	for {
		select {
		case payload := <-eventChannel:
			eventType := payload.Get().Type
			event := &BrokerEvent{
				Type: eventType,
				Data: string(payload.Get().Body),
			}
			go func() { brokerData.Channel <- *event }()
			logger.Errors(payload.Acknowledge(false, false))
		case <-ctx.Done():
			return
		}
	}
}
