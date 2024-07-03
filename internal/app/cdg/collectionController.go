package cdg

import (
	"fmt"
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/storage"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/types"
)

type CollectionController struct {
	server.ControllerBase
	config *config.Config
	repo   LogCollectionRepository
	store  storage.Storage
}

func NewCollectionController(config *config.Config, repo LogCollectionRepository, store storage.Storage) *CollectionController {
	return &CollectionController{config: config, repo: repo, store: store}
}

func (c *CollectionController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet,
			"/collection",
			endpoint.Handler(c.List),
			endpoint.Summary("Get collection list"),
			endpoint.Response(http.StatusOK, "Collections", endpoint.SchemaResponseOption([]dto.CollectionDto{})),
			endpoint.Tags("collection"),
			endpoint.Produces("application/json"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
		endpoint.New(
			http.MethodGet,
			"/collection/{id}",
			endpoint.Path("id", types.String, "Id of the Collection", true),
			endpoint.Handler(c.Get),
			endpoint.Summary("Get collection"),
			endpoint.Response(http.StatusOK, "Collection", endpoint.SchemaResponseOption(dto.CollectionDto{})),
			endpoint.Tags("collection"),
			endpoint.Produces("application/json"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
		endpoint.New(
			http.MethodGet,
			"/collection/{id}/archive",
			endpoint.Path("id", types.String, "Id of the Collection", true),
			endpoint.Handler(c.GetFileContent),
			endpoint.Summary("File"),
			endpoint.Response(http.StatusOK, "File contents"),
			endpoint.Produces("application/octet-stream"),
			endpoint.Tags("collection"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
	}
}

func (c *CollectionController) List(writer http.ResponseWriter, _ *http.Request) {
	items, err := c.repo.GetMany(dto.GetCollectionFilterDto{})
	if err != nil {
		c.WriteResponse(writer, nil, err, "application/json")
		return
	}
	c.WriteResponse(writer, items, err, "application/json")
}

func (c *CollectionController) Get(writer http.ResponseWriter, req *http.Request) {
	idPath := strings.Split(req.URL.Path, "/")[2]
	id, err := uuid.Parse(idPath)
	if err != nil {
		c.WriteResponse(writer, nil, err, "application/json")
		return
	}
	item, err := c.repo.Get(id)
	c.WriteResponse(writer, item, err, "application/json")
}

func (c *CollectionController) GetFileContent(writer http.ResponseWriter, req *http.Request) {
	idPath := strings.Split(req.URL.Path, "/")[2]
	id, err := uuid.Parse(idPath)
	if err != nil {
		c.WriteResponse(writer, nil, err, "text/plain")
		return
	}
	bytes, err := c.store.GetFile(id)
	if err != nil {
		c.WriteResponse(writer, nil, err, "text/plain")
		return
	}

	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename="%s.zip"`, id))

	writer.Write(bytes)
}
