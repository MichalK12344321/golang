package cdg

import (
	"context"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/database"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
)

type Subscriber struct {
	subscriber broker.EventSubscriber
}

func NewSubscriber(
	ctx context.Context,
	connection broker.BrokerConnection,
	logger logging.Logger,
	config *config.Config,
	publisher *Publisher,
	factoryProvider broker.SubscriberFactoryProvider,
	dataContext database.DataContext,
	repo LogCollectionRepository,
) *Subscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:    broker.EXCHANGE,
		Queue:       QUEUE_NAME,
		RoutingKeys: []string{broker.COLLECTION_CREATE_KEY, broker.RUN_CREATE_KEY, broker.RUN_UPDATE_KEY},
		Job: func(c <-chan broker.EventPayload) {
			SubscriberJob(ctx, c, logger, dataContext, repo)
		},
	}
	return &Subscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func SubscriberJob(
	ctx context.Context,
	eventChannel <-chan broker.EventPayload,
	logger logging.Logger,
	dataContext database.DataContext,
	repo LogCollectionRepository,
) {
	for {
		select {
		case payload := <-eventChannel:
			switch payload.Get().Type {
			case util.GetSimpleName(events.CollectionCreateEvent{}):
				handleCollectionCreate(payload, logger, repo)
			case util.GetSimpleName(events.RunCreateEvent{}):
				handleRunCreatePayload(payload, logger, repo)
			case util.GetSimpleName(events.RunUpdateEvent{}):
				handleRunUpdatePayload(payload, logger, repo)
			}
		case <-ctx.Done():
			return
		}
	}
}

func handleCollectionCreate(payload broker.EventPayload, logger logging.Logger, repo LogCollectionRepository) {
	event := &events.CollectionCreateEvent{}
	var parseError error = payload.ParseBody(event, events.FromJson)
	if parseError != nil {
		ackError := payload.Acknowledge(true, false)
		logger.Errors(ackError, parseError)
		return
	}

	insertError := repo.Insert(&dto.CollectionDto{
		CollectionCreateEvent: *event,
		Runs:                  make([]*dto.RunDto, 0),
	})
	if insertError != nil {
		ackError := payload.Acknowledge(true, true)
		logger.Errors(ackError, insertError)
		return
	}

	ackError := payload.Acknowledge(false, false)
	logger.Errors(ackError)
}

func handleRunCreatePayload(payload broker.EventPayload, logger logging.Logger, repo LogCollectionRepository) {
	event := &events.RunCreateEvent{}
	var parseError error = payload.ParseBody(event, events.FromJson)

	if parseError != nil {
		ackError := payload.Acknowledge(false, false)
		logger.Errors(parseError, ackError)
		return
	}

	runUpdateEvent := &events.RunUpdateEvent{RunCreateEvent: *event, Status: events.CollectionStatusCreated}
	updateError := repo.Update((*dto.RunDto)(runUpdateEvent))
	ackError := payload.Acknowledge(updateError != nil, true)
	logger.Errors(updateError, ackError)
}

func handleRunUpdatePayload(payload broker.EventPayload, logger logging.Logger, repo LogCollectionRepository) {
	event := &events.RunUpdateEvent{}
	var parseError error = payload.ParseBody(event, events.FromJson)

	if parseError != nil {
		ackError := payload.Acknowledge(false, false)
		logger.Errors(parseError, ackError)
		return
	}

	updateError := repo.Update((*dto.RunDto)(event))
	ackError := payload.Acknowledge(updateError != nil, true)
	logger.Errors(updateError, ackError)
}
