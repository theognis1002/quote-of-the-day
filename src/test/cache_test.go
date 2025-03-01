package test

import (
	"testing"
	"time"

	"quote-of-the-day/src/internal"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := &internal.QuoteCache{}
	testQuote := "Test quote"
	testDate := time.Now().Format("2006-01-02")

	// Test initial state
	quote, date := cache.GetQuote()
	assert.Empty(t, quote)
	assert.Empty(t, date)

	// Test setting and getting quote
	cache.SetQuote(testQuote, testDate)
	quote, date = cache.GetQuote()
	assert.Equal(t, testQuote, quote)
	assert.Equal(t, testDate, date)

	// Test clearing cache
	cache.Clear()
	quote, date = cache.GetQuote()
	assert.Empty(t, quote)
	assert.Empty(t, date)
}
