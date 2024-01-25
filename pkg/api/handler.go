package api

import (
	"web-hotlist-middleware/pkg/statistics"
)

// Handler struct holds a reference to the statistics service
type Handler struct {
	StatsService *statistics.StatisticsService
}

// NewHandler creates a new Handler with the given statistics service
func NewHandler(s *statistics.StatisticsService) *Handler {
	return &Handler{StatsService: s}
}

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
