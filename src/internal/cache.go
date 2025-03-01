package internal

import (
	"os"
	"sync"
)

type QuoteCache struct {
	Mutex   sync.RWMutex
	Quote   string
	Date    string
	Service QuoteService
}

var cache = &QuoteCache{
	Service: NewOpenAIQuoteService(getEnv("OPENAI_API_KEY", "")),
}

func (c *QuoteCache) GetQuote() (string, string) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.Quote, c.Date
}

func (c *QuoteCache) SetQuote(quote, date string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Quote = quote
	c.Date = date
}

func (c *QuoteCache) Clear() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Quote = ""
	c.Date = ""
}

func SetTestCache(c *QuoteCache) {
	cache = c
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
