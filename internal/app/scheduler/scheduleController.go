package scheduler

import (
	"lca/api/server"
	"lca/internal/pkg/dto"
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

type SchedulerController struct {
	server.ControllerBase
	scheduler *Scheduler
}

func NewSchedulerController(s *Scheduler) *SchedulerController {
	return &SchedulerController{scheduler: s}
}

func (c *SchedulerController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodPost,
			"/schedule/ssh",
			endpoint.Handler(c.ScheduleSSh),
			endpoint.Body(dto.ScheduleSSHCollectionDto{}, "Request", true),
			endpoint.Summary("Schedule SSH collection"),
			endpoint.Response(http.StatusOK, "Response", endpoint.SchemaResponseOption(dto.ScheduleResponseDto{})),
			endpoint.Tags("schedule"),
			endpoint.Produces("application/json"),
			endpoint.Consumes("application/json"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
		endpoint.New(
			http.MethodPost,
			"/schedule/go",
			endpoint.Handler(c.ScheduleGo),
			endpoint.Body(dto.ScheduleGoCollectionDto{}, "Request", true),
			endpoint.Summary("Schedule Go collection"),
			endpoint.Response(http.StatusOK, "Response", endpoint.SchemaResponseOption(dto.ScheduleResponseDto{})),
			endpoint.Tags("schedule"),
			endpoint.Produces("application/json"),
			endpoint.Consumes("application/json"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
		endpoint.New(
			http.MethodPost,
			"/terminate",
			endpoint.Handler(c.Terminate),
			endpoint.Body(dto.TerminateRequestDto{}, "Terminate request", true),
			endpoint.Summary("Terminate log collection"),
			endpoint.Response(http.StatusOK, "Response", endpoint.SchemaResponseOption(dto.TerminateResponseDto{})),
			endpoint.Tags("schedule"),
			endpoint.Produces("application/json"),
			endpoint.Consumes("application/json"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
	}
}

func (c *SchedulerController) ScheduleSSh(writer http.ResponseWriter, request *http.Request) {
	schedule := &dto.ScheduleSSHCollectionDto{}
	if c.ParseBody(writer, request, schedule) != nil {
		return
	}
	result, err := c.scheduler.ScheduleSSH(schedule)
	c.WriteResponse(writer, result, err, "application/json")
}

func (c *SchedulerController) ScheduleGo(writer http.ResponseWriter, request *http.Request) {
	schedule := &dto.ScheduleGoCollectionDto{}
	if c.ParseBody(writer, request, schedule) != nil {
		return
	}
	result, err := c.scheduler.ScheduleGo(schedule)
	c.WriteResponse(writer, result, err, "application/json")
}

func (c *SchedulerController) Terminate(writer http.ResponseWriter, request *http.Request) {
	terminate := &dto.TerminateRequestDto{}
	if c.ParseBody(writer, request, terminate) != nil {
		return
	}
	err := c.scheduler.Terminate(*terminate)
	c.WriteResponse(writer, &dto.TerminateResponseDto{}, err, "application/json")
}
