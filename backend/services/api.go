package services

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strings"

	"github.com/google/uuid"
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

// InterpretTarotCards interprets tarot cards using the OpenAI API
func InterpretTarotCards(apiKey string, cards []string, RequestID uuid.UUID) (string, error) {
	client := &http.Client{}

	userStory := "I've just started a new job and I don't know if it was the right decision."
	prompt := fmt.Sprintf("I'm doing a tarot card reading. They drew %s, %s, and %s. Please interpret these cards in relation to their story: '%s'. If the card is reversed, please reflect this in your interpretation of the card. Please format your response in the style of a tarot card reader, and keep your response lower than 200 words", cards[0:2], cards[2:4], cards[4:6], userStory)
	payload := fmt.Sprintf(`{"model": "gpt-3.5-turbo-instruct", "prompt": "%s", "max_tokens": 400}`, prompt)
	fmt.Println(prompt)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var responseBody strings.Builder
	if _, err := io.Copy(&responseBody, resp.Body); err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	type Response struct {
		Choices []struct {
			Message struct {
				Text string `json:"text"`
			} `json:"message"`
		} `json:"choices"`
	}

	var response Response
	if err := json.Unmarshal([]byte(responseBody.String()), &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)

	}
	result := response.Choices[0].Message.Text
	fmt.Println(result)
	fmt.Println(resp.StatusCode)

	return result, nil

}
