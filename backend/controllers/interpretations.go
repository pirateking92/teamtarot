package controllers

import (
	//"math/rand"
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/services"
)

// This function takes the AI-powered response provided by GetAIResponse(), converts that data into JSON so it can be sent to the client.

func interpretCards(ctx *gin.Context) {
	AIText, err := services.GetAIResponse("")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch interpretation"})
		return
	}

	var myInterpretation models.AIResponse
	AIText += myInterpretation.Response

	JSONAIResponse := models.JSONAIResponse{JSONResponse: AIText}
	ctx.JSON(http.StatusOK, gin.H{"response": JSONAIResponse})

}
