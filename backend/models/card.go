package models

// Card is one of the 78 tarot cards, seeded once from tarotapi.dev into Postgres.
// JSON tags define the shape returned to the frontend (camelCase / clean names).
type Card struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
	NameShort      string `gorm:"uniqueIndex;size:16" json:"nameShort"`
	Arcana         string `gorm:"size:8" json:"arcana"` // "major" | "minor"
	Suit           string `gorm:"size:16" json:"suit"`  // wands|cups|swords|pentacles, or "" for major arcana
	MeaningUp      string `gorm:"type:text" json:"meaningUp"`
	MeaningReverse string `gorm:"type:text" json:"meaningReverse"`
	Description    string `gorm:"type:text" json:"description"`
	ImageFileName  string `json:"imageFileName"` // "<nameShort>.jpg"
}
