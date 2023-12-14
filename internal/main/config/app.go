package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/main/routes"
)

func GetApp() *fiber.App {
	app := fiber.New()

	routes.SetupRoutes(app)

	return app
}
