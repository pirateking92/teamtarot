package models

type JSONCard struct {
	CardName       string // API
	Type           string // API
	MeaningUp      string // API
	MeaningReverse string // API
	Description    string // API
	ImageName      string // taken from name_short (API)
}
