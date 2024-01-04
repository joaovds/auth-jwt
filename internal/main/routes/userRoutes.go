package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/auth-jwt/internal/application"
	"github.com/joaovds/auth-jwt/internal/main/middlewares"
	"github.com/joaovds/auth-jwt/internal/presentation/handlers"
)

func handleUserRoutes(router fiber.Router) {
	userRouter := router.Group("/users")

	userService := application.NewUserUseCases()
	userHandler := handlers.NewUserHandler(userService)

	userRouter.Get("/", middlewares.Auth, userHandler.GetAll)
	userRouter.Get("/:id", userHandler.GetByID)
  userRouter.Post("/", userHandler.Create)
  userRouter.Post("/login", userHandler.Login)
}
