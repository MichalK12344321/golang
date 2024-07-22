package collector

import (
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/database"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/storage"

	"github.com/defval/di"
)

const QUEUE_NAME = "collector"

func ProvideOptions() di.Option {
	return di.Options(
		di.Provide(logging.GetLogger, di.As(new(logging.Logger))),
		di.Provide(server.NewAppController, di.As(new(server.Controller))),
		di.Provide(serverConfig),
		di.Provide(brokerConfig),
		di.Provide(func() (*config.Config, error) {
			return config.NewConfig("configs/collector.yaml")
		}),
		di.Provide(broker.NewBrokerConnection, di.As(new(broker.BrokerConnection))),
		di.Provide(broker.NewPublisherRMQFactoryProvider, di.As(new(broker.PublisherFactoryProvider))),
		di.Provide(broker.NewSubscriberRMQFactoryProvider, di.As(new(broker.SubscriberFactoryProvider))),
		di.Provide(NewSubscriber),
		di.Provide(NewPublisher),
		di.Provide(storage.NewDiskStorage, di.As(new(storage.Storage))),
		di.Provide(ssh.NewVSSHFactoryProvider, di.As(new(ssh.SSHFactoryProvider))),
		di.Provide(NewSinkFactory, di.As(new(CollectionSinkFactory))),
		di.Provide(NewCollectionJobManager, di.As(new(JobManager))),

		di.Provide(NewGoRunner, di.As(new(CollectionRunner))),
		di.Provide(NewSSHRunner, di.As(new(CollectionRunner))),

		di.Provide(database.NewContext),
		di.Provide(NewDataRepository, di.As(new(Repository))),

		di.Invoke(broker.NewStreamQueueRMQ),
		di.Invoke(broker.NewDeadLetterQueueRMQ),

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

func invoke(
	_ *Subscriber,
	dataContext database.DataContext,
) error {
	return dataContext.RunMigrations(ENTITIES...)
}
