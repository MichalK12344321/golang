package web

import (
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/database"
	"lca/internal/pkg/logging"

	"github.com/defval/di"
)

const QUEUE_NAME = "web"

func ProvideOptions() di.Option {
	return di.Options(
		di.Provide(logging.GetLogger, di.As(new(logging.Logger))),
		di.Provide(server.NewAppController, di.As(new(server.Controller))),
		di.Provide(NewPreviewController, di.As(new(server.Controller))),
		di.Provide(NewEventController, di.As(new(server.Controller))),
		di.Provide(serverConfig),
		di.Provide(brokerConfig),
		di.Provide(func() (*config.Config, error) {
			return config.NewConfig("configs/web.yaml")
		}),

		di.Provide(NewBrokerEventData),
		di.Provide(newHub),

		di.Provide(broker.NewBrokerConnection, di.As(new(broker.BrokerConnection))),
		di.Provide(broker.NewSubscriberRMQFactoryProvider, di.As(new(broker.SubscriberFactoryProvider))),
		di.Provide(NewEventsSubscriber),
		di.Invoke(broker.NewDeadLetterQueueRMQ),
		di.Provide(database.NewContext),
		di.Invoke(invoke),
	)
}

func serverConfig(config *config.Config) *server.ServerConfig {
	return &server.ServerConfig{
		Version: config.Version,
		Name:    config.Name,
		Port:    config.Port,
	}
}

func brokerConfig(config *config.Config) *broker.BrokerConfig {
	return &broker.BrokerConfig{
		ClientName: config.Name,
		Uri:        config.RMQ.Uri,
	}
}

func invoke(_ *EventsSubscriber, hub *Hub) error {
	go hub.run()
	return nil
}
