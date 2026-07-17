// Package controllers holds the Gin HTTP handlers. All handlers hang off a
// single Handler struct carrying shared dependencies (DB, config).
package controllers

import (
	"github.com/pirateking92/teamtarot/backend/env"
	"gorm.io/gorm"
)

// Handler carries dependencies shared by all HTTP handlers.
type Handler struct {
	DB  *gorm.DB
	Cfg env.Config
}

// New builds a Handler.
func New(gormDB *gorm.DB, cfg env.Config) *Handler {
	return &Handler{DB: gormDB, Cfg: cfg}
}
