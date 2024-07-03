package database

import (
	"lca/internal/pkg/dto"
	"lca/internal/pkg/events"
	"lca/internal/pkg/util"
	"reflect"

	"github.com/google/uuid"
)

type CollectionRepository struct {
	context DataContext
}

func NewCollectionRepository(context DataContext) *CollectionRepository {
	return &CollectionRepository{context: context}
}

func (repo *CollectionRepository) Insert(dto *dto.CollectionDto) error {
	db := repo.context.Database()
	model := mapToModel(dto)
	result := db.Create(model)
	return result.Error
}

func (repo *CollectionRepository) Update(dto *dto.CollectionDto) error {
	db := repo.context.Database()
	model := mapToModel(dto)
	result := db.Omit("created_at").Save(model)
	return result.Error
}

func (repo *CollectionRepository) Get(id uuid.UUID) (*dto.CollectionDto, error) {
	result := &Collection{}
	db := repo.context.Database()
	fetch := db.Model(&Collection{}).Where("id = ?", id.String()).Find(result)
	return mapFromModel(result), fetch.Error
}

func (repo *CollectionRepository) GetMany(filter dto.GetCollectionFilterDto) ([]*dto.CollectionDto, error) {
	result := []*Collection{}
	db := repo.context.Database()
	query := db.Model(&Collection{})
	fetch := query.Order("created_at DESC").Find(&result)
	if fetch.Error != nil {
		return nil, fetch.Error
	}
	mapped, err := util.ParallelMap(result, mapFromModel, reflect.TypeOf(&dto.CollectionDto{}))
	return (mapped).([]*dto.CollectionDto), err
}

func mapFromModel(c *Collection) *dto.CollectionDto {
	return &dto.CollectionDto{
		CollectionScheduleEvent: events.CollectionScheduleEvent{
			Id:     c.Id,
			Host:   c.Host,
			Port:   c.Port,
			Script: c.Script,
			SSHCredentials: events.SSHCredentials{
				User:     c.User,
				Password: c.Password,
			},
		},
		Error:  c.Error,
		Path:   c.Path,
		Status: events.CollectionStatus(c.Status),
	}
}

func mapToModel(c *dto.CollectionDto) *Collection {
	return &Collection{
		Id:       c.Id,
		Host:     c.Host,
		Port:     c.Port,
		Script:   c.Script,
		User:     c.User,
		Password: c.Password,
		Error:    c.Error,
		Path:     c.Path,
		Status:   string(c.Status),
	}
}
