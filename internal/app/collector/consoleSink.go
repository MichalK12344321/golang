package collector

import (
	"lca/internal/pkg/logging"
)

type ConsoleSink struct {
	logger logging.Logger
}

func (c *ConsoleSink) AppendStderr(data []byte) {
	c.logger.Debug("CONSOLE SINK: %s - %s", "APPEND STDERR", data)
}

func (c *ConsoleSink) AppendStdout(data []byte) {
	c.logger.Debug("CONSOLE SINK: %s - %s", "APPEND STDOUT", data)
}

func (c *ConsoleSink) Finalize() error {
	c.logger.Debug("CONSOLE SINK: %s", "FINALIZE")
	return nil
}

func (c *ConsoleSink) Initialize() error {
	c.logger.Debug("CONSOLE SINK: %s", "INITIALIZE")
	return nil
}

func NewConsoleSink(logger logging.Logger) *ConsoleSink {
	return &ConsoleSink{logger: logger}
}

var _ CollectionSink = new(ConsoleSink)
