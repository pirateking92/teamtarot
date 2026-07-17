package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is a registered account. Guests never get a User row — readings simply
// carry a null UserID. Accounts are optional and only unlock saved history.
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;size:320" json:"email"`
	PasswordHash string    `json:"-"`
	DisplayName  string    `json:"displayName"`
	CreatedAt    time.Time `json:"createdAt"`
}

// BeforeCreate assigns a UUID if one is not already set.
func (u *User) BeforeCreate(*gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
