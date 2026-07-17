package controllers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pirateking92/teamtarot/backend/models"
	"github.com/pirateking92/teamtarot/backend/services"
	"gorm.io/gorm"
)

type createReadingRequest struct {
	QuerentName string `json:"querentName"`
	Story       string `json:"story"`
	SpreadType  string `json:"spreadType"`
}

// CreateReading draws cards for the chosen spread, persists the reading, and
// kicks off asynchronous interpretation.
// POST /api/readings
func (h *Handler) CreateReading(c *gin.Context) {
	var req createReadingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	spreadType := req.SpreadType
	if spreadType == "" {
		spreadType = services.DefaultSpreadType
	}
	spread, ok := services.GetSpread(spreadType)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unknown spread type"})
		return
	}

	n := len(spread.Positions)
	cards, err := services.DrawCards(h.DB, n)
	if err != nil || len(cards) < n {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not draw cards (is the deck seeded?)"})
		return
	}

	reading := models.Reading{
		UserID:      currentUserID(c),
		QuerentName: req.QuerentName,
		Story:       req.Story,
		SpreadType:  spreadType,
		Status:      models.ReadingPending,
	}
	reading.Cards = make([]models.ReadingCard, n)
	interpretInput := services.InterpretInput{
		QuerentName: req.QuerentName,
		Story:       req.Story,
		SpreadName:  spread.Name,
	}
	for i, card := range cards {
		reversed := services.RandomReversed()
		reading.Cards[i] = models.ReadingCard{
			CardID:        card.ID,
			PositionIndex: i,
			PositionLabel: spread.Positions[i].Label,
			Reversed:      reversed,
		}

		meaning := card.MeaningUp
		name := card.Name
		if reversed {
			meaning = card.MeaningReverse
			name = card.Name + " (Reversed)"
		}
		interpretInput.Cards = append(interpretInput.Cards, services.DrawnCard{
			Position: spread.Positions[i].Label,
			Name:     name,
			Reversed: reversed,
			Meaning:  meaning,
		})
	}

	// Omit the Card association so GORM uses the CardID FK and never writes to the cards table.
	if err := h.DB.Omit("Cards.Card").Create(&reading).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save reading"})
		return
	}

	go h.generateInterpretation(reading.ID, interpretInput)

	out, err := h.loadReading(reading.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load reading"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"reading": out, "spread": spread})
}

// GetReading returns a reading by id, including its cards, status, and
// interpretation (which may still be pending). This is the poll target.
// GET /api/readings/:id
func (h *Handler) GetReading(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reading id"})
		return
	}

	reading, err := h.loadReading(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "reading not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load reading"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reading": reading})
}

// ListReadings returns the authenticated user's reading history (newest first).
// GET /api/readings  (requires auth)
func (h *Handler) ListReadings(c *gin.Context) {
	uid := currentUserID(c)
	if uid == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	var readings []models.Reading
	err := h.DB.
		Preload("Cards", func(db *gorm.DB) *gorm.DB { return db.Order("reading_cards.position_index") }).
		Preload("Cards.Card").
		Where("user_id = ?", *uid).
		Order("created_at DESC").
		Find(&readings).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"readings": readings})
}

// loadReading fetches a reading with its cards (ordered) and each card's detail.
func (h *Handler) loadReading(id uuid.UUID) (models.Reading, error) {
	var reading models.Reading
	err := h.DB.
		Preload("Cards", func(db *gorm.DB) *gorm.DB { return db.Order("reading_cards.position_index") }).
		Preload("Cards.Card").
		First(&reading, "id = ?", id).Error
	return reading, err
}

// generateInterpretation runs in a goroutine, writing the result back into the
// reading row. Falls back to a stub when TESTING or no API key is configured.
func (h *Handler) generateInterpretation(readingID uuid.UUID, in services.InterpretInput) {
	var text string
	var genErr error

	if h.Cfg.Testing || !h.Cfg.HasAnthropicCredential() {
		text = "This is a test interpretation. Cassandra senses the cards clearly, but her voice is stubbed for now."
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
		defer cancel()
		text, genErr = services.GenerateInterpretation(ctx, in)
	}

	updates := map[string]any{}
	if genErr != nil {
		updates["status"] = models.ReadingError
		updates["interpretation"] = "The cards blur before Cassandra's eyes — the reading could not be completed. (" + genErr.Error() + ")"
	} else {
		updates["status"] = models.ReadingComplete
		updates["interpretation"] = text
	}
	h.DB.Model(&models.Reading{}).Where("id = ?", readingID).Updates(updates)
}

// currentUserID returns the authenticated user's id, or nil for guests.
// The optionalAuth/requireAuth middleware (Phase 3) sets "userID" in context.
func currentUserID(c *gin.Context) *uuid.UUID {
	v, ok := c.Get("userID")
	if !ok {
		return nil
	}
	if id, ok := v.(uuid.UUID); ok {
		return &id
	}
	return nil
}
