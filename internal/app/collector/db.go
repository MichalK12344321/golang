package collector

import (
	"lca/internal/pkg/database"
	"lca/internal/pkg/util"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

var ENTITIES = []any{&Event{}}

type Event struct {
	Id      uint `gorm:"primarykey"`
	EventId uuid.UUID
	Type    string
	Raw     pgtype.JSONB `gorm:"type:jsonb;default:'[]';not null"`
}

type DataRepository struct {
	context database.DataContext
}

func NewDataRepository(context database.DataContext) *DataRepository {
	return &DataRepository{context: context}
}

var _ Repository = new(DataRepository)

func (r *DataRepository) SaveEvent(event any, id uuid.UUID) error {
	db := r.context.Database()

	err := db.Transaction(func(tx *gorm.DB) error {

		dbEvent := &Event{
			EventId: id,
			Type:    util.GetSimpleName(event),
		}

		dbEvent.Raw.Set(event)

		create := tx.Create(dbEvent)
		return create.Error
	})
	return err
}

func (r *DataRepository) GetEvent(event any, id uuid.UUID) (*Event, error) {
	db := r.context.Database()
	result := &Event{}
	tx := db.Model(result).Where("event_id = ? AND type = ?", id.String(), util.GetSimpleName(event)).First(result)
	return result, tx.Error
}
