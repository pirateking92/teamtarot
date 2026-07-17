// Package db handles the PostgreSQL connection and schema migration via GORM.
package db

import (
	"github.com/pirateking92/teamtarot/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens a GORM connection to Postgres using the given DSN/URL.
func Connect(databaseURL string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
}

// AutoMigrate creates/updates the schema for all models. For this project's
// scale, GORM AutoMigrate is sufficient; versioned migrations are noted as a
// future hardening step.
func AutoMigrate(gormDB *gorm.DB) error {
	return gormDB.AutoMigrate(
		&models.Card{},
		&models.User{},
		&models.Session{},
		&models.Reading{},
		&models.ReadingCard{},
		&models.DailyCard{},
	)
}
