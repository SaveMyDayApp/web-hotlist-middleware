package main

import (
	"log"
	"net/http"

	"web-hotlist-middleware/pkg/api"
	"web-hotlist-middleware/pkg/statistics"
	"web-hotlist-middleware/pkg/storage"
)

func main() {
	// Initialize Redis connection
	redisClient, err := storage.NewRedisClient("45.76.165.171:6379", "CFredis123", 1)
	if err != nil {
		log.Fatalf("Failed to initialize Redis client: %v", err)
	}
	//Initialize statistics service
	statisticsService := statistics.NewStatisticsService(redisClient.Client)
	// Initialize API handlers
	handler := api.NewHandler(statisticsService)

	// Set up routes
	router := api.NewRouter(handler)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
