package models

import (
	"time"

	"github.com/google/uuid"
)

// Session is a server-side auth session. The opaque Token is stored in an
// httpOnly cookie; sessions are revocable by deleting the row.
type Session struct {
	Token     string    `gorm:"primaryKey;size:64" json:"-"`
	UserID    uuid.UUID `gorm:"type:uuid;index" json:"-"`
	ExpiresAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}
