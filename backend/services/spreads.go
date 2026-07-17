package services

// SpreadPosition is one slot in a spread, with a human-readable label the
// interpretation and UI use (e.g. "Past", "Present", "Future").
type SpreadPosition struct {
	Index int    `json:"index"`
	Label string `json:"label"`
}

// Spread is a named tarot layout. Spreads are static config, not DB rows.
type Spread struct {
	Type        string           `json:"type"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Positions   []SpreadPosition `json:"positions"`
}

// DefaultSpreadType is used when a reading request omits a spread.
const DefaultSpreadType = "past-present-future"

func pos(labels ...string) []SpreadPosition {
	out := make([]SpreadPosition, len(labels))
	for i, l := range labels {
		out[i] = SpreadPosition{Index: i, Label: l}
	}
	return out
}

// spreadOrder preserves a stable, sensible order for AllSpreads().
var spreadOrder = []string{"single", "past-present-future", "five-card", "celtic-cross"}

var spreads = map[string]Spread{
	"single": {
		Type:        "single",
		Name:        "Single Card",
		Description: "One card for a focused question or a quick daily nudge.",
		Positions:   pos("The Card"),
	},
	"past-present-future": {
		Type:        "past-present-future",
		Name:        "Past · Present · Future",
		Description: "The classic three-card spread tracing your path through time.",
		Positions:   pos("Past", "Present", "Future"),
	},
	"five-card": {
		Type:        "five-card",
		Name:        "Five Card Cross",
		Description: "A fuller picture: the situation, what challenges it, its roots, recent past, and the likely outcome.",
		Positions:   pos("The Situation", "The Challenge", "The Foundation", "The Recent Past", "The Outcome"),
	},
	"celtic-cross": {
		Type:        "celtic-cross",
		Name:        "Celtic Cross",
		Description: "The deep ten-card reading for a thorough look at a complex matter.",
		Positions: pos(
			"The Present", "The Challenge", "The Foundation", "The Recent Past",
			"The Crown", "The Near Future", "Your Attitude", "External Influences",
			"Hopes & Fears", "The Outcome",
		),
	},
}

// AllSpreads returns every spread in a stable display order.
func AllSpreads() []Spread {
	out := make([]Spread, 0, len(spreadOrder))
	for _, t := range spreadOrder {
		out = append(out, spreads[t])
	}
	return out
}

// GetSpread looks up a spread by type. Returns ok=false if unknown.
func GetSpread(spreadType string) (Spread, bool) {
	s, ok := spreads[spreadType]
	return s, ok
}
