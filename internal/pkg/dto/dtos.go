package dto

import (
	"lca/internal/pkg/events"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CollectionDto struct {
	events.CollectionCreateEvent
	Runs []*RunDto `json:"runs"`
}

type RunScheduleDto events.RunCreateEvent

type RunDto events.RunUpdateEvent

type GetCollectionFilterDto struct {
	Statuses []GetCollectionFilterStatusEnum `json:"statuses" required:"" example:"['success']"`
	Limit    int                             `json:"take"`
	Cursor   uuid.UUID                       `json:"cursor"`
}

type GetCollectionFilterStatusEnum events.CollectionStatus

type ScheduleRequestDto struct {
	Repeat bool   `json:"repeat" required:"" example:"false"`
	Cron   string `json:"cron" required:"" example:"* * * * *" description:"cron for repeat collection - leave empty if one off"`
}

type ScheduleSSHCollectionDto struct {
	ScheduleRequestDto
	Targets []TargetDto `json:"targets" required:""`
	Script  string      `json:"script" required:"" example:"/scripts/hello.sh"`
	Timeout string      `json:"timeout" required:"" example:"1m"`
}

type ScheduleGoCollectionDto struct {
	ScheduleRequestDto
	Script  string `json:"script" required:"" example:"package main; func main(){return}"`
	Timeout string `json:"timeout" required:"" example:"5s"`
}

type TargetDto struct {
	Host     string `json:"host" required:"" example:"192.168.49.2"`
	Port     int    `json:"port" required:"" example:"30022"`
	User     string `json:"user" required:"" example:"root"`
	Password string `json:"password" required:"" example:"root"`
}

type ScheduleResponseDto struct {
	Runs []*RunScheduleDto `json:"runs" required:""`
}

type TerminateRequestDto struct {
	RunId uuid.UUID `json:"runId,string" required:"" example:"a0a73fd7-1161-4b5b-bce7-01907c1353c9"`
}

type TerminateResponseDto struct{}

type AppendLineDto struct {
	RunId uuid.UUID `json:"runId,string" required:"" example:"a0a73fd7-1161-4b5b-bce7-01907c1353c9"`
	File  string    `json:"file" required:"" example:"stdout.log"`
	Line  string    `json:"line" required:"" example:"output text"`
}

func (target *TargetDto) ToEvent(*events.SSHInfo, error) {}

func (schedule *ScheduleSSHCollectionDto) ToEvents() []*events.CollectionCreateEvent {
	timeout, err := time.ParseDuration(schedule.Timeout)
	if err != nil {
		timeout = time.Duration(time.Minute)
	}

	collectionId := uuid.New()

	result := lo.Map(schedule.Targets, func(target TargetDto, _ int) *events.CollectionCreateEvent {
		return &events.CollectionCreateEvent{
			CollectionId: collectionId,
			Type:         events.CollectionTypeSSH,
			SSH: &events.SSHInfo{
				User:     target.User,
				Password: target.Password,
				Host:     target.Host,
				Port:     target.Port,
				Script:   schedule.Script,
				Timeout:  timeout,
			},
		}
	})

	return result
}

func (schedule *ScheduleGoCollectionDto) ToEvents() []*events.CollectionCreateEvent {
	timeout, err := time.ParseDuration(schedule.Timeout)
	if err != nil {
		timeout = time.Duration(time.Minute)
	}

	collectionId := uuid.New()

	result := &events.CollectionCreateEvent{
		CollectionId: collectionId,
		Type:         events.CollectionTypeGO,
		Go: &events.GoInfo{
			Script:  schedule.Script,
			Timeout: timeout,
		},
	}

	return []*events.CollectionCreateEvent{result}
}

func (terminate *TerminateRequestDto) ToEvent() *events.RunTerminateEvent {
	return &events.RunTerminateEvent{
		RunId: terminate.RunId,
	}
}
