package api

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/statistics", handleGetStatistics)
	router.HandleFunc("/api/submit-statistics", handlePostStatistics)

	return router
}

func handleGetStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the handler for getting statistics
}

func handlePostStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the handler for posting statistics
}