package broker

import (
	"lca/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewStreamQueueRMQ(config *config.Config) error {
	return NewStreamQueue(config, &QueueConfig{Queue: QUEUE_NAME_STREAM, RoutingKey: RUN_STREAM_KEY, Exchange: EXCHANGE})
}

func StreamQueueArgs() map[string]any {
	args := make(map[string]any)
	args["x-queue-type"] = "stream"
	args["x-max-age"] = "1D"
	return args
}

func NewStreamQueue(config *config.Config, qc *QueueConfig) error {
	conn, err := amqp.Dial(config.RMQ.Uri)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	args := StreamQueueArgs()

	err = ch.ExchangeDeclare(qc.Exchange, "topic", true, false, false, false, make(amqp.Table))
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(qc.Queue, true, false, false, false, args)

	if err != nil {
		return err
	}

	err = ch.QueueBind(qc.Queue, qc.RoutingKey, qc.Exchange, false, make(amqp.Table))
	return err
}
