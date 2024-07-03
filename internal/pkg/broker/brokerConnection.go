package broker

import (
	"context"
	"lca/internal/pkg/logging"
	"time"

	"github.com/aaronjan/hunch"
	"github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
)

type BrokerConnectionRMQ rabbitmq.Conn

func NewBrokerConnection(ctx context.Context, config *BrokerConfig, logger logging.Logger) *BrokerConnectionRMQ {
	connectionConfig := rabbitmq.Config{Properties: amqp091.Table{}}
	connectionConfig.Properties.SetClientConnectionName(config.ClientName)

	connection, err := hunch.Retry(
		ctx,
		8,
		func(ctx context.Context) (interface{}, error) {
			connection, err := rabbitmq.NewConn(
				config.Uri,
				rabbitmq.WithConnectionOptionsConfig(connectionConfig),
				rabbitmq.WithConnectionOptionsLogger(logger),
			)

			if err != nil {
				time.Sleep(time.Second * 1)
			}

			return connection, err
		},
	)

	if err != nil {
		panic(err)
	}

	rmqConn := connection.(*rabbitmq.Conn)

	go func() {
		<-ctx.Done()
		rmqConn.Close()
	}()

	return (*BrokerConnectionRMQ)(rmqConn)
}

var _ BrokerConnection = new(BrokerConnectionRMQ)
