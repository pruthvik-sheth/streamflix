package api

import (
	"net/http"
	"streamflix/internal/api/handlers"
	"streamflix/internal/service"
)

func SetupRouter(userService *service.UserService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/ping", handlers.PingHandler)

	userHandler := handlers.NewUserHandler(userService)
	mux.HandleFunc("/api/register", userHandler.Register)
	mux.HandleFunc("/api/login", userHandler.Login)

	return mux
}
