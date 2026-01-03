package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/database"
	"github.com/savanyv/zenith-pay/internal/delivery/handlers"
	"github.com/savanyv/zenith-pay/internal/middlewares"
	"github.com/savanyv/zenith-pay/internal/model"
	"github.com/savanyv/zenith-pay/internal/repository"
	"github.com/savanyv/zenith-pay/internal/usecase"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func productRegisterRoutes(app fiber.Router, jwtService helpers.JWTService) {
	repo := repository.NewProductRepository(database.DB)
	categoryRepo := repository.NewCategoryRepository(database.DB)
	uc := usecase.NewProductUsecase(repo, categoryRepo)
	handler := handlers.NewProductHandler(uc)

	productRoutes := app.Group("/products", middlewares.JWTMiddleware(jwtService))
	productRoutes.Get("/", handler.ListProduct)
	productRoutes.Get("/:id", handler.GetProductByID)

	admin := productRoutes.Group("/admin", middlewares.RoleMiddleware(model.AdminRole))
	admin.Post("/", handler.CreateProduct)
	admin.Put("/:id", handler.UpdateProduct)
	admin.Delete("/:id", handler.DeleteProduct)
}
