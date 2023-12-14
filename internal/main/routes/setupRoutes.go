package routes


import (
  "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  _ = app.Group("/api/v1")
}
