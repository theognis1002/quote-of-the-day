package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuoteResponse(t *testing.T) {
	// Test JSON unmarshaling
	jsonData := `{
		"contents": {
			"quotes": [
				{
					"quote": "Test quote"
				}
			]
		}
	}`

	var response QuoteResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NoError(t, err)
	assert.Equal(t, "Test quote", response.Contents.Quotes[0].Quote)
}
