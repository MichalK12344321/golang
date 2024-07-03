package collector

import (
	"lca/api/server"
	"lca/internal/config"
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

type VersionController struct {
	server.ControllerBase
	config *config.Config
}

func NewVersionController(config *config.Config) *VersionController {
	return &VersionController{config: config}
}

func (c *VersionController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet,
			"/version",
			endpoint.Handler(c.Get),
			endpoint.Summary("Get application version"),
			endpoint.Response(http.StatusOK, "Version number"),
			endpoint.Produces("text/plain"),
		),
	}
}

func (c *VersionController) Get(writer http.ResponseWriter, _ *http.Request) {
	c.WriteResponse(writer, c.config.Version, nil, "text/plain")
}
