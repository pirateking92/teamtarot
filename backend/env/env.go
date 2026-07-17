// Package env loads runtime configuration from environment variables
// (optionally seeded from a local .env file).
package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all runtime configuration for the Cassandra backend.
type Config struct {
	DatabaseURL       string
	AnthropicKey      string
	AnthropicAuthToken string
	Port              string
	GinMode           string
	FrontendOrigin    string
	Testing           bool
}

// HasAnthropicCredential reports whether any Anthropic credential is configured
// (an API key or a bearer auth token). The SDK resolves which header to use from
// the environment; this just decides real-vs-stub interpretation.
func (c Config) HasAnthropicCredential() bool {
	return c.AnthropicKey != "" || c.AnthropicAuthToken != ""
}

// Load reads configuration from the environment. A local .env file, if present,
// is loaded first but is entirely optional — unlike the old behaviour, a missing
// .env is not fatal.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("env: no .env file loaded (this is fine outside local dev)")
	}

	return Config{
		DatabaseURL:        getenv("DATABASE_URL", "postgres://cassandra:cassandra@localhost:5432/cassandra?sslmode=disable"),
		AnthropicKey:       os.Getenv("ANTHROPIC_API_KEY"),
		AnthropicAuthToken: os.Getenv("ANTHROPIC_AUTH_TOKEN"),
		Port:               getenv("PORT", "8082"),
		GinMode:            os.Getenv("GIN_MODE"),
		FrontendOrigin:     getenv("FRONTEND_ORIGIN", "http://localhost:5173"),
		Testing:            os.Getenv("TESTING") == "True",
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
