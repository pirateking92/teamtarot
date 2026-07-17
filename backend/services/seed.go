package services

import (
	"fmt"

	"github.com/pirateking92/teamtarot/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SeedCards imports the 78-card deck from tarotapi.dev into the cards table.
// It upserts on name_short, so it is safe to run repeatedly. Returns the number
// of cards imported.
func SeedCards(gormDB *gorm.DB) (int, error) {
	cards, err := FetchTarotCardsFromAPI()
	if err != nil {
		return 0, err
	}
	if len(cards) == 0 {
		return 0, fmt.Errorf("tarot API returned no cards")
	}

	// Upsert by name_short: insert new cards, update meanings/description for existing.
	err = gormDB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "name_short"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "arcana", "suit", "meaning_up", "meaning_reverse", "description", "image_file_name",
		}),
	}).Create(&cards).Error
	if err != nil {
		return 0, fmt.Errorf("seeding cards: %w", err)
	}

	var count int64
	if err := gormDB.Model(&models.Card{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
