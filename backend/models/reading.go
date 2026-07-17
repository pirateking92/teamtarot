package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reading status values.
const (
	ReadingPending  = "pending"
	ReadingComplete = "complete"
	ReadingError    = "error"
)

// Reading is a single tarot reading. UserID is nil for guest readings.
// The AI interpretation is generated asynchronously and written back here,
// so it survives restarts and can appear in a user's history.
type Reading struct {
	ID             uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	UserID         *uuid.UUID    `gorm:"type:uuid;index" json:"-"`
	QuerentName    string        `json:"querentName"`
	Story          string        `gorm:"type:text" json:"story"`
	SpreadType     string        `gorm:"size:32" json:"spreadType"`
	Status         string        `gorm:"size:16" json:"status"`
	Interpretation string        `gorm:"type:text" json:"interpretation"`
	CreatedAt      time.Time     `json:"createdAt"`
	Cards          []ReadingCard `gorm:"constraint:OnDelete:CASCADE" json:"cards"`
}

// BeforeCreate assigns a UUID if one is not already set.
func (r *Reading) BeforeCreate(*gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
