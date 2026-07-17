package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pirateking92/teamtarot/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Auth errors surfaced to controllers.
var (
	ErrEmailTaken         = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

// SessionDuration is how long a login session stays valid.
const SessionDuration = 30 * 24 * time.Hour

// Register creates a new user and an initial session, returning the session token.
func Register(gormDB *gorm.DB, email, password, displayName string) (models.User, string, error) {
	email = normalizeEmail(email)

	var count int64
	if err := gormDB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return models.User{}, "", err
	}
	if count > 0 {
		return models.User{}, "", ErrEmailTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, "", err
	}

	if displayName == "" {
		displayName = strings.SplitN(email, "@", 2)[0]
	}
	user := models.User{Email: email, PasswordHash: string(hash), DisplayName: displayName}
	if err := gormDB.Create(&user).Error; err != nil {
		return models.User{}, "", err
	}

	token, err := createSession(gormDB, user.ID)
	return user, token, err
}

// Login verifies credentials and creates a new session, returning its token.
func Login(gormDB *gorm.DB, email, password string) (models.User, string, error) {
	email = normalizeEmail(email)

	var user models.User
	err := gormDB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, "", ErrInvalidCredentials
	}
	if err != nil {
		return models.User{}, "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return models.User{}, "", ErrInvalidCredentials
	}

	token, err := createSession(gormDB, user.ID)
	return user, token, err
}

// UserForSession resolves a session token to its (non-expired) user.
func UserForSession(gormDB *gorm.DB, token string) (models.User, error) {
	var session models.Session
	err := gormDB.Where("token = ? AND expires_at > ?", token, time.Now()).First(&session).Error
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	err = gormDB.First(&user, "id = ?", session.UserID).Error
	return user, err
}

// DeleteSession revokes a session (logout).
func DeleteSession(gormDB *gorm.DB, token string) error {
	return gormDB.Where("token = ?", token).Delete(&models.Session{}).Error
}

func createSession(gormDB *gorm.DB, userID uuid.UUID) (string, error) {
	token, err := randomToken()
	if err != nil {
		return "", err
	}
	session := models.Session{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(SessionDuration),
	}
	if err := gormDB.Create(&session).Error; err != nil {
		return "", err
	}
	return token, nil
}

func randomToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
