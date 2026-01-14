package main

import (
	"log"
	"os"

	"github.com/savanyv/zenith-pay/config"
	"github.com/savanyv/zenith-pay/internal/app"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	log.Printf("🚀 Starting %s (%s environment)",
		cfg.AppName,
		cfg.AppEnv,
	)

	server := app.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Printf("Server stopped with error: %v", err)
		os.Exit(1)
	}
}
