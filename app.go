package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

// Add cache structure
type QuoteCache struct {
	mutex sync.RWMutex
	quote string
	date  string
}

var cache = &QuoteCache{}

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message: "Welcome! Please hit the `/quote-of-the-day` API to get the quote of the day.",
	})
}

func quoteOfTheDayHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02")

		cache.mutex.RLock()
		if cache.date == date && cache.quote != "" {
			log.Println("Cache Hit for date ", date)
			json.NewEncoder(w).Encode(Response{Message: cache.quote})
			cache.mutex.RUnlock()
			return
		}
		cache.mutex.RUnlock()

		// Cache miss - get new quote
		log.Println("Cache miss for date ", date)
		quote, err := getQuoteFromLLM()
		if err != nil {
			log.Println("Error getting quote from LLM: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "Sorry! We could not get the Quote of the Day. Please try again.",
			})
			return
		}

		// Update cache
		cache.mutex.Lock()
		cache.quote = quote
		cache.date = date
		cache.mutex.Unlock()

		json.NewEncoder(w).Encode(Response{Message: quote})
	}
}

func main() {
	// Load .env file at the start of main
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/quote-of-the-day", quoteOfTheDayHandler())
	r.HandleFunc("/clear-cache", clearCacheHandler()).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
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

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func clearCacheHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cache.mutex.Lock()
		cache.quote = ""
		cache.date = ""
		cache.mutex.Unlock()

		json.NewEncoder(w).Encode(Response{
			Message: "Cache cleared successfully",
		})
	}
}
