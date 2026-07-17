package services

import (
	"math/rand"

	"github.com/pirateking92/teamtarot/backend/models"
	"gorm.io/gorm"
)

// DrawCards returns n distinct cards drawn at random from the deck.
func DrawCards(gormDB *gorm.DB, n int) ([]models.Card, error) {
	var cards []models.Card
	err := gormDB.Order("RANDOM()").Limit(n).Find(&cards).Error
	return cards, err
}

// RandomReversed returns a random orientation for a drawn card.
func RandomReversed() bool {
	return rand.Intn(2) == 0
}
