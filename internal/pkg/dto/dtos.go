package dto

import (
	"lca/internal/pkg/events"

	"github.com/google/uuid"
)

type SSHCredentialsDto events.SSHCredentials

type CollectionDto events.CollectionJobUpdatedEvent

type GetCollectionFilterDto struct{}

type CollectionRequestDto events.CollectionScheduleEvent

type ScheduleRequestDto struct {
	Host     string `json:"host" required:"" example:"ne"`
	Port     int    `json:"port" required:"" example:"22"`
	User     string `json:"user" required:"" example:"root"`
	Password string `json:"password" required:"" example:"root"`
	Script   string `json:"script" required:"" example:"/scripts/hello.sh"`
}

type ScheduleResponseDto struct {
	Id uuid.UUID `json:"id,string" required:"" example:"a0a73fd7-1161-4b5b-bce7-01907c1353c9"`
}

func (schedule *ScheduleRequestDto) ToEvent() *events.CollectionScheduleEvent {
	return &events.CollectionScheduleEvent{
		Id:     uuid.New(),
		Host:   schedule.Host,
		Port:   schedule.Port,
		Script: schedule.Script,
		SSHCredentials: events.SSHCredentials{
			User:     schedule.User,
			Password: schedule.Password,
		},
	}
}
