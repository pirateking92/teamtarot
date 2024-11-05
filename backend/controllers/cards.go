package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main.go/errors"
	"main.go/models"
	"main.go/services"
)

// Local storage for interpretations and UUIDs
var LocalStorage map[string]string = make(map[string]string)

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
	deck, err := services.FetchTarotCards()
	if err != nil {
		errors.SendInternalError(ctx, err)
		return
	}

	requestID := uuid.New()
	var threeCards []models.Card
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))

	var jsonCards []models.JSONCard
	var cardNames []string
	for _, card := range threeCards {
		reversed := ReverseRandomiser()
		var FinalCardName string
		if reversed {
			FinalCardName = card.CardName + " (Reversed)"
		} else {
			FinalCardName = card.CardName
		}
		jsonCards = append(jsonCards, models.JSONCard{
			CardName:       FinalCardName,
			Type:           card.Type,
			MeaningUp:      card.MeaningUp,
			MeaningReverse: card.MeaningReverse,
			Description:    card.Description,
			ImageName:      card.ShortName + ".jpg",
			Reversed:       reversed,
		})
		var reversedValue string
		card.Reversed = reversed
		if card.Reversed {
			reversedValue = "(Reversed)"
		} else {
			reversedValue = ""
		}
		cardNames = append(cardNames, card.CardName, reversedValue)
	}

	ctx.JSON(http.StatusOK, gin.H{"cards": jsonCards, "requestID": requestID})
	userStory := ctx.Query("userstory")
	userName := ctx.Query("name")
	fmt.Print(userName, userStory)

	go func() {
		testing := os.Getenv("TESTING")
		if testing == "True" {
			interpretation := "This is a test interpretation"
			LocalStorage[requestID.String()] = interpretation
			return
		}

		// CHANGE 1: Updated environment variable name to match Anthropic's convention
		apiKey := os.Getenv("ANTHROPIC_API_KEY")
		if apiKey == "" {
			errors.SendInternalError(ctx, fmt.Errorf("ANTHROPIC_API_KEY not found in environment"))
			return
		}

		// CHANGE 2: Added error checking for interpretation
		interpretation, err := services.InterpretTarotCards(apiKey, cardNames, requestID, userStory, userName)
		if err != nil {
			// CHANGE 3: Store error message in LocalStorage instead of sending error immediately
			// This allows the frontend to handle the error gracefully
			LocalStorage[requestID.String()] = fmt.Sprintf("Error generating interpretation: %v", err)
			return
		}

		LocalStorage[requestID.String()] = interpretation

		// CHANGE 4: Removed redundant GetInterpretation call
		// The interpretation will be fetched by the frontend using the separate endpoint
		fmt.Println(interpretation)
	}()
}

func GetInterpretation(ctx *gin.Context) {
	requestID := ctx.Param("uuid")
	interpretation, ok := LocalStorage[requestID]
	if !ok {
		// CHANGE 5: More detailed error response
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "No interpretation found for this UUID",
			"details": "The interpretation might still be generating or there was an error in the process",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"interpretation": interpretation})
}

func ReverseRandomiser() bool {
	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomBool := randomiser.Intn(2)
	return randomBool == 0
}
