package storage

import (
	"io/fs"
	"lca/internal/config"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mholt/archiver/v3"
)

const ext = ".zip"

type DiskStorage struct {
	rootPath string
}

func NewDiskStorage(config *config.Config) *DiskStorage {
	path, err := filepath.Abs(config.App.DataPath)
	if err != nil {
		panic(err)
	}

	return &DiskStorage{rootPath: path}
}

func (d *DiskStorage) GetFile(id uuid.UUID) ([]byte, error) {
	pathToFile := filepath.Join(d.rootPath, id.String()+ext)
	return os.ReadFile(pathToFile)
}

func (d *DiskStorage) CreateFile(id uuid.UUID, stdout string, stderr string) (string, error) {
	var err error
	archivePath := filepath.Join(d.rootPath, id.String()+ext)
	tempPath := filepath.Join(d.rootPath, id.String())
	err = os.Mkdir(tempPath, os.ModeTemporary)
	defer os.RemoveAll(tempPath)
	if err != nil {
		return "", err
	}

	stdoutPath := filepath.Join(tempPath, "stdout.log")

	err = os.WriteFile(stdoutPath, []byte(stdout), fs.FileMode(os.O_CREATE))
	if err != nil {
		return "", err
	}
	stderrPath := filepath.Join(tempPath, "stderr.log")
	err = os.WriteFile(stderrPath, []byte(stderr), fs.FileMode(os.O_CREATE))
	if err != nil {
		return "", err
	}
	err = archiver.Archive([]string{stdoutPath, stderrPath}, archivePath)
	return archivePath, err
}

func (d *DiskStorage) ListFiles(id uuid.UUID) ([]string, error) {
	path := filepath.Join(d.rootPath, id.String()+ext)
	result := []string{}
	err := archiver.Walk(path, func(f archiver.File) error {
		result = append(result, f.Name())
		return nil
	})
	return result, err
}
