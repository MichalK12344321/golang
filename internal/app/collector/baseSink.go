package collector

import (
	"context"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
)

type BaseSink struct {
	logger  logging.Logger
	jobData JobData
	ctx     context.Context
	event   *events.RunCreateEvent
}
