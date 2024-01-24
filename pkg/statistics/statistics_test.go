package statistics

import (
	"testing"
)

func TestGetStatistics(t *testing.T) {
	// Initialize a new Statistics object
	s := NewStatistics()

	// Call the GetStatistics method
	stats, err := s.GetStatistics("POST", "daily")

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
	s := NewStatistics()

	// Call the UpdateStatistics method
	err := s.UpdateStatistics("POST", "identifier", 1)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}