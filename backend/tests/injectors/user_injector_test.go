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

func TestInjectorUserDI(t *testing.T) {
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

	t.Run("should initialize UserRepository and UserService", func(t *testing.T) {
		injectors.InjectorUserDI(dbHandler, container)

		assert.NotNil(t, container.UserRepository, "UserRepository should not be nil")
		assert.NotNil(t, container.UserService, "UserService should not be nil")

		userRepo := container.UserRepository
		assert.NotNil(t, userRepo, "UserRepository should not be nil")
		assert.Equal(t, dbHandler, userRepo.IDbHandler, "UserRepository should use the provided dbHandler")

		userService := container.UserService
		assert.NotNil(t, userService, "userService should not be nil")
		assert.Equal(t, container.UserRepository, userService.Repository, "UserService should use the initialized UserRepository")
	})
}
