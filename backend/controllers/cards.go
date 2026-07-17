package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/models"
	"github.com/pirateking92/teamtarot/backend/services"
	"gorm.io/gorm"
)

// ListCards returns the deck, optionally filtered by ?arcana= and ?suit=.
// GET /api/cards
func (h *Handler) ListCards(c *gin.Context) {
	query := h.DB.Model(&models.Card{})
	if arcana := c.Query("arcana"); arcana != "" {
		query = query.Where("arcana = ?", arcana)
	}
	if suit := c.Query("suit"); suit != "" {
		query = query.Where("suit = ?", suit)
	}

	var cards []models.Card
	if err := query.Order("id").Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load cards"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cards": cards})
}

// GetCard returns a single card by its short name (e.g. "ar00", "swqu").
// GET /api/cards/:nameShort
func (h *Handler) GetCard(c *gin.Context) {
	nameShort := c.Param("nameShort")

	var card models.Card
	err := h.DB.Where("name_short = ?", nameShort).First(&card).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "card not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"card": card})
}

// DailyCard returns the shared card of the day.
// GET /api/daily
func (h *Handler) DailyCard(c *gin.Context) {
	today := time.Now().Format("2006-01-02")
	daily, err := services.GetOrCreateDailyCard(h.DB, today)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load daily card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"daily": daily})
}
