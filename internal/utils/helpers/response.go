package helpers

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Code int `json:"code,omitempty"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Code: statusCode,
		Status: "Success",
		Message: message,
		Data: data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Code: statusCode,
		Status: "Error",
		Message: message,
	})
}
