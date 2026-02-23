package chains

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var Vacations []*Vacation

func GetVacationFromDb(id uuid.UUID) (Vacation, error) {
	idx := slices.IndexFunc(Vacations, func(v *Vacation) bool { return v.Id == id })
	if idx == -1 {
		return Vacation{}, errors.New("Id not found")
	}
	return *Vacations[idx], nil
}

// call the open ai api key
func GenerateVacationIdeaChange(id uuid.UUID, budget int, season string, hobbies []string) {
	log.Printf("Generating new vacation with ID : %s ", id)

	v := &Vacation{Id: id, Completed: false, Idea: ""}
	Vacations = append(Vacations, v)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Printf("OPENAI_API_KEY is not set")
		return
	}

	llmClient, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithModel("gpt-4o-mini"),
	)
	if err != nil {
		log.Printf("OpenAI client error: %v", err)
		return
	}

	prompt := fmt.Sprintf(
		"You are an AI travel agent. Create one practical vacation itinerary.\n"+
			"Season: %s\n"+
			"Hobbies: %s\n"+
			"Budget (USD): %d\n"+
			"Return a concise day-by-day plan plus estimated major costs.",
		season,
		strings.Join(hobbies, ", "),
		budget,
	)

	idea, err := llms.GenerateFromSinglePrompt(
		ctx,
		llmClient,
		prompt,
		llms.WithTemperature(0.4),
		openai.WithMaxCompletionTokens(350),
	)
	if err != nil {
		log.Printf("OpenAI generation error: %v", err)
		return
	}

	v.Idea = strings.TrimSpace(idea)
	v.Completed = true
	log.Printf("Generation for %s is done!", v.Id)
}
