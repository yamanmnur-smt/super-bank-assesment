package injectors

import (
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/internal/services"
	"yamanmnur/simple-dashboard/pkg/db"
)

func InjectorUserDI(dbHandler *db.IDbHandler, container *Container) {
	container.UserRepository = &repositories.UserRepository{IDbHandler: dbHandler}
	container.UserService = &services.UserService{Repository: container.UserRepository}
}
