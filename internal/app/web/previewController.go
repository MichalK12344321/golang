package web

import (
	"context"
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/logging"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/types"
)

type PreviewController struct {
	server.ControllerBase
	config *config.Config
	logger logging.Logger
}

func NewPreviewController(config *config.Config, logger logging.Logger) *PreviewController {
	return &PreviewController{config: config, logger: logger}
}

func (c *PreviewController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet,
			"/run/{id}/preview",
			endpoint.Path("id", types.String, "Id of the run", true),
			endpoint.Handler(c.ConnectSocket),
			endpoint.Tags("run"),
			endpoint.Response(http.StatusOK, "Response"),
		),
	}
}

func (c *PreviewController) ConnectSocket(writer http.ResponseWriter, req *http.Request) {
	idPath := strings.Split(req.URL.Path, "/")[2]
	id, err := uuid.Parse(idPath)
	if err != nil {
		c.WriteResponse(writer, nil, err, "text/plain")
		return
	}

	var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // use default options

	cc, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		c.WriteResponse(writer, nil, err, "text/plain")
		return
	}
	defer cc.Close()

	c.Sub(req.Context(), id, cc.WriteJSON)
}

func (c *PreviewController) Sub(ctx context.Context, id uuid.UUID, handler func(any) error) {
	conn, err := amqp.Dial(c.config.RMQ.Uri)
	c.logger.Errors(err)
	defer conn.Close()

	ch, err := conn.Channel()
	ch.Qos(1024, 0, false)
	c.logger.Errors(err)
	defer ch.Close()

	args := make(map[string]any)
	args["x-queue-type"] = "stream"
	args["x-stream-offset"] = "first"
	args["x-stream-filter"] = id.String()

	msgs, err := ch.Consume(
		broker.QUEUE_NAME_STREAM, // queue
		QUEUE_NAME,               // consumer
		false,                    // auto ack
		false,                    // exclusive
		false,                    // no local
		false,                    // no wait
		args,                     // args
	)
	c.logger.Errors(err)

	go func() {
		for delivery := range msgs {
			if err == nil {
				c.logger.Errors(err, handler(&BrokerEvent{Type: delivery.Type, Data: string(delivery.Body)}))
			}
		}
	}()

	// TODO: fix broken pipe

	<-ctx.Done()
}
