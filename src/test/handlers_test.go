package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"quote-of-the-day/src/internal"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/", internal.IndexHandler)
	r.GET("/quote-of-the-day", internal.QuoteOfTheDayHandler)
	r.POST("/clear-cache", internal.ClearCacheHandler)
	return r
}

func TestIndexHandler(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["message"], "Welcome!")
}

func TestQuoteOfTheDayHandler(t *testing.T) {
	testCases := []struct {
		name           string
		mockQuote      string
		mockErr        error
		expectedStatus int
		checkResponse  func(*testing.T, map[string]string)
	}{
		{
			name:           "successful quote",
			mockQuote:      "Test quote - Author",
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, response map[string]string) {
				assert.Equal(t, "Test quote - Author", response["message"])
			},
		},
		{
			name:           "service error",
			mockQuote:      "",
			mockErr:        assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			checkResponse: func(t *testing.T, response map[string]string) {
				assert.Contains(t, response["error"], "Sorry!")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up mock service before test
			mockService := NewMockQuoteService(tc.mockQuote, tc.mockErr)
			internal.SetTestCache(&internal.QuoteCache{
				Service: mockService,
			})

			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/quote-of-the-day", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			tc.checkResponse(t, response)
		})
	}
}

func TestClearCacheHandler(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/clear-cache", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Cache cleared successfully", response["message"])
}
