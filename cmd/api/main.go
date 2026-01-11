package main

import (
	"fmt"
	"log"
	"net/http"
	"streamflix/internal/api"
	"streamflix/internal/config"
	"streamflix/internal/repository"
	"streamflix/internal/service"
	"streamflix/pkg/postgres"
)

func main() {
	cfg := config.Load()

	//Connect to database
	dbConfig := postgres.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	}

	db, err := postgres.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to PostgreSQL successfully!")

	err = postgres.RunMigrations(db, "migrations")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	// Initalizing Repo
	userRepo := repository.NewPostgresUserRepository(db)

	// Initializing Service
	userService := service.NewUserService(userRepo)

	router := api.SetupRouter(userService)

	addr := ":" + cfg.Port
	fmt.Printf("Server starting on %s\n", addr)
	http.ListenAndServe(addr, router)
}
