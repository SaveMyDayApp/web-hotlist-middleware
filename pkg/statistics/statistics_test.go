package statistics

import (
	"log"
	"testing"
	"web-hotlist-middleware/pkg/storage"
)

func TestGetStatistics(t *testing.T) {
	// Initialize a new Statistics object
	redisClient, err := storage.NewRedisClient("45.76.165.171:6379", "CFredis123", 1)
	if err != nil {
		log.Fatalf("Failed to initialize Redis client: %v", err)
	}
	s := NewStatisticsService(redisClient.Client)

	// Call the GetStatistics method
	stats, err := s.GetStatistics("POST", "day", 10)
	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the result
	if len(stats) == 0 {
		t.Errorf("Expected non-empty statistics")
	}
}

func TestUpdateStatistics(t *testing.T) {
	// Initialize a new Statistics object
	redisClient, err := storage.NewRedisClient("45.76.165.171:6379", "CFredis123", 1)
	if err != nil {
		log.Fatalf("Failed to initialize Redis client: %v", err)
	}
	s := NewStatisticsService(redisClient.Client)

	// Call the UpdateStatistics method
	err = s.SubmitStatistic("POST", "identifier", 1)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
