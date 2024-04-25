package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/services"
)

// Local storage for interpretations and UUIDs
var localStorage map[string]string = make(map[string]string)

// Function to select 3 random tarot cards from the deck
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

// Function to get and interpret 3 tarot cards
func GetandInterpretThreeCards(ctx *gin.Context) {
	deck, err := services.FetchTarotCards() //returns a type of []Card
	if err != nil {
		SendInternalError(ctx, err)
		return
	}
	requestID := uuid.New()
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
			Reversed:       ReverseRandomiser(),
		})

		var reversedValue string
		if card.Reversed {
			reversedValue = "notReversed"
		} else {
			reversedValue = "reversed"
		}

		cardNames = append(cardNames, card.CardName, reversedValue)

		// reversedValue := strconv.FormatBool(card.Reversed)
		// cardNames = append(cardNames, card.CardName, reversedValue)
	}

	//here we send our three Cards and the requestID in JSON form to the client, to be rendered in the UI.

	ctx.JSON(http.StatusOK, gin.H{"cards": jsonCards, "requestID": requestID})

	// here we use Open AI's API to generate a reading of our three cards, we store this reading locally to return it to the user later.
	go func() {
		apiKey := os.Getenv("API_KEY")
		interpretation, err := services.InterpretTarotCards(apiKey, cardNames, requestID)
		if err != nil {
			SendInternalError(ctx, err)
			return
		}
		localStorage[requestID.String()] = interpretation
		GetInterpretation(ctx)
		fmt.Println(interpretation)
	}()

}

// function to send the interpretation from the internal storage to the frontend
func GetInterpretation(ctx *gin.Context) {
	// Get the UUID from the request parameters
	requestID := ctx.Param("uuid")

	// Retrieve the interpretation from local storage
	interpretation, ok := localStorage[requestID]

	// Check if the interpretation was found
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No interpretation found for this UUID"})
		return
	}

	// Return the interpretation
	ctx.JSON(http.StatusOK, gin.H{"interpretation": interpretation})
}

// function to generate reversed cards at random
func ReverseRandomiser() bool {
	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomBool := randomiser.Intn(2)
	return randomBool == 0
}
