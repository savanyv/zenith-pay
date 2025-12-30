package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/database"
	"github.com/savanyv/zenith-pay/internal/delivery/handlers"
	"github.com/savanyv/zenith-pay/internal/middlewares"
	"github.com/savanyv/zenith-pay/internal/repository"
	"github.com/savanyv/zenith-pay/internal/usecase"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func categoryRegisterRoutes(app fiber.Router, jwt helpers.JWTService) {
	repository := repository.NewCategoryRepository(database.DB)
	usecase := usecase.NewCategoryUsecase(repository)
	handler := handlers.NewCategoryHandler(usecase)

	categoryRoutes := app.Group("/zenith-pay/categories", middlewares.JWTMiddleware(jwt))

	categoryRoutes.Post("/", handler.CreateCategory)
	categoryRoutes.Get("/", handler.ListCategories)
	categoryRoutes.Get("/:id", handler.GetCategoryByID)
	categoryRoutes.Put("/:id", handler.UpdateCategory)
	categoryRoutes.Delete("/:id", middlewares.RoleMiddleware("admin"), handler.DeleteCategory)
}
