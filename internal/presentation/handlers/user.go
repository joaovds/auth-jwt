package handlers

import (
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
  Create(ctx *fiber.Ctx) error
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *handler) Create(ctx *fiber.Ctx) error {
  user := new(domain.User)
  err := ctx.BodyParser(user)
  if err != nil {
    println(err.Error())
    return ctx.Status(fiber.StatusBadRequest).JSON("invalid body")
  }

  if user.Name == "" {
    return ctx.Status(fiber.StatusBadRequest).JSON("name is required")
  }

  if user.Email == "" {
    return ctx.Status(fiber.StatusBadRequest).JSON("email is required")
  }

  _, err = mail.ParseAddress(user.Email)
  if err != nil {
    return ctx.Status(fiber.StatusBadRequest).JSON("email is invalid")
  }

  if user.Password == "" {
    return ctx.Status(fiber.StatusBadRequest).JSON("password is required")
  }

  err = h.userService.Create(user)
  if err != nil {
    return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
  }
  return ctx.Status(fiber.StatusCreated).JSON(nil)
}
