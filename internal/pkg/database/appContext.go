package database

import (
	"lca/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppContext struct {
	db *gorm.DB
}

func newAppContextConn(conn string) (*AppContext, error) {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &AppContext{db: db}, nil
}

func (appContext *AppContext) Database() *gorm.DB {
	return appContext.db
}

func (appContext *AppContext) RunMigrations(v ...any) error {
	return appContext.db.Migrator().AutoMigrate(
		v...,
	)
}

func NewContext(config *config.Config) (DataContext, error) {
	return newAppContextConn(config.DB.CONNECTION)
}
