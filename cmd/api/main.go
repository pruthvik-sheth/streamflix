package main

import (
	"fmt"
	"net/http"
	"streamflix/internal/api"
	"streamflix/internal/config"
	"streamflix/internal/repository"
	"streamflix/internal/service"
)

func main() {
	cfg := config.Load()

	// Initalizing Repo
	userRepo := repository.NewMemoryUserRepository()

	// Initializing Service
	userService := service.NewUserService(userRepo)

	router := api.SetupRouter(userService)

	addr := ":" + cfg.Port
	fmt.Printf("Server starting on %s\n", addr)
	http.ListenAndServe(addr, router)
}
