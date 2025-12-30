package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/database"
	"github.com/savanyv/zenith-pay/internal/delivery/handlers"
	"github.com/savanyv/zenith-pay/internal/repository"
	"github.com/savanyv/zenith-pay/internal/usecase"
)

func userRegisterRoutes(app fiber.Router) {
	repo := repository.NewUserRepository(database.DB)
	usecase := usecase.NewUserUsecase(repo)
	handler := handlers.NewUserHandler(usecase)

	app.Post("/zenith-pay/auth/register", handler.Register)
	app.Post("/zenith-pay/auth/login", handler.Login)
}
