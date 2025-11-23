package api

import (
	"net/http"
	"streamflix/internal/api/handlers"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/ping", handlers.PingHandler)

	return mux
}
