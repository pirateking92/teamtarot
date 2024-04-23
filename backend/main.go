package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/routes"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	// if the above line is not called in main.go the App will always run in debug mode
	// and detailed errors will be shown to the user.
	app := setupApp()
	app.Run(":8082")
}

func setupApp() *gin.Engine {
	app := gin.Default()
	setupCORS(app)
	routes.SetupRoutes(app)
	return app
}

func setupCORS(app *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept"}

	app.Use(cors.New(config))
}
