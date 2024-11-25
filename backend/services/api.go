package services

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strings"

	"github.com/google/uuid"
	"main.go/errors"
	"main.go/models"
)

// FetchTarotCards makes a GET request to the API to fetch tarot cards
// and returns a slice of Card structs representing the cards.
func FetchTarotCards() ([]models.Card, error) {

	apiUrl := "http://tarotapi.dev/api/v1/cards"

	// Send GET request to the API
	resp, err := http.Get(apiUrl)
	if err != nil {
		errors.SendInternalError(nil, fmt.Errorf("failed to make GET request: %v", err))
	}
	defer resp.Body.Close()

	// Decode JSON response into a slice of Card structs
	var cardsResponse struct {
		Cards []models.Card `json:"cards"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&cardsResponse); err != nil {
		errors.SendInternalError(nil, fmt.Errorf("failed to decode JSON response: %v", err))
	}

	return cardsResponse.Cards, nil
}

// InterpretTarotCards interprets tarot cards using the Anthropic API
func InterpretTarotCards(apiKey string, cards []string, requestID uuid.UUID, userStory string, userName string) (string, error) {
	client := &http.Client{}

	// Create the prompt with the card names and format
	prompt := fmt.Sprintf("You're doing a tarot card reading for %s as a tarot card reader called Cassandra. They drew %s (for their past), %s (for their present), and %s (for their future). This is their story: %s. Give them a reading based on the cards they have drawn, and the story they have given you. Don't ask a follow up question.", userName, cards[0:2], cards[2:4], cards[4:6], userStory)

	requestBody := map[string]interface{}{
		"model":      "claude-3-5-sonnet-20241022",
		"max_tokens": 2000,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", strings.NewReader(string(jsonBody)))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body) // Read error response body for logging
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var responseBody strings.Builder
	if _, err := io.Copy(&responseBody, resp.Body); err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	type Response struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}

	var response Response
	if err := json.Unmarshal([]byte(responseBody.String()), &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if len(response.Content) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	// Clean text if necessary
	cleanedText := strings.ReplaceAll(response.Content[0].Text, "[", "")
	cleanedText = strings.ReplaceAll(cleanedText, "]", "")

	return cleanedText, nil
}
