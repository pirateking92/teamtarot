package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/services"
)

// ListSpreads returns the available spread types and their positions.
// GET /api/spreads
func (h *Handler) ListSpreads(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"spreads": services.AllSpreads()})
}
