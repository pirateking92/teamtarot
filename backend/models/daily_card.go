package models

// DailyCard is the shared "card of the day". Date is the primary key
// (YYYY-MM-DD), so everyone drawing on the same day sees the same card.
type DailyCard struct {
	Date     string `gorm:"primaryKey;size:10" json:"date"` // YYYY-MM-DD
	CardID   uint   `json:"-"`
	Card     Card   `json:"card"`
	Reversed bool   `json:"reversed"`
}
