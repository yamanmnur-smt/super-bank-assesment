package injectors

import (
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/internal/services"
	"yamanmnur/simple-dashboard/pkg/db"
)

func InjectorCustomerDI(dbHandler *db.IDbHandler, container *Container) {
	container.CustomerRepository = &repositories.CustomerRepository{IDbHandler: dbHandler}
	container.CustomerService = &services.CustomerService{Repository: container.CustomerRepository, MinioClient: container.MinioClient}
	container.CustomerController = &handlers.CustomerController{Service: container.CustomerService}
}
