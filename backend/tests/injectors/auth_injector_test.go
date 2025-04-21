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

func TestInjectorAuthDI(t *testing.T) {
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

	t.Run("should initialize AuthService and AuthController", func(t *testing.T) {
		injectors.InjectorUserDI(dbHandler, container)
		injectors.InjectorCustomerDI(dbHandler, container)
		injectors.InjectorAuthDI(dbHandler, container)

		assert.NotNil(t, container.AuthService, "AuthService should not be nil")
		assert.NotNil(t, container.AuthController, "AuthController should not be nil")

		authService := container.AuthService
		assert.NotNil(t, authService, "AuthService should not be nil")
		assert.Equal(t, container.UserService, authService.UserService, "AuthService should use the provided dbHandler")

		authController := container.AuthController
		assert.NotNil(t, authController, "authController should not be nil")
		assert.Equal(t, container.AuthService, authController.Service, "AuthController should use the initialized AuthService")
	})
}
