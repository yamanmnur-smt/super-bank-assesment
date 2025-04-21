package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDashboardService struct {
	mock.Mock
}

func (m *MockDashboardService) GetDashboard() (data.DashboardData, error) {
	args := m.Called()
	return args.Get(0).(data.DashboardData), args.Error(1)
}

func TestDashboardController_GetDashboard_Success(t *testing.T) {
	// Arrange
	mockService := new(MockDashboardService)
	controller := handlers.DashboardController{Service: mockService}

	mockService.On("GetDashboard").Return(data.DashboardData{
		TotalCards: data.DashboardTotalCard{
			TotalCustomers: "10",
			TotalDeposits:  "10",
			TotalBalance:   "10",
		},
		PieData: []data.ChartData{
			{
				Label: "7 Days",
				Value: "1 Months",
			},
		},

		BarData: []data.ChartData{
			{
				Label: "Jan",
				Value: "10",
			},
		},
	}, nil)

	app := fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		return middlewares.ErrorHandler(ctx, ctx.Next())
	})
	app.Get("/dashboard", controller.GetDashboard)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestDashboardController_GetDashboard_Error(t *testing.T) {
	// Arrange
	mockService := new(MockDashboardService)
	mockService.On("GetDashboard").Return(data.DashboardData{}, errors.New("service error"))

	controller := handlers.DashboardController{Service: mockService}
	app := fiber.New()
	app.Get("/dashboard", controller.GetDashboard)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}
