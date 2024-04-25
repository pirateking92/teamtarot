package controllers

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/services"
)

var localStorage map[string]string = make(map[string]string)

func GetRandomCard(deck []models.Card, currentCards []models.Card) models.Card {
	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		randomIndex := randomiser.Intn(len(deck))
		randomCard := deck[randomIndex]

		isDuplicate := false
		for _, card := range currentCards {
			if card.CardName == randomCard.CardName {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			return randomCard
		}
	}
}

func GetandInterpretThreeCards(ctx *gin.Context) {
	deck, _ := services.FetchTarotCards() //returns a type of []Card
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
	var cardNames []string

	for _, card := range threeCards {
		jsonCards = append(jsonCards, models.JSONCard{
			CardName:       card.CardName,
			Type:           card.Type,
			MeaningUp:      card.MeaningUp,
			MeaningReverse: card.MeaningReverse,
			Description:    card.Description,
			ImageName:      card.ShortName + ".jpg",
		})
		cardNames = append(cardNames, card.CardName)
	}

	// var cardNames []string
	// for _, card := range threeCards {
	// 	cardNames = append(cardNames, card.CardName)
	// }

	// interpret, _ := services.InterpretTarotCards("sk-proj-ciSJ9dZTcV5znq0TrIOnT3BlbkFJcQiiBmTQc31dAoOGaHap", cardNames)
	ctx.JSON(http.StatusOK, gin.H{"cards": jsonCards})

	// Interpret the cards
	go func() {
		apiKey := os.Getenv("API_KEY")
		requestID := uuid.New()
		interpretation, err := services.InterpretTarotCards(apiKey, cardNames, requestID)
		if err != nil {
			SendInternalError(ctx, err)
			return
		}
		localStorage[requestID.String()] = interpretation
	}()
	// ctx.JSON(http.StatusOK, gin.H{"interpret": interpretation, "requestID": requestID})

	// func InterpretThreeCards(ctx *gin.Context) {
	// 	var request struct {
	// 		Cards []string `json:"cards"`
	// 		}

	// 		requestID := uuid.New()
	// 	if err := ctx.BindJSON(&request); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	// 		return
	// 	}
	// 	// apiKey := os.Getenv("API_KEY")
	// 	// requestID = uuid.New()

	// // interpret, _ := services.InterpretTarotCards(apiKey, request.Cards, requestID)
	// ctx.JSON(http.StatusOK, gin.H{"interpret": interpretation, "requestID": requestID})
}
