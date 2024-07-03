package events

import (
	"encoding/json"
	"fmt"
	"lca/internal/pkg/util"

	"github.com/google/uuid"
)

type SSHCredentials struct {
	User     string `json:"user" required:"" example:"root"`
	Password string `json:"password" required:"" example:"root"`
}

type CollectionScheduleEvent struct {
	Id     uuid.UUID `json:"id,string" required:"" example:"0ccd924e-751e-4e1c-b51c-d3956d0cf2a7"`
	Host   string    `json:"host" required:"" example:"ne"`
	Port   int       `json:"port" required:"" example:"22"`
	Script string    `json:"script" required:"" example:"/scripts/hello.sh"`
	SSHCredentials
}

type CollectionJobUpdatedEvent struct {
	CollectionScheduleEvent
	Status CollectionStatus `json:"status" enum:"initialized,started,failure,success" required:"" example:"started"`
	Error  string           `json:"error" required:"" example:"failed to run"`
	Path   string           `json:"path" required:"" example:"/var/data/123.zip"`
}

type CollectionStatus string

const CollectionStatusInitialized CollectionStatus = "initialized"
const CollectionStatusStarted CollectionStatus = "started"
const CollectionStatusFailure CollectionStatus = "failure"
const CollectionStatusSuccess CollectionStatus = "success"

func FromJson(event any, body []byte, eventType string) error {
	if eventType != util.GetSimpleName(event) {
		return fmt.Errorf("provided struct %s does not match event type %s", event, eventType)
	}
	return json.Unmarshal(body, event)
}

func (e *CollectionScheduleEvent) GetCredentials() (string, string) {
	return e.User, e.Password
}

func (e *CollectionScheduleEvent) GetAddress() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}
