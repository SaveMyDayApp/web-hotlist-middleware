package api

import (
	"net/http"
	"encoding/json"
	"web-hotlist-middleware/pkg/statistics"
)

// Handler struct holds a reference to the statistics service
type Handler struct {
	StatsService *statistics.Service
}

// NewHandler creates a new Handler with the given statistics service
func NewHandler(s *statistics.Service) *Handler {
	return &Handler{StatsService: s}
}

// GetStatistics handles the GET /api/statistics endpoint
func (h *Handler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract parameters from request and call the appropriate method on the statistics service
	// Write the response as JSON
}

// SubmitStatistics handles the POST /api/submit-statistics endpoint
func (h *Handler) SubmitStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract data from request body and call the appropriate method on the statistics service
	// Write the response as JSON
}