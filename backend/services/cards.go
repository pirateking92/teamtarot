// Package services holds business logic: fetching/seeding cards, drawing spreads,
// generating interpretations, and (later) auth.
package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pirateking92/teamtarot/backend/models"
)

const tarotAPIURL = "http://tarotapi.dev/api/v1/cards"

// apiCard mirrors the JSON shape returned by tarotapi.dev. It is only used at
// seed time to map the external payload onto our own models.Card.
type apiCard struct {
	Name       string `json:"name"`
	NameShort  string `json:"name_short"`
	Type       string `json:"type"` // "major" | "minor"
	Suit       string `json:"suit"`
	MeaningUp  string `json:"meaning_up"`
	MeaningRev string `json:"meaning_rev"`
	Desc       string `json:"desc"`
}

// FetchTarotCardsFromAPI pulls the full 78-card deck from tarotapi.dev and maps
// it onto our Card model. This is only called at seed time, never per-reading.
func FetchTarotCardsFromAPI() ([]models.Card, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(tarotAPIURL)
	if err != nil {
		return nil, fmt.Errorf("fetching tarot cards: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tarot API returned status %d", resp.StatusCode)
	}

	var payload struct {
		Cards []apiCard `json:"cards"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decoding tarot API response: %w", err)
	}

	cards := make([]models.Card, 0, len(payload.Cards))
	for _, c := range payload.Cards {
		cards = append(cards, models.Card{
			Name:           c.Name,
			NameShort:      c.NameShort,
			Arcana:         c.Type,
			Suit:           c.Suit,
			MeaningUp:      c.MeaningUp,
			MeaningReverse: c.MeaningRev,
			Description:    c.Desc,
			ImageFileName:  c.NameShort + ".jpg",
		})
	}
	return cards, nil
}
