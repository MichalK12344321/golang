package broker

import (
	"context"
	"errors"
	"lca/internal/pkg/logging"
	"sync"

	"github.com/samber/lo"
	"github.com/wagslane/go-rabbitmq"
)

type SubscriberRMQ struct {
	connection *BrokerConnectionRMQ
	logger     logging.Logger
	channel    chan EventPayload
}

func NewSubscriberRMQ(ctx context.Context, connection *BrokerConnectionRMQ, logger logging.Logger, subConfig *SubscribeConfig) *SubscriberRMQ {
	result := &SubscriberRMQ{connection: connection, logger: logger}
	result.subscribe(ctx, subConfig)
	return result
}

func (sub *SubscriberRMQ) Channel() chan EventPayload {
	return sub.channel
}

func (sub *SubscriberRMQ) subscribe(ctx context.Context, subConfig *SubscribeConfig) {
	sub.channel = make(chan EventPayload)
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go sub.subscribeQueue(ctx, sub.channel, subConfig, &waitGroup)

	go func() {
		go subConfig.Job(sub.channel)
		waitGroup.Wait()
		close(sub.channel)
	}()
}

func (sub *SubscriberRMQ) subscribeQueue(ctx context.Context, channel chan<- EventPayload, subConfig *SubscribeConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	if subConfig.Args == nil {
		subConfig.Args = map[string]any{}
	}

	if subConfig.Args["x-queue-type"] != "stream" {
		subConfig.Args["x-dead-letter-exchange"] = EXCHANGE_DEAD_LETTER
	}

	optionFuncs := lo.Map(
		subConfig.RoutingKeys,
		func(routingKey string, _ int) func(*rabbitmq.ConsumerOptions) {
			return rabbitmq.WithConsumerOptionsRoutingKey(routingKey)
		},
	)

	optionFuncs = append(optionFuncs, rabbitmq.WithConsumerOptionsExchangeName(subConfig.Exchange))
	optionFuncs = append(optionFuncs, rabbitmq.WithConsumerOptionsQueueArgs(subConfig.Args))
	optionFuncs = append(optionFuncs, rabbitmq.WithConsumerOptionsQueueDurable)
	optionFuncs = append(optionFuncs, rabbitmq.WithConsumerOptionsLogger(logging.GetLogger()))

	consumer, err := rabbitmq.NewConsumer(
		(*rabbitmq.Conn)(sub.connection),
		subConfig.Queue,
		optionFuncs...,
	)

	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		consumer.Close()
	}()

	err = consumer.Run(func(delivery rabbitmq.Delivery) rabbitmq.Action {
		sub.logger.Debug(
			"CONSUME '%s' [%s | %s | %s]\n%s",
			delivery.Type,
			subConfig.Queue,
			delivery.RoutingKey,
			delivery.CorrelationId,
			delivery.Body,
		)
		if delivery.Type == "" {
			panic(errors.New("'Type' property is required"))
		}
		channel <- (EventPayload)(Payload(delivery))
		return rabbitmq.Manual
	})

	if err != nil {
		sub.logger.Fatalf("%s", err)
	}
}

var _ EventSubscriber = new(SubscriberRMQ)
