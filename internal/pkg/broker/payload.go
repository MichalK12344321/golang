package broker

import (
	"lca/internal/pkg/util"

	"github.com/wagslane/go-rabbitmq"
)

type Payload rabbitmq.Delivery

func (payload Payload) Get() *rabbitmq.Delivery {
	return (*rabbitmq.Delivery)(&payload)
}

func (payload Payload) ParseBody(target any, parser ParseBodyFn) error {
	return parser(target, payload.Body, payload.Type)
}

func (payload Payload) Acknowledge(err ...error) error {
	if util.AnyErrors(err...) {
		return payload.Nack(false, true)
	}
	return payload.Ack(false)
}

var _ EventPayload = new(Payload)
