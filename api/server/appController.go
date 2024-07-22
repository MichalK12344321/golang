package server

import (
	"lca/internal/config"
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

type AppController struct {
	ControllerBase
	config *config.Config
}

func NewAppController(config *config.Config) *AppController {
	return &AppController{config: config}
}

func (c *AppController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet,
			"/version",
			endpoint.Handler(c.GetVersion),
			endpoint.Summary("Get application version"),
			endpoint.Response(http.StatusOK, "Version number", endpoint.SchemaResponseOption(&ValueObject{})),
			endpoint.Produces("text/plain"),
		),
		endpoint.New(
			http.MethodGet,
			"/healthz",
			endpoint.Handler(c.GetHealthz),
			endpoint.Summary("Get application status"),
			endpoint.Response(http.StatusOK, "", endpoint.SchemaResponseOption(&ValueObject{})),
			endpoint.Produces("text/plain"),
		),
	}
}

func (c *AppController) GetVersion(writer http.ResponseWriter, _ *http.Request) {
	c.WriteResponse(writer, &ValueObject{Value: c.config.Version}, nil, "text/plain")
}

func (c *AppController) GetHealthz(writer http.ResponseWriter, _ *http.Request) {
	c.WriteResponse(writer, &ValueObject{Value: "ok"}, nil, "text/plain")
}
