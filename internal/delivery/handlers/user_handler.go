package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/zenith-pay/internal/dto"
	"github.com/savanyv/zenith-pay/internal/usecase"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	validator *helpers.CustomValidator
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		validator: helpers.NewCustomValidtor(),
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dtos.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.validator.Validate(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.Register(&req)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", res)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dtos.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.validator.Validate(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.Login(&req)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Login successful", res)
}
