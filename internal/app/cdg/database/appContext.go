package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppContext struct {
	db *gorm.DB
}

func NewAppContext(conn string) (*AppContext, error) {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &AppContext{db: db}, nil
}

func (appContext *AppContext) Database() *gorm.DB {
	return appContext.db
}
