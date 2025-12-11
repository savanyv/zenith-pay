package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/config"
	"github.com/savanyv/zenith-pay/internal/database"
)

// Server represents the application server
type Server struct {
	app *fiber.App
	config *config.Config
}

// NewServer creates a new Server instance with the given configuration
func NewServer(config *config.Config) *Server {
	return &Server{
		app: fiber.New(),
		config: config,
	}
}

// Start initializes the database, sets up routes, and starts the server
func (s *Server) Start() error {
	// Database
	_, err := database.InitDatabase(s.config)
	if err != nil {
		return err
	}

	// Routes

	// Server Listen
	if err := s.app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
		return err
	}

	log.Println("Server is running on port 3000")
	return nil
}
