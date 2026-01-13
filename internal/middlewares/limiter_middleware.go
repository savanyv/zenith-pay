package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(max int, duration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max: max,
		Expiration: duration,
		KeyGenerator: func(c *fiber.Ctx) string {
			if userID := c.Locals("userID"); userID != nil {
				return userID.(string)
			}
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"code": 429,
				"status": "Too Many Request",
				"message": "Request limit exceed. Please try again later.",
			})
		},
	})
}
