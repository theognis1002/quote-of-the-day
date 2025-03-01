package main

import (
	"log"
	"quote-of-the-day/src/internal"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}
	internal.Run()
}
