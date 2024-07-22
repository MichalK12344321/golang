package collector

import (
	"context"
	"lca/internal/config"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/storage"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mholt/archiver/v3"
)

type StorageSink struct {
	BaseSink
	config *config.Config
	id     uuid.UUID
}

func NewStorageSink(config *config.Config, ctx context.Context, event *events.RunCreateEvent, logger logging.Logger, jobData JobData) *StorageSink {
	return &StorageSink{
		BaseSink: BaseSink{
			logger:  logger,
			ctx:     ctx,
			event:   event,
			jobData: jobData,
		},
		config: config,
		id:     event.RunId,
	}
}

func (d *StorageSink) workDir() string {
	return filepath.Join(d.config.DataPath, d.id.String())
}

func (d *StorageSink) Initialize() error {
	workDir := d.workDir()
	err := os.MkdirAll(workDir, 0777)
	if err != nil {
		return err
	}
	d.logger.Debug("job - created temp: '%s'", workDir)

	stdoutPath := storage.GetStdOutPath(workDir)
	stderrPath := storage.GetStdErrPath(workDir)

	err = d.AppendToFile(stdoutPath, []byte{})
	if err != nil {
		return err
	}

	err = d.AppendToFile(stderrPath, []byte{})
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageSink) AppendStdout(data []byte) {
	stdoutPath := storage.GetStdOutPath(s.workDir())
	line := data
	line = append(line, 10)
	err := s.AppendToFile(stdoutPath, line)
	if err != nil {
		s.jobData.ErrorChannel() <- err
	}
}

func (s *StorageSink) AppendStderr(data []byte) {
	stdoutPath := storage.GetStdErrPath(s.workDir())
	line := data
	line = append(line, 10)
	err := s.AppendToFile(stdoutPath, line)
	if err != nil {
		s.jobData.ErrorChannel() <- err
	}
}

func (s *StorageSink) Finalize() error {
	workDir := s.workDir()
	stdoutPath := storage.GetStdOutPath(workDir)
	stderrPath := storage.GetStdErrPath(workDir)
	defer os.RemoveAll(workDir)
	archivePath := workDir + storage.EXTENSION
	err := archiver.Archive([]string{stdoutPath, stderrPath}, archivePath)
	return err
}

func (d *StorageSink) AppendToFile(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	return err
}

var _ CollectionSink = new(StorageSink)
