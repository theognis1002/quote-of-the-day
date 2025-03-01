package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome! Please hit the `/quote-of-the-day` API to get the quote of the day.",
	})
}

func QuoteOfTheDayHandler(c *gin.Context) {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	cache.Mutex.RLock()
	if cache.Date == date && cache.Quote != "" {
		log.Println("Cache Hit for date ", date)
		cache.Mutex.RUnlock()
		c.JSON(http.StatusOK, gin.H{"message": cache.Quote})
		return
	}
	cache.Mutex.RUnlock()

	log.Println("Cache miss for date ", date)
	quote, err := cache.Service.GetQuote()
	if err != nil {
		log.Println("Error getting quote from service: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sorry! We could not get the Quote of the Day. Please try again.",
		})
		return
	}

	cache.Mutex.Lock()
	cache.Quote = quote
	cache.Date = date
	cache.Mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": quote})
}

func ClearCacheHandler(c *gin.Context) {
	cache.Mutex.Lock()
	cache.Quote = ""
	cache.Date = ""
	cache.Mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": "Cache cleared successfully",
	})
}
