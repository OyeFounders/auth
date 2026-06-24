package main

import (
	"log"

	"github.com/OyeFounders/auth/internal/config"
	"github.com/OyeFounders/auth/internal/logger"
	"github.com/OyeFounders/auth/internal/server"
)

func main() {
	cfg := config.Load()
	appLogger := logger.New()
	router := server.NewRouter(appLogger)

	appLogger.Printf("starting auth service on port %s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
