package injectors_test

import (
	"testing"
	"yamanmnur/simple-dashboard/internal/injectors"
	"yamanmnur/simple-dashboard/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestInjectorCustomerDI(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	gormDb, _ := gorm.Open(dialector, &gorm.Config{})

	dbHandler := &db.IDbHandler{
		DB: gormDb,
	}

	container := &injectors.Container{}

	t.Run("should initialize CustomerRepository and CustomerService", func(t *testing.T) {
		injectors.InjectorUserDI(dbHandler, container)
		injectors.InjectorCustomerDI(dbHandler, container)

		assert.NotNil(t, container.CustomerRepository, "CustomerRepository should not be nil")
		assert.NotNil(t, container.CustomerService, "CustomerService should not be nil")

		userRepo := container.CustomerRepository
		assert.NotNil(t, userRepo, "CustomerRepository should not be nil")
		assert.Equal(t, dbHandler, userRepo.IDbHandler, "CustomerRepository should use the provided dbHandler")

		userService := container.CustomerService
		assert.NotNil(t, userService, "userService should not be nil")
		assert.Equal(t, container.CustomerRepository, userService.Repository, "CustomerService should use the initialized CustomerRepository")

		customerController := container.CustomerController
		assert.NotNil(t, customerController, "CustomerController should not be nil")
		assert.Equal(t, container.CustomerService, customerController.Service, "CustomerController should use the initialized CustomerService")
	})
}
