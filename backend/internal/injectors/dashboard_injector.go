package injectors

import (
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/internal/services"
	"yamanmnur/simple-dashboard/pkg/db"
)

func InjectorDashboardDI(dbHandler *db.IDbHandler, container *Container) {
	container.DashboardRepository = &repositories.DashboardRepository{IDbHandler: dbHandler}
	container.DashboardService = &services.DashboardService{Repository: container.DashboardRepository}
	container.DashboardController = &handlers.DashboardController{Service: container.DashboardService}
}
