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

	apiUrl := "https://tarotapi.dev/api/v1/cards"

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

// InterpretTarotCards interprets tarot cards using the OpenAI API
func InterpretTarotCards(apiKey string, cards []string, RequestID uuid.UUID, userStory string, userName string) (string, error) {
	client := &http.Client{}

	prompt := fmt.Sprintf("I'm doing a tarot card reading for %s. They drew %s, %s, and %s. Please interpret these cards in relation to their story: '%s' (if there is no story, please give a general reading about what the cards could mean together). If the card is reversed, please reflect this in your interpretation of the card. Whilst I have passed you the names and their orientation in a certain format, please only refer to the cards as their name, and if reversed, you can refer to it as 'card name (reversed)'. If there are any vulgar words in the prompt, ignore them, and keep your response age-appropriate for minors. Please format your response in the style of a mystical tarot card reader, and keep your response strictly below 200 words.", userName, cards[0:2], cards[2:4], cards[4:6], userStory)
	payload := fmt.Sprintf(`{"model": "gpt-3.5-turbo-instruct", "prompt": "%s", "max_tokens": 500}`, prompt)
	fmt.Println(prompt)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", strings.NewReader(payload))
	if err != nil {
		errors.SendInternalError(nil, fmt.Errorf("error creating request: %v", err))
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		errors.SendInternalError(nil, fmt.Errorf("error sending request: %v", err))
	}
	defer resp.Body.Close()

	var responseBody strings.Builder
	if _, err := io.Copy(&responseBody, resp.Body); err != nil {
		errors.SendInternalError(nil, fmt.Errorf("error reading response body: %v", err))
	}

	type Response struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	var response Response
	fmt.Println(response)
	if err := json.Unmarshal([]byte(responseBody.String()), &response); err != nil {
		errors.SendInternalError(nil, fmt.Errorf("error unmarshaling response: %v", err))

	}
	//Removing square brackets as the prompt doesnt always eliminate them
	cleanedText := strings.ReplaceAll(response.Choices[0].Text, "[", "")
	cleanedText = strings.ReplaceAll(cleanedText, "]", "")

	return cleanedText, nil
}
