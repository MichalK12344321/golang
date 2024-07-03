package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Collection struct {
	Id        uuid.UUID `gorm:"primarykey"`
	Host      string
	Port      int
	User      string
	Password  string
	Script    string
	Status    string
	Error     string
	Path      string
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}

func (u *Collection) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = "***"
	return
}
