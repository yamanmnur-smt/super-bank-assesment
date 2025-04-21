package routes

import (
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoutes(app *fiber.App, controller *handlers.AuthController) {
	authRoutes := app.Group("/api/v1/auth")
	authRoutes.Post("/login", controller.Login)
	authProfileRoutes := app.Group("/api/v1/auth/profile", middlewares.JwtMiddleware)
	authProfileRoutes.Get("", controller.Profile)
}
