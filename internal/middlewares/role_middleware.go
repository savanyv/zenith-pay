package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/model"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func RoleMiddleware(allowedRoles ...model.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleValue := c.Locals("role")
		if roleValue == nil {
			return helpers.ErrorResponse(c, fiber.StatusForbidden, "Access Denied")
		}

		userRoleStr, ok := roleValue.(string)
		if !ok {
			return helpers.ErrorResponse(c, fiber.StatusForbidden, "Access Denied")
		}

		userRole := model.Role(userRoleStr)
		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}

		return helpers.ErrorResponse(c, fiber.StatusForbidden, "Access Denied")
	}
}
