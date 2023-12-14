package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler interface {
	GetAll(ctx *fiber.Ctx) error
  GetByID(ctx *fiber.Ctx) error
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

func (h *handler) GetByID(ctx *fiber.Ctx) error {
  id := ctx.Params("id")
  if id == "" {
    return ctx.Status(fiber.StatusBadRequest).JSON("id is required")
  }

  _, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return ctx.Status(fiber.StatusBadRequest).JSON("id is invalid")
  }

  user, err := h.userService.GetByID(id)
  if err != nil {
    return ctx.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return ctx.Status(fiber.StatusOK).JSON(user)
}
