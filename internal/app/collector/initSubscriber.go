package collector

import (
	"context"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/storage"
)

type CollectionInitSubscriber struct {
	subscriber broker.EventSubscriber
}

func NewCollectionInitSubscriber(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, config *config.Config, publisher *CollectionUpdatePublisher, store storage.Storage, factoryProvider broker.SubscriberFactoryProvider, sshFactoryProvider ssh.SSHFactoryProvider) *CollectionInitSubscriber {
	subConfig := &broker.SubscribeConfig{
		Exchange:   broker.EXCHANGE,
		Queue:      "collector-collection-inits",
		RoutingKey: broker.LOG_COLLECTION_JOB_INIT_KEY,
		Job: func(c <-chan broker.EventPayload) {
			CollectionInitSubscriberJob(ctx, c, logger, publisher, store, sshFactoryProvider)
		},
	}
	return &CollectionInitSubscriber{factoryProvider.GetSubscriberFactory()(ctx, connection, logger, subConfig)}
}

func CollectionInitSubscriberJob(ctx context.Context, eventChannel <-chan broker.EventPayload, logger logging.Logger, publisher *CollectionUpdatePublisher, store storage.Storage, sshFactoryProvider ssh.SSHFactoryProvider) {
	for {
		select {
		case payload := <-eventChannel:
			handleInitPayload(ctx, payload, publisher, logger, store, sshFactoryProvider)
		case <-ctx.Done():
			return
		}
	}
}

func handleInitPayload(ctx context.Context, payload broker.EventPayload, publisher *CollectionUpdatePublisher, logger logging.Logger, store storage.Storage, sshFactoryProvider ssh.SSHFactoryProvider) {
	event := events.CollectionJobUpdatedEvent{}
	var parseError error = payload.ParseBody(&event, events.FromJson)
	event.Status = events.CollectionStatusStarted

	if parseError != nil {
		handleInitPayloadResult(logger, payload, &event, publisher, "", parseError)
		return
	}

	publisher.Exec(&event)

	client := sshFactoryProvider.GetSSHFactory()()
	sshConnectError := client.Connect(&event)
	if sshConnectError != nil {
		handleInitPayloadResult(logger, payload, &event, publisher, "", sshConnectError)
		return
	}

	stdout, stderr, runError := client.Run(ctx, event.Script)
	if runError != nil {
		handleInitPayloadResult(logger, payload, &event, publisher, "", runError)
		return
	}

	path, createFileError := store.CreateFile(event.Id, stdout, stderr)
	if createFileError != nil {
		handleInitPayloadResult(logger, payload, &event, publisher, "", createFileError)
		return
	}

	handleInitPayloadResult(logger, payload, &event, publisher, path, nil)
}

func handleInitPayloadResult(logger logging.Logger, payload broker.EventPayload, event *events.CollectionJobUpdatedEvent, publisher *CollectionUpdatePublisher, path string, err error) {
	if err != nil {
		event.Status = events.CollectionStatusFailure
		event.Error = err.Error()
	} else {
		event.Status = events.CollectionStatusSuccess
		event.Path = path
	}
	publishError := publisher.Exec(event)
	ackError := payload.Acknowledge(err, publishError)
	logger.Errors(ackError, err, publishError)
}
