package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

func GetThreeCards(ctx *gin.Context) {
	deck, _ := services.GetAllCards() //returns a type of []Card
	// if err != nil {
	// 	SendInternalError(ctx, err)
	// 	return
	// } ADD THIS BACK IN WHEN ERROR CONTROLLER EXISTS

	var threeCards []models.Card
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))

	// from here below we convert the three Card into three JSONCard

	var jsonCards []models.JSONCard

	for _, card := range threeCards {
		jsonCards = append(jsonCards, models.JSONCard{
			CardName:       card.CardName,
			Type:           card.Type,
			MeaningUp:      card.MeaningUp,
			MeaningReverse: card.MeaningReverse,
			Description:    card.Description,
			ImageName:      card.ShortName + ".png",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"cards": jsonCards})
}

func GetRandomCard(deck []models.Card, currentCards []models.Card) models.Card {
	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randomiser.Intn(len(deck))
	randomCard := deck[randomIndex]
	return randomCard
}
