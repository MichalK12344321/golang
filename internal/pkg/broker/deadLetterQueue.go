package broker

import (
	"lca/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewDeadLetterQueueRMQ(config *config.Config) error {
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

	err = ch.ExchangeDeclare(EXCHANGE_DEAD_LETTER, "fanout", true, false, false, false, make(amqp.Table))
	if err != nil {
		return err
	}

	args := make(map[string]any)

	_, err = ch.QueueDeclare(QUEUE_NAME_DEAD_LETTER, true, false, false, false, args)

	if err != nil {
		return err
	}

	err = ch.QueueBind(QUEUE_NAME_DEAD_LETTER, "", EXCHANGE_DEAD_LETTER, false, make(amqp.Table))
	return err
}
