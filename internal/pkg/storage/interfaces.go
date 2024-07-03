package storage

import (
	"github.com/google/uuid"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Storage
type Storage interface {
	CreateFile(id uuid.UUID, stdout string, stderr string) (string, error)
	GetFile(id uuid.UUID) ([]byte, error)
	ListFiles(id uuid.UUID) ([]string, error)
}
