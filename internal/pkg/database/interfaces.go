package database

import (
	"gorm.io/gorm"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . DataContext
type DataContext interface {
	Database() *gorm.DB
	RunMigrations(entities ...any) error
}
