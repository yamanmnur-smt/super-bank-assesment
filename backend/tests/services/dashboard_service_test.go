package services_test

import (
	"errors"
	"testing"

	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetTotalCard() (data.DashboardTotalCard, error) {
	args := m.Called()
	return args.Get(0).(data.DashboardTotalCard), args.Error(1)
}

func (m *MockRepository) GetPieData() ([]data.ChartData, error) {
	args := m.Called()
	return args.Get(0).([]data.ChartData), args.Error(1)
}

func (m *MockRepository) GetMonthlyRegisteredCustomers() ([]data.ChartData, error) {
	args := m.Called()
	return args.Get(0).([]data.ChartData), args.Error(1)
}

func TestDashboardService_GetDashboard_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.DashboardService{Repository: mockRepo}

	mockTotalCard := data.DashboardTotalCard{
		TotalBalance:  "1000000",
		TotalDeposits: "500000",
	}
	mockPieData := []data.ChartData{
		{Label: "Category A", Value: "300000"},
		{Label: "Category B", Value: "700000"},
	}
	mockBarData := []data.ChartData{
		{Label: "January", Value: "10"},
		{Label: "February", Value: "20"},
	}

	mockRepo.On("GetTotalCard").Return(mockTotalCard, nil)
	mockRepo.On("GetPieData").Return(mockPieData, nil)
	mockRepo.On("GetMonthlyRegisteredCustomers").Return(mockBarData, nil)

	result, err := service.GetDashboard()

	assert.NoError(t, err)
	assert.Equal(t, "Rp1.000.000", result.TotalCards.TotalBalance)
	assert.Equal(t, "Rp500.000", result.TotalCards.TotalDeposits)
	assert.Equal(t, "Rp300.000", result.PieData[0].Value)
	assert.Equal(t, "Rp700.000", result.PieData[1].Value)
	assert.Equal(t, mockBarData, result.BarData)

	mockRepo.AssertExpectations(t)
}

func TestDashboardService_GetDashboard_ErrorOnTotalCard(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.DashboardService{Repository: mockRepo}

	mockRepo.On("GetTotalCard").Return(data.DashboardTotalCard{}, errors.New("database error"))

	_, err := service.GetDashboard()

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDashboardService_GetDashboard_ErrorOnPieData(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.DashboardService{Repository: mockRepo}

	mockTotalCard := data.DashboardTotalCard{
		TotalBalance:  "1000000",
		TotalDeposits: "500000",
	}

	mockRepo.On("GetTotalCard").Return(mockTotalCard, nil)
	mockRepo.On("GetPieData").Return([]data.ChartData{}, errors.New("database error"))

	_, err := service.GetDashboard()

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDashboardService_GetDashboard_ErrorOnBarData(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.DashboardService{Repository: mockRepo}

	mockTotalCard := data.DashboardTotalCard{
		TotalBalance:  "1000000",
		TotalDeposits: "500000",
	}
	mockPieData := []data.ChartData{
		{Label: "Category A", Value: "300000"},
		{Label: "Category B", Value: "700000"},
	}

	mockRepo.On("GetTotalCard").Return(mockTotalCard, nil)
	mockRepo.On("GetPieData").Return(mockPieData, nil)
	mockRepo.On("GetMonthlyRegisteredCustomers").Return([]data.ChartData{}, errors.New("database error"))

	_, err := service.GetDashboard()

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	mockRepo.AssertExpectations(t)
}
