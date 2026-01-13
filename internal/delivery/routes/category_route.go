package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/database"
	"github.com/savanyv/zenith-pay/internal/delivery/handlers"
	"github.com/savanyv/zenith-pay/internal/middlewares"
	"github.com/savanyv/zenith-pay/internal/model"
	"github.com/savanyv/zenith-pay/internal/repository"
	"github.com/savanyv/zenith-pay/internal/usecase"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func categoryRegisterRoutes(app fiber.Router, jwt helpers.JWTService) {
	repository := repository.NewCategoryRepository(database.DB)
	usecase := usecase.NewCategoryUsecase(repository)
	handler := handlers.NewCategoryHandler(usecase)

	categoryRoutes := app.Group("/categories", middlewares.JWTMiddleware(jwt), middlewares.RateLimiter(100, 1*time.Minute))
	categoryRoutes.Get("/", handler.ListCategories)
	categoryRoutes.Get("/:id", handler.GetCategoryByID)

	admin := categoryRoutes.Group("/admin", middlewares.RoleMiddleware(model.AdminRole), middlewares.RateLimiter(50, 1*time.Minute))
	admin.Post("/", handler.CreateCategory)
	admin.Put("/:id", handler.UpdateCategory)
	admin.Delete("/:id", handler.DeleteCategory)
}
