package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role").(string)
		for _, role := range allowedRoles {
			if role == userRole {
				return c.Next()
			}
		}
		return helpers.ErrorResponse(c, fiber.StatusForbidden, "You do not have permission to access this resource")
	}
}
