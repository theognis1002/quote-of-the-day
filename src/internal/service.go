package internal

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type QuoteService interface {
	GetQuote() (string, error)
}

type OpenAIQuoteService struct {
	apiKey string
}

func NewOpenAIQuoteService(apiKey string) *OpenAIQuoteService {
	return &OpenAIQuoteService{apiKey: apiKey}
}

func (s *OpenAIQuoteService) GetQuote() (string, error) {
	return getQuoteFromLLM()
}

func getQuoteFromLLM() (string, error) {
	apiKey := getEnv("OPENAI_API_KEY", "")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "You are a quote generator. Provide a single meaningful quote about life, " +
						"success, or wisdom. Include the author. Keep it concise.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate a quote of the day.",
				},
			},
			MaxTokens: 100,
		},
	)

	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
