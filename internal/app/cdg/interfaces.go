package cdg

import (
	"lca/internal/app/cdg/database"
	"lca/internal/pkg/dto"

	"github.com/google/uuid"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . LogCollectionRepository
type LogCollectionRepository interface {
	Insert(*dto.CollectionDto) error
	Update(*dto.CollectionDto) error
	Get(uuid.UUID) (*dto.CollectionDto, error)
	GetMany(dto.GetCollectionFilterDto) ([]*dto.CollectionDto, error)
}

var _ LogCollectionRepository = new(database.CollectionRepository)
