package test

import (
	"encoding/json"
	"testing"

	"quote-of-the-day/src/model"

	"github.com/stretchr/testify/assert"
)

func TestQuoteResponse(t *testing.T) {
	jsonData := `{
		"contents": {
			"quotes": [
				{
					"quote": "Test quote"
				}
			]
		}
	}`

	var response model.QuoteResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NoError(t, err)
	assert.Equal(t, "Test quote", response.Contents.Quotes[0].Quote)
}
