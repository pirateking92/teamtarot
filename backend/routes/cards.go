package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func setupCardRoutes(baseRouter *gin.RouterGroup) {
	posts := baseRouter.Group("/cards")

	posts.GET("", controllers.GetThreeCards)

}
