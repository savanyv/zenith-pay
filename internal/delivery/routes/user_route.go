package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/database"
	"github.com/savanyv/zenith-pay/internal/delivery/handlers"
	"github.com/savanyv/zenith-pay/internal/repository"
	"github.com/savanyv/zenith-pay/internal/usecase"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func userRegisterRoutes(app fiber.Router) {
	jwt := helpers.NewJWTService()
	repo := repository.NewUserRepository(database.DB)
	usecase := usecase.NewUserUsecase(repo, jwt)
	handler := handlers.NewUserHandler(usecase)

	app.Post("/zenith-pay/auth/register", handler.Register)
	app.Post("/zenith-pay/auth/login", handler.Login)
}
