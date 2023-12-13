package handlers

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
  GetAll(ctx *fiber.Ctx) error
}

type handler struct {}

func NewUserHandler() UserHandler {
  return &handler{}
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
  return ctx.JSON("GetAll")
}
