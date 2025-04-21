package injectors

import (
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/services"
	"yamanmnur/simple-dashboard/pkg/db"
)

func InjectorAuthDI(dbHandler *db.IDbHandler, container *Container) {
	container.AuthService = &services.AuthService{UserService: container.UserService}
	container.AuthController = &handlers.AuthController{Service: container.AuthService}
}
