package scheduler

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/database"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
)

type RunUpdateSubscriber struct {
	subscriber broker.EventSubscriber
}

func NewSubscriber(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, dataContext database.DataContext, factoryProvider broker.SubscriberFactoryProvider) *RunUpdateSubscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:    broker.EXCHANGE,
		Queue:       QUEUE_NAME,
		RoutingKeys: []string{broker.RUN_UPDATE_KEY},
		Job:         func(c <-chan broker.EventPayload) { RunUpdateSubscriberJob(ctx, c, logger, dataContext) },
	}

	return &RunUpdateSubscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func RunUpdateSubscriberJob(ctx context.Context, eventChannel <-chan broker.EventPayload, logger logging.Logger, dataContext database.DataContext) {
	for {
		select {
		case payload := <-eventChannel:
			switch payload.Get().Type {
			case util.GetSimpleName(events.RunUpdateEvent{}):
				handleUpdatePayload(payload, logger, dataContext)
			}
		case <-ctx.Done():
			return
		}
	}
}

func handleUpdatePayload(payload broker.EventPayload, logger logging.Logger, dataContext database.DataContext) {
	event := &events.RunUpdateEvent{}
	var parseError error = payload.ParseBody(event, events.FromJson)

	if parseError != nil {
		ackError := payload.Acknowledge(false, false)
		logger.Errors(parseError, ackError)
		return
	}

	updateError := InsertEvent(dataContext, event, event.RunId)
	ackError := payload.Acknowledge(updateError != nil, true)
	logger.Errors(updateError, ackError)
}
