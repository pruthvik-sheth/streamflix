package main

import (
	"fmt"
	"net/http"
	"streamflix/internal/api"
	"streamflix/internal/config"
)

func main() {
	cfg := config.Load()
	router := api.SetupRouter()

	addr := ":" + cfg.Port
	fmt.Printf("Server starting on %s\n", addr)
	http.ListenAndServe(addr, router)
}
