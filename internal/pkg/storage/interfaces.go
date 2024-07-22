package storage

import (
	"github.com/google/uuid"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Storage
type Storage interface {
	GetFile(runId uuid.UUID) ([]byte, error)
	ListFiles(runId uuid.UUID) ([]string, error)
}
