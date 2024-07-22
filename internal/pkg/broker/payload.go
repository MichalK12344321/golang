package broker

import (
	"github.com/wagslane/go-rabbitmq"
)

type Payload rabbitmq.Delivery

func (payload Payload) Get() *rabbitmq.Delivery {
	return (*rabbitmq.Delivery)(&payload)
}

func (payload Payload) ParseBody(target any, parser ParseBodyFn) error {
	return parser(target, payload.Body, payload.Type)
}

func (payload Payload) Acknowledge(nack bool, requeue bool) error {
	if nack {
		return payload.Nack(false, requeue)
	}
	return payload.Ack(false)
}

var _ EventPayload = new(Payload)
