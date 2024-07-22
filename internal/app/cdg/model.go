package cdg

import (
	"lca/internal/pkg/events"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ENTITIES = []any{&Collection{}, &Run{}, &SshInfo{}, &GoInfo{}}

type Collection struct {
	Id   uuid.UUID `gorm:"primarykey"`
	Type events.CollectionType

	SSH *SshInfo
	Go  *GoInfo

	Runs []*Run

	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}

type Run struct {
	Id uuid.UUID `gorm:"type:uuid;primary_key;"`

	Collection   Collection
	CollectionId uuid.UUID

	Status events.CollectionStatus
	Error  string

	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}

type SshInfo struct {
	Collection   Collection
	CollectionId uuid.UUID
	Host         string
	Port         int
	User         string
	Password     string
	Script       string
	Path         string
	Timeout      time.Duration
	CreatedAt    time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:true"`
}

func (u *SshInfo) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = "***"
	return
}

type GoInfo struct {
	Collection   Collection
	CollectionId uuid.UUID
	Script       string
	Timeout      time.Duration
	CreatedAt    time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:true"`
}
