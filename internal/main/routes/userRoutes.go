package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/presentation/handlers"
)

func handleUserRoutes(router fiber.Router) {
  userRouter := router.Group("/users")

  userHandler := handlers.NewUserHandler()

  userRouter.Get("/", userHandler.GetAll)
}
