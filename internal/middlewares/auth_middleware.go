package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

func AuthMiddleware(jwtService helpers.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "Authorization header is missing")
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid authorization header format")
		}

		token := bearerToken[1]
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		c.Locals("userID", claims.UserID)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}
