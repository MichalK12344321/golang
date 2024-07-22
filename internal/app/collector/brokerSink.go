package collector

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"

	"github.com/google/uuid"
)

type BrokerSink struct {
	brokerConnection         broker.BrokerConnection
	pub                      broker.EventPublisher
	id                       uuid.UUID
	publisherFactoryProvider broker.PublisherFactoryProvider
	BaseSink
}

func NewBrokerSink(
	ctx context.Context,
	event *events.RunCreateEvent,
	brokerConnection broker.BrokerConnection,
	logger logging.Logger,
	jobData JobData,
	publisherFactory broker.PublisherFactoryProvider,
) *BrokerSink {
	result := &BrokerSink{
		brokerConnection: brokerConnection,
		BaseSink: BaseSink{
			logger:  logger,
			ctx:     ctx,
			event:   event,
			jobData: jobData,
		},
		id:                       event.RunId,
		publisherFactoryProvider: publisherFactory,
	}
	return result
}

func (r *BrokerSink) Initialize() error {
	headers := make(map[string]any)
	headers["x-stream-filter-value"] = r.id.String()

	pubConfig := &broker.PublisherConfig{
		Exchange:   broker.EXCHANGE,
		RoutingKey: broker.RUN_STREAM_KEY,
		Headers:    headers,
	}

	r.pub = r.publisherFactoryProvider.GetPublisherFactory()(
		r.ctx,
		(r.brokerConnection),
		r.logger,
		pubConfig,
	)

	return nil
}

func (r *BrokerSink) AppendStdout(data []byte) {
	err := r.pub.Publish(&events.AppendLineEvent{
		RunId: r.id,
		File:  "stdout.log",
		Line:  string(data),
	})
	if err != nil {
		r.jobData.ErrorChannel() <- err
	}
}

func (r *BrokerSink) AppendStderr(data []byte) {
	err := r.pub.Publish(&events.AppendLineEvent{
		RunId: r.id,
		File:  "stderr.log",
		Line:  string(data),
	})
	if err != nil {
		r.jobData.ErrorChannel() <- err
	}
}

func (r *BrokerSink) Finalize() error {
	return nil
}

var _ CollectionSink = new(BrokerSink)
