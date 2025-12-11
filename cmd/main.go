package main

import (
	"github.com/savanyv/zenith-pay/config"
	"github.com/savanyv/zenith-pay/internal/app"
)

func main() {
	// Load configuration
	config := config.LoadConfig()

	// Initialize and start the server
	server := app.NewServer(config)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
