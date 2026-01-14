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

func shiftRegisterRoute(app fiber.Router, jwtService helpers.JWTService) {
	repo := repository.NewShiftRepository(database.DB)
	uc := usecase.NewShiftUsecase(repo)
	handler := handlers.NewShiftHandler(uc)

	shiftRoutes := app.Group("/shifts", middlewares.JWTMiddleware(jwtService), middlewares.RoleMiddleware(model.CashierRole), middlewares.RateLimiter(10, 1*time.Minute))

	shiftRoutes.Post("/open", handler.OpenShift)
	shiftRoutes.Post("/close", handler.CloseShift)
	shiftRoutes.Get("/active", handler.GetActiveShift)
}
