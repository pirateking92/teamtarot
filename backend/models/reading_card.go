package models

import "github.com/google/uuid"

// ReadingCard links a Reading to one drawn Card, recording its spread position
// and whether it landed reversed.
type ReadingCard struct {
	ID            uint      `gorm:"primaryKey" json:"-"`
	ReadingID     uuid.UUID `gorm:"type:uuid;index" json:"-"`
	CardID        uint      `json:"-"`
	Card          Card      `json:"card"`
	PositionIndex int       `json:"positionIndex"`
	PositionLabel string    `json:"positionLabel"`
	Reversed      bool      `json:"reversed"`
}
