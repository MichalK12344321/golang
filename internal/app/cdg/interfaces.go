package cdg

import (
	"lca/internal/pkg/dto"

	"github.com/google/uuid"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . LogCollectionRepository
type LogCollectionRepository interface {
	Insert(*dto.CollectionDto) error
	Update(*dto.RunDto) error
	Get(uuid.UUID) (*dto.CollectionDto, error)
	GetMany(*dto.GetCollectionFilterDto) ([]*dto.CollectionDto, error)
	GetRun(uuid.UUID) (*dto.RunDto, error)
}

var _ LogCollectionRepository = new(CollectionRepository)
