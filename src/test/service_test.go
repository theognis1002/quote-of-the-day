package test

import (
	"testing"

	"quote-of-the-day/src/internal"

	"github.com/stretchr/testify/assert"
)

type MockQuoteService struct {
	quote string
	err   error
}

func NewMockQuoteService(quote string, err error) *MockQuoteService {
	return &MockQuoteService{quote: quote, err: err}
}

func (s *MockQuoteService) GetQuote() (string, error) {
	return s.quote, s.err
}
func TestOpenAIQuoteService(t *testing.T) {
	t.Run("new service with api key", func(t *testing.T) {
		service := internal.NewOpenAIQuoteService("test-key")
		assert.NotNil(t, service)
	})

	t.Run("get quote without api key", func(t *testing.T) {
		service := internal.NewOpenAIQuoteService("")
		quote, err := service.GetQuote()
		assert.Error(t, err)
		assert.Empty(t, quote)
		assert.Contains(t, err.Error(), "OPENAI_API_KEY environment variable is not set")
	})
}

func TestMockQuoteService(t *testing.T) {
	t.Run("successful quote", func(t *testing.T) {
		expectedQuote := "Test quote"
		service := NewMockQuoteService(expectedQuote, nil)
		quote, err := service.GetQuote()
		assert.NoError(t, err)
		assert.Equal(t, expectedQuote, quote)
	})

	t.Run("error case", func(t *testing.T) {
		expectedErr := assert.AnError
		service := NewMockQuoteService("", expectedErr)
		quote, err := service.GetQuote()
		assert.Error(t, err)
		assert.Empty(t, quote)
		assert.Equal(t, expectedErr, err)
	})
}
