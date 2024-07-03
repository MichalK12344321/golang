package collector

import (
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/storage"

	"github.com/defval/di"
)

func ProvideOptions() di.Option {
	return di.Options(
		di.Provide(logging.GetLogger, di.As(new(logging.Logger))),
		di.Provide(NewVersionController, di.As(new(server.Controller))),
		di.Provide(serverConfig),
		di.Provide(brokerConfig),
		di.Provide(func() (*config.Config, error) {
			return config.NewConfig("configs/collector.yaml")
		}),
		di.Provide(broker.NewBrokerConnection, di.As(new(broker.BrokerConnection))),
		di.Provide(broker.NewPublisherRMQFactoryProvider, di.As(new(broker.PublisherFactoryProvider))),
		di.Provide(broker.NewSubscriberRMQFactoryProvider, di.As(new(broker.SubscriberFactoryProvider))),
		di.Provide(NewCollectionInitSubscriber),
		di.Provide(NewCollectionUpdatePublisher),
		di.Provide(storage.NewDiskStorage, di.As(new(storage.Storage))),
		di.Provide(ssh.NewVSSHFactoryProvider, di.As(new(ssh.SSHFactoryProvider))),

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

func invoke(*CollectionInitSubscriber) {}
