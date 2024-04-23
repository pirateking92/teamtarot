package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"main.go/models"
)

// FetchTarotCards makes a GET request to the API to fetch tarot cards
// and returns a slice of Card structs representing the cards.
func FetchTarotCards() ([]models.Card, error) {

	apiUrl := "https://tarotapi.dev/api/v1/cards"

	// Send GET request to the API
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Decode JSON response into a slice of Card structs
	var cardsResponse struct {
		Cards []models.Card `json:"cards"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&cardsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return cardsResponse.Cards, nil
}


//Below is the code for sending a POST request to the OpenAI API to fetch interpretations of our previously fetched tarot cards.
//The function returns the AI-generated interpretation as a string.

const openAIKey = "OPENAI_API" //see .env for details.

func authenticateRequest(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+openAIKey)
}

const openAIEndpoint = "https://api.openai.com/v1/engines/davinci/completions" // davinci has been deprecated and the more recent APIs are not open-source/free. We'll need an alternative.

type Request struct {
	Prompt string `json:"prompt"`
}

type Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func GetAIResponse(Prompt string) (string, error) {
	data := Request{
		Prompt: Prompt,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	authenticateRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if len(result.Choices) == 0 {
		return "No completions found from the API", nil
	}

	return result.Choices[0].Text, nil
}
