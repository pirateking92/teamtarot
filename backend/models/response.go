package models

type JSONCard struct {
	CardName       string `json:"name"`
	Type           string `json:"type"`
	MeaningUp      string `json:"meaning_up"`
	MeaningReverse string `json:"meaning_rev"`
	Description    string `json:"desc"`
	ImageName      string `json:"image_file_name"`
	Reversed       bool   `json:"reversed"`
}
