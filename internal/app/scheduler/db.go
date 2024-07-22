package scheduler

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

func InsertEvent(context database.DataContext, event any, id uuid.UUID) error {
	db := context.Database()

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
