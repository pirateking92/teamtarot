package services

import (
	"errors"

	"github.com/pirateking92/teamtarot/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetOrCreateDailyCard returns the shared card of the day for the given date
// (YYYY-MM-DD), creating it on first request. Everyone sees the same card.
func GetOrCreateDailyCard(gormDB *gorm.DB, date string) (models.DailyCard, error) {
	var daily models.DailyCard

	err := gormDB.Preload("Card").Where("date = ?", date).First(&daily).Error
	if err == nil {
		return daily, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return daily, err
	}

	// Not yet chosen for today — pick a random card.
	var card models.Card
	if err := gormDB.Order("RANDOM()").First(&card).Error; err != nil {
		return daily, err
	}

	daily = models.DailyCard{Date: date, CardID: card.ID, Reversed: RandomReversed()}
	// DoNothing on conflict handles the race where a concurrent request created it first.
	if err := gormDB.Clauses(clause.OnConflict{DoNothing: true}).Create(&daily).Error; err != nil {
		return daily, err
	}

	// Reload with the associated card (and the winner's row if we lost the race).
	if err := gormDB.Preload("Card").Where("date = ?", date).First(&daily).Error; err != nil {
		return daily, err
	}
	return daily, nil
}
