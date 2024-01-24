package main

import (
	"log"
	"net/http"

	"github.com/yourusername/web-hotlist-middleware/pkg/api"
	"github.com/yourusername/web-hotlist-middleware/pkg/storage"
)

func main() {
	// Initialize Redis connection
	redisClient, err := storage.NewRedisClient("localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Initialize API handlers
	handler := api.NewHandler(redisClient)

	// Set up routes
	router := api.NewRouter(handler)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}