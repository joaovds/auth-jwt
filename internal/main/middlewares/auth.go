package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/infra/cryptography"
)

func Auth(ctx *fiber.Ctx) error {
  token := ctx.Cookies("Authentication")

  if token == "" {
    return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "message": "Unauthorized",
    })
  }
 
  crypt := cryptography.NewCryptography()

  err := crypt.Decrypt(token)
  if err != nil {
    return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "message": "Unauthorized",
    })
  }

  return ctx.Next()
}
