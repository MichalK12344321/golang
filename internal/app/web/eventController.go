package web

import (
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/logging"
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

type EventController struct {
	server.ControllerBase
	config       *config.Config
	logger       logging.Logger
	brokerEvents *BrokerEventData
	hub          *Hub
}

func NewEventController(config *config.Config, logger logging.Logger, brokerEvent *BrokerEventData, hub *Hub) *EventController {
	return &EventController{config: config, logger: logger, brokerEvents: brokerEvent, hub: hub}
}

func (c *EventController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet,
			"/event",
			endpoint.Handler(c.ConnectSocket),
			endpoint.Tags("event"),
			endpoint.Response(http.StatusOK, "Response"),
		),
	}
}

func (c *EventController) ConnectSocket(writer http.ResponseWriter, req *http.Request) {
	serveWs(c.hub, writer, req)
}
