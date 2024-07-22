package storage

import (
	"lca/internal/config"
	"lca/internal/pkg/logging"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mholt/archiver/v3"
)

const EXTENSION = ".zip"

type DiskStorage struct {
	rootPath string
	logger   logging.Logger
}

func NewDiskStorage(config *config.Config, logger logging.Logger) *DiskStorage {
	path, err := filepath.Abs(config.App.DataPath)
	if err != nil {
		panic(err)
	}

	return &DiskStorage{rootPath: path, logger: logger}
}

func (d *DiskStorage) GetFile(runId uuid.UUID) ([]byte, error) {
	pathToFile := filepath.Join(d.rootPath, runId.String()+EXTENSION)
	return os.ReadFile(pathToFile)
}

func (d *DiskStorage) ListFiles(runId uuid.UUID) ([]string, error) {
	path := filepath.Join(d.rootPath, runId.String()+EXTENSION)
	result := []string{}
	err := archiver.Walk(path, func(f archiver.File) error {
		result = append(result, f.Name())
		return nil
	})
	return result, err
}

var _ Storage = new(DiskStorage)
