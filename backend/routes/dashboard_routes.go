package routes

import (
	"yamanmnur/simple-dashboard/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitDashboardRoutes(app *fiber.App, controller *handlers.DashboardController) {
	authRoutes := app.Group("/api/v1/dashboard")
	authRoutes.Get("", controller.GetDashboard)
}
