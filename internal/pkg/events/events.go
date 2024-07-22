package events

import (
	"encoding/json"
	"fmt"
	"lca/internal/pkg/util"
	"time"

	"github.com/google/uuid"
)

type AppendLineEvent struct {
	RunId uuid.UUID `json:"runId,string"`
	File  string    `json:"file"`
	Line  string    `json:"line"`
}

type SSHInfo struct {
	User     string        `json:"user"`
	Password string        `json:"password"`
	Host     string        `json:"host"`
	Port     int           `json:"port"`
	Script   string        `json:"script"`
	Timeout  time.Duration `json:"timeout"`
}

type GoInfo struct {
	Script  string        `json:"script"`
	Timeout time.Duration `json:"timeout"`
}

type CollectionCreateEvent struct {
	CollectionId uuid.UUID      `json:"collectionId,string"`
	Type         CollectionType `json:"type"`
	SSH          *SSHInfo       `json:"ssh"`
	Go           *GoInfo        `json:"go"`
}

type RunCreateEvent struct {
	RunId        uuid.UUID `json:"runId,string"`
	CollectionId uuid.UUID `json:"collectionId,string"`
}

type RunUpdateEvent struct {
	RunCreateEvent
	Status CollectionStatus `json:"status"`
	Error  string           `json:"error"`
}

type RunTerminateEvent struct {
	RunId uuid.UUID `json:"runId,string"`
}

type CollectionType string

const CollectionTypeSSH CollectionType = "ssh"
const CollectionTypeGO CollectionType = "go"

type CollectionStatus string

const CollectionStatusCreated CollectionStatus = "created"
const CollectionStatusStarted CollectionStatus = "started"
const CollectionStatusFailure CollectionStatus = "failure"
const CollectionStatusSuccess CollectionStatus = "success"
const CollectionStatusTerminating CollectionStatus = "terminating"
const CollectionStatusTerminated CollectionStatus = "terminated"

func FromJson(event any, body []byte, eventType string) error {
	eventTargetType := util.GetSimpleName(event)
	if eventType != eventTargetType {
		return fmt.Errorf(
			"provided struct %s (%s) does not match event type %s",
			event,
			eventTargetType,
			eventType,
		)
	}
	return json.Unmarshal(body, event)
}

func (e *SSHInfo) GetCredentials() (string, string) {
	return e.User, e.Password
}

func (e *SSHInfo) GetAddress() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

func (e *SSHInfo) GetTimeout() time.Duration {
	return e.Timeout
}

func (e *SSHInfo) GetCommand() string {
	return e.Script
}

func (c *CollectionCreateEvent) GetTimeout() time.Duration {
	if c.Type == CollectionTypeSSH {
		return c.SSH.Timeout
	}

	return c.Go.Timeout
}
