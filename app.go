package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome! Please hit the `/quote-of-the-day` API to get the quote of the day.",
	})
}

func quoteOfTheDayHandler(c *gin.Context) {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	cache.mutex.RLock()
	if cache.date == date && cache.quote != "" {
		log.Println("Cache Hit for date ", date)
		cache.mutex.RUnlock()
		c.JSON(http.StatusOK, gin.H{"message": cache.quote})
		return
	}
	cache.mutex.RUnlock()

	// Cache miss - get new quote
	log.Println("Cache miss for date ", date)
	quote, err := getQuoteFromLLM()
	if err != nil {
		log.Println("Error getting quote from LLM: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sorry! We could not get the Quote of the Day. Please try again.",
		})
		return
	}

	// Update cache
	cache.mutex.Lock()
	cache.quote = quote
	cache.date = date
	cache.mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": quote})
}

func clearCacheHandler(c *gin.Context) {
	cache.mutex.Lock()
	cache.quote = ""
	cache.date = ""
	cache.mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": "Cache cleared successfully",
	})
}

func main() {
	// Load .env file at the start of main
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Create Gin router
	r := gin.Default() // This includes logging and recovery middleware

	// Routes
	r.GET("/", indexHandler)
	r.GET("/quote-of-the-day", quoteOfTheDayHandler)
	r.POST("/clear-cache", clearCacheHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	// Start Server
	go func() {
		log.Println("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
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
