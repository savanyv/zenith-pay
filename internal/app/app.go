package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/config"
	"github.com/savanyv/zenith-pay/internal/database"
	"github.com/savanyv/zenith-pay/internal/database/seed"
	"github.com/savanyv/zenith-pay/internal/delivery/routes"
	"github.com/savanyv/zenith-pay/internal/middlewares"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

// Server represents the application server
type Server struct {
	app *fiber.App
	config *config.Config
}

// NewServer creates a new Server instance with the given configuration
func NewServer(config *config.Config) *Server {
	app := fiber.New(fiber.Config{
		AppName: config.AppName,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
	})

	return &Server{
		app: app,
		config: config,
	}
}

// Start initializes the database, sets up routes, and starts the server
func (s *Server) Start() error {
	// database
	if _, err := database.InitDatabase(s.config); err != nil {
		return fmt.Errorf("init database: %w", err)
	}

	// Seed
	if s.config.AppEnv == "development" {
		bcHelper := helpers.NewBcryptHelper()
		seed.SeedAdmin(database.DB, bcHelper)
	}

	// Middlewares And Routes
	s.app.Use(middlewares.CORSMiddleware())
	s.app.Use(middlewares.MethodValidationMiddleware())
	routes.RegisterRoutes(s.app)

	// Start server
	addr := fmt.Sprintf(":%s", s.config.AppPort)
	go func() {
		log.Printf("🚀 Server running on %s", addr)
		if err := s.app.Listen(addr); err != nil {
			log.Printf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("⏳ Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	log.Println("✅ Server shut down successfully")
	return nil
}
