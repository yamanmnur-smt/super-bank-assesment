package injectors

import (
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/internal/services"
	"yamanmnur/simple-dashboard/pkg/db"
	"yamanmnur/simple-dashboard/pkg/util"
	"yamanmnur/simple-dashboard/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type Container struct {
	UserRepository      *repositories.UserRepository
	UserService         *services.UserService
	AuthService         *services.AuthService
	CustomerRepository  *repositories.CustomerRepository
	CustomerService     *services.CustomerService
	DashboardRepository *repositories.DashboardRepository
	DashboardService    *services.DashboardService

	CustomerController  *handlers.CustomerController
	DashboardController *handlers.DashboardController
	AuthController      *handlers.AuthController

	MinioClient *util.MinioClient
}

func InitDI(app *fiber.App, dbHandler *db.IDbHandler, minioClient *minio.Client) {
	container := &Container{}

	InjectorMinioDI(minioClient, container)

	InjectorUserDI(dbHandler, container)

	InjectorAuthDI(dbHandler, container)

	InjectorCustomerDI(dbHandler, container)

	InjectorDashboardDI(dbHandler, container)

	routes.InitAuthRoutes(app, container.AuthController)
	routes.InitCustomerRoute(app, container.CustomerController)
	routes.InitDashboardRoutes(app, container.DashboardController)
}
