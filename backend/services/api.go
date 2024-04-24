package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

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
func InterpretTarotCards(apiKey string, cards []string) (string, error) {
	// Set up the HTTP client
	client := &http.Client{}

	// Create a JSON payload for the request
	prompt := fmt.Sprintf("You drew %s, %s, and %s. Please interpret these cards.", cards[0], cards[1], cards[2])
	payload := fmt.Sprintf(`{"model": "gpt-3.5-turbo-instruct", "prompt": "%s", "max_tokens": 100}`, prompt)

	// Create a request object
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set the API key in the request headers
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and process the response
	var responseBody strings.Builder
	if _, err := io.Copy(&responseBody, resp.Body); err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	fmt.Println(responseBody.String())
	fmt.Println(resp.StatusCode)

	return responseBody.String(), nil
}

func SeeModels() {
	apiKey := "sk-proj-xxg71fBX7UwL5XI3JPeQT3BlbkFJs4TMMqsxh8odON7eCkXQ"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a GET request to the /v1/engines endpoint
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/engines", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set the API key in the request headers
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Print the response body (list of available engines/models)
	fmt.Println(string(body))
}
