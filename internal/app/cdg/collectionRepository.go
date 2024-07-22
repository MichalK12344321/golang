package cdg

import (
	"lca/internal/pkg/database"
	"lca/internal/pkg/dto"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type CollectionRepository struct {
	context database.DataContext
}

func NewCollectionRepository(context database.DataContext) *CollectionRepository {
	return &CollectionRepository{context: context}
}

func (repo *CollectionRepository) Insert(dto *dto.CollectionDto) error {
	db := repo.context.Database()

	collectionDb := mapCollectionToModel(dto)

	err := db.Transaction(func(tx *gorm.DB) error {

		sshInfoDb := collectionDb.SSH
		collectionDb.SSH = nil

		goInfoDb := collectionDb.Go
		collectionDb.Go = nil

		create := tx.Create(collectionDb)
		if create.Error != nil {
			return create.Error
		}

		if sshInfoDb != nil {
			sshInfoDb.Collection = *collectionDb
			sshInfoDb.CollectionId = collectionDb.Id
			create = tx.Create(sshInfoDb)
			if create.Error != nil {
				return create.Error
			}
		}

		if goInfoDb != nil {
			goInfoDb.Collection = *collectionDb
			goInfoDb.CollectionId = collectionDb.Id
			create = tx.Create(goInfoDb)
			if create.Error != nil {
				return create.Error
			}
		}

		return nil
	})
	return err
}

func (repo *CollectionRepository) Update(dto *dto.RunDto) error {
	db := repo.context.Database()
	err := db.Transaction(func(tx *gorm.DB) error {
		dbModel := mapCollectionRunToModel(dto)
		save := db.Save(dbModel)
		return save.Error
	})
	return err
}

func (repo *CollectionRepository) Get(id uuid.UUID) (*dto.CollectionDto, error) {
	result := &Collection{}
	db := repo.context.Database()
	fetch := db.Model(&Collection{}).Preload("SSH").Preload("Go").Preload("Runs").Where("id = ?", id.String()).Find(result)
	return mapCollectionFromDto(result, 0), fetch.Error
}

func (repo *CollectionRepository) GetMany(filter *dto.GetCollectionFilterDto) ([]*dto.CollectionDto, error) {
	result := []*Collection{}
	db := repo.context.Database()
	query := db.Model(&Collection{}).Limit(filter.Limit).Preload("SSH").Preload("Go").Preload("Runs")

	if filter.Cursor != uuid.Nil {
		query = query.Where("id <= ?", filter.Cursor)
	}

	if len(filter.Statuses) > 0 {
		query = query.Joins("left join runs on collections.id = runs.collection_id").Where("runs.status IN ?", filter.Statuses)
	}

	fetch := query.Order("created_at DESC").Find(&result)
	if fetch.Error != nil {
		return nil, fetch.Error
	}
	mapped := lo.Map(result, mapCollectionFromDto)
	return mapped, nil
}

func (repo *CollectionRepository) GetRun(id uuid.UUID) (*dto.RunDto, error) {
	result := &Run{}
	db := repo.context.Database()
	tx := db.Model(result).Where("id = ?", id.String()).First(result)

	return mapCollectionRunFromModel(result, 0), tx.Error
}
