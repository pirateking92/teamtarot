// Package routes wires HTTP endpoints to controllers under the /api prefix.
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/controllers"
	"github.com/pirateking92/teamtarot/backend/env"
	"github.com/pirateking92/teamtarot/backend/middleware"
	"gorm.io/gorm"
)

// SetupRoutes registers all API routes.
func SetupRoutes(engine *gin.Engine, gormDB *gorm.DB, cfg env.Config) {
	h := controllers.New(gormDB, cfg)

	api := engine.Group("/api")
	// Attach the current user to every request if a valid session is present.
	api.Use(middleware.OptionalAuth(gormDB))

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api.GET("/spreads", h.ListSpreads)
	api.GET("/daily", h.DailyCard)

	cards := api.Group("/cards")
	cards.GET("", h.ListCards)
	cards.GET("/:nameShort", h.GetCard)

	readings := api.Group("/readings")
	readings.POST("", h.CreateReading)                        // guest or logged-in
	readings.GET("", middleware.RequireAuth(), h.ListReadings) // history — auth only
	readings.GET("/:id", h.GetReading)

	auth := api.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/logout", h.Logout)
	auth.GET("/me", middleware.RequireAuth(), h.Me)
}
