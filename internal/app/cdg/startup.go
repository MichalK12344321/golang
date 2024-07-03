package cdg

import (
	"lca/api/server"
	"lca/internal/app/cdg/database"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/storage"

	"github.com/defval/di"
)

func ProvideOptions() di.Option {
	return di.Options(
		di.Provide(logging.GetLogger, di.As(new(logging.Logger))),
		di.Provide(NewVersionController, di.As(new(server.Controller))),
		di.Provide(NewCollectionController, di.As(new(server.Controller))),
		di.Provide(serverConfig),
		di.Provide(brokerConfig),
		di.Provide(func() (*config.Config, error) {
			return config.NewConfig("configs/cdg.yaml")
		}),

		di.Provide(NewCollectionInitPublisher),
		di.Provide(NewCollectionScheduleRequestSubscriber),
		di.Provide(NewCollectionUpdateSubscriber),
		di.Provide(database.NewCollectionRepository, di.As(new(LogCollectionRepository))),
		di.Provide(storage.NewDiskStorage, di.As(new(storage.Storage))),
		di.Provide(newAppContext, di.As(new(database.DataContext))),
		di.Provide(broker.NewBrokerConnection, di.As(new(broker.BrokerConnection))),
		di.Provide(broker.NewPublisherRMQFactoryProvider, di.As(new(broker.PublisherFactoryProvider))),
		di.Provide(broker.NewSubscriberRMQFactoryProvider, di.As(new(broker.SubscriberFactoryProvider))),
		di.Invoke(invoke),
	)
}

func newAppContext(config *config.Config) (*database.AppContext, error) {
	return database.NewAppContext(config.DB.CONNECTION)
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

func invoke(_ *CollectionInitPublisher, _ *CollectionScheduleRequestSubscriber, _ *CollectionUpdateSubscriber, dataContext database.DataContext) error {
	return dataContext.Database().Migrator().AutoMigrate(&database.Collection{})
}
