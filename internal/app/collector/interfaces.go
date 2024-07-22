package collector

import (
	"context"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"

	"github.com/google/uuid"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Job
type Job interface {
	Id() uuid.UUID
	Run() []error
	Terminate() error
	JobData() JobData
}

//counterfeiter:generate . JobManager
type JobManager interface {
	Start(payload broker.EventPayload) error
	Terminate(id uuid.UUID) error
}

//counterfeiter:generate . JobData
type JobData interface {
	StdoutChannel() chan []byte
	StderrChannel() chan []byte
	ErrorChannel() chan error
	DoneChannel() chan any
}

//counterfeiter:generate . CollectionSink
type CollectionSink interface {
	Initialize() error
	AppendStdout(data []byte)
	AppendStderr(data []byte)
	Finalize() error
}

//counterfeiter:generate . CollectionSinkFactory
type CollectionSinkFactory interface {
	New(ctx context.Context, event *events.RunCreateEvent, job JobData) []CollectionSink
}

type CollectionRunnerOptions any

//counterfeiter:generate . CollectionRunner
type CollectionRunner interface {
	Run(context.Context, JobData, CollectionRunnerOptions)
}

//counterfeiter:generate . Repository
type Repository interface {
	SaveEvent(event any, id uuid.UUID) error
	GetEvent(event any, id uuid.UUID) (*Event, error)
}
