package scheduler

import (
	"io/ioutil"
	"lca/api/server"
	"lca/internal/config"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/util"
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

type SchedulerController struct {
	server.ControllerBase
	config    *config.Config
	publisher *CollectionRequestPublisher
}

func NewSchedulerController(config *config.Config, publisher *CollectionRequestPublisher) *SchedulerController {
	return &SchedulerController{config: config, publisher: publisher}
}

func (c *SchedulerController) GetEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodPost,
			"/schedule",
			endpoint.Handler(c.Post),
			endpoint.Body(dto.ScheduleRequestDto{}, "Schedule request", true),
			endpoint.Summary("Schedule log collection"),
			endpoint.Response(http.StatusOK, "Collection Id", endpoint.SchemaResponseOption(dto.ScheduleResponseDto{})),
			endpoint.Tags("schedule"),
			endpoint.Produces("application/json"),
			endpoint.Response(http.StatusBadRequest, "Failure response"),
		),
	}
}

func (c *SchedulerController) Post(writer http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.WriteResponse(writer, "Error reading request body", err, "application/json")
		return
	}
	defer r.Body.Close()

	schedule := &dto.ScheduleRequestDto{}
	err = util.FromJson(schedule, bodyBytes)
	if err != nil {
		c.WriteResponse(writer, "Error parsing body", err, "application/json")
		return
	}
	scheduleEvent := schedule.ToEvent()
	err = c.publisher.Publish(scheduleEvent)
	c.WriteResponse(writer, &dto.ScheduleResponseDto{Id: scheduleEvent.Id}, err, "application/json")
}
