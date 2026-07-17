package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pirateking92/teamtarot/backend/db"
	"github.com/pirateking92/teamtarot/backend/env"
	"github.com/pirateking92/teamtarot/backend/routes"
	"github.com/pirateking92/teamtarot/backend/services"
	"gorm.io/gorm"
)

func main() {
	cfg := env.Load()

	gormDB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	if err := db.AutoMigrate(gormDB); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	// Subcommands: `go run . seed` imports the 78 cards from tarotapi.dev.
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		count, err := services.SeedCards(gormDB)
		if err != nil {
			log.Fatalf("seed failed: %v", err)
		}
		log.Printf("seeded %d cards", count)
		return
	}

	if cfg.GinMode != "" {
		gin.SetMode(cfg.GinMode)
	}

	app := setupApp(cfg, gormDB)
	log.Printf("Cassandra backend listening on :%s", cfg.Port)
	if err := app.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func setupApp(cfg env.Config, gormDB *gorm.DB) *gin.Engine {
	app := gin.Default()
	// No CORS middleware: this API is an internal service behind the SvelteKit
	// BFF. The browser never calls it directly — only the SvelteKit server does
	// (server-to-server), so a browser CORS gate is unnecessary and would wrongly
	// reject those calls based on their Origin. Do not expose this API publicly.
	routes.SetupRoutes(app, gormDB, cfg)
	return app
}
