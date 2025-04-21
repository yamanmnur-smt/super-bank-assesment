package routes

import (
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitCustomerRoute(app *fiber.App, controller *handlers.CustomerController) {
	customerRoutes := app.Group("/api/v1/customer", middlewares.JwtMiddleware)
	customerRoutes.Get("/list", controller.GetCustomersWithPagination)
	customerRoutes.Get("/detail/:id", controller.GetCustomerByID)
	customerRoutes.Post("/create", controller.Create)
	customerRoutes.Delete("/delete/:id", controller.GetCustomerByID)
	customerRoutes.Put("/update-photo", controller.UpdatePhotoCustomer)
	customerRoutes.Post("/upload-photo", controller.UploadPhotoCustomer)
	customerRoutes.Get("/get-photo", controller.GetPhotoCustomer)

	testFile := app.Group("/api/v1/file/customer")
	testFile.Get("/photo", controller.GetPhotoCustomer)

}
