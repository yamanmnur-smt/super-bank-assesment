package injectors

import (
	"testing"
	"yamanmnur/simple-dashboard/internal/injectors"
	"yamanmnur/simple-dashboard/pkg/db"
	"yamanmnur/simple-dashboard/pkg/util"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestInitDI(t *testing.T) {
	t.Skip("Temporary Skip")
	app := fiber.New()
	viper.SetConfigFile("../../.env")
	viper.ReadInConfig()

	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	gormDb, _ := gorm.Open(dialector, &gorm.Config{})

	dbHandler := &db.IDbHandler{
		DB: gormDb,
	}

	client := util.InitMinio()

	// Call the function
	injectors.InitDI(app, dbHandler, client)

	// Assertions
	assert.NotNil(t, app, "Fiber app should not be nil")

	// Check if routes are registered
	routes := app.GetRoutes()
	assert.NotEmpty(t, routes, "Routes should be registered")
}
