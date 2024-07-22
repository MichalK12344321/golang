package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
)

type PublisherRMQ struct {
	config    *PublisherConfig
	publisher *rabbitmq.Publisher
	logger    logging.Logger
}

func NewPublisherRMQ(ctx context.Context, connection *BrokerConnectionRMQ, logger logging.Logger, pubConfig *PublisherConfig) *PublisherRMQ {
	publisher, err := rabbitmq.NewPublisher(
		(*rabbitmq.Conn)(connection),
		rabbitmq.WithPublisherOptionsLogger(logging.GetLogger()),
		rabbitmq.WithPublisherOptionsExchangeName(pubConfig.Exchange),
		rabbitmq.WithPublisherOptionsExchangeKind("topic"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		publisher.Close()
	}()

	return &PublisherRMQ{logger: logger, config: pubConfig, publisher: publisher}
}

func (pub *PublisherRMQ) Publish(data any) error {
	dataType := util.GetSimpleName(data)
	if dataType == "" {
		return fmt.Errorf("data type must be provided")
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to parse message %s", err)
	}
	correlationId := strings.ReplaceAll(uuid.New().String(), "-", "")
	pub.logger.Debug("PUBLISH '%s' [%s | %s]\n%s", dataType, pub.config.RoutingKey, correlationId, jsonData)
	err = pub.publisher.Publish(
		jsonData,
		[]string{pub.config.RoutingKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(pub.config.Exchange),
		rabbitmq.WithPublishOptionsTimestamp(time.Now()),
		rabbitmq.WithPublishOptionsCorrelationID(correlationId),
		rabbitmq.WithPublishOptionsType(dataType),
		rabbitmq.WithPublishOptionsHeaders(pub.config.Headers),
	)
	return err
}

var _ EventPublisher = new(PublisherRMQ)
