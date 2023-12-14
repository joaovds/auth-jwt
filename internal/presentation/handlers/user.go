package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/domain"
)

type UserHandler interface {
	GetAll(ctx *fiber.Ctx) error
}

type handler struct {
	userService domain.UserUseCases
}

func NewUserHandler(userService domain.UserUseCases) UserHandler {
	return &handler{
		userService,
	}
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}
