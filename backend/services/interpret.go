package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// DrawnCard is a single card in a reading, with its resolved meaning, as fed to
// the interpretation prompt.
type DrawnCard struct {
	Position string
	Name     string
	Reversed bool
	Meaning  string
}

// InterpretInput is everything the interpreter needs to write a reading.
type InterpretInput struct {
	QuerentName string
	Story       string
	SpreadName  string
	Cards       []DrawnCard
}

const cassandraSystemPrompt = `You are Cassandra, a warm, perceptive tarot reader with a slightly theatrical, old-world voice. You give tarot readings that feel personal and grounded, never generic.

Guidelines:
- Speak directly to the seeker, by name if you know it.
- Weave the cards together into a single coherent narrative that follows the spread's positions in order — don't just list each card's textbook meaning.
- Tie the reading to the seeker's own story where they've shared one.
- Honour each card's orientation (upright vs reversed) in what it signifies.
- Be encouraging and honest; acknowledge difficulty without doom.
- Keep it to a few flowing paragraphs. Do not use markdown headers or bullet lists.
- End with a gentle closing thought. Do NOT ask the seeker a follow-up question.`

// GenerateInterpretation calls Claude (claude-opus-4-8) to produce a tarot
// reading from the drawn cards and the seeker's story. Credentials and base URL
// are resolved by the SDK from the environment (ANTHROPIC_API_KEY →
// x-api-key, ANTHROPIC_AUTH_TOKEN → bearer, ANTHROPIC_BASE_URL → endpoint),
// which is loaded from backend/.env at startup.
func GenerateInterpretation(ctx context.Context, in InterpretInput) (string, error) {
	var opts []option.RequestOption
	// OAuth-style bearer tokens (ANTHROPIC_AUTH_TOKEN, e.g. a gateway/base-URL
	// setup) require the oauth beta header on /v1/messages; plain API keys don't.
	if os.Getenv("ANTHROPIC_AUTH_TOKEN") != "" && os.Getenv("ANTHROPIC_API_KEY") == "" {
		opts = append(opts, option.WithHeader("anthropic-beta", "oauth-2025-04-20"))
	}
	client := anthropic.NewClient(opts...)

	name := strings.TrimSpace(in.QuerentName)
	if name == "" {
		name = "the seeker"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Give a %q reading for %s.\n\n", in.SpreadName, name)
	if story := strings.TrimSpace(in.Story); story != "" {
		fmt.Fprintf(&b, "Their story, in their words: %s\n\n", story)
	}
	b.WriteString("The cards they drew, in order:\n")
	for _, c := range in.Cards {
		orient := "upright"
		if c.Reversed {
			orient = "reversed"
		}
		fmt.Fprintf(&b, "- %s — %s (%s): %s\n", c.Position, c.Name, orient, c.Meaning)
	}
	b.WriteString("\nGive them their reading now.")

	msg, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeOpus4_8,
		MaxTokens: 2048,
		System:    []anthropic.TextBlockParam{{Text: cassandraSystemPrompt}},
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(b.String())),
		},
	})
	if err != nil {
		return "", fmt.Errorf("anthropic request failed: %w", err)
	}

	var out strings.Builder
	for _, block := range msg.Content {
		if block.Type == "text" {
			out.WriteString(block.Text)
		}
	}
	text := strings.TrimSpace(out.String())
	if text == "" {
		return "", fmt.Errorf("empty interpretation returned")
	}
	return text, nil
}
