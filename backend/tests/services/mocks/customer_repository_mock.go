package mocks_test

import (
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"

	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) FindById(id uint) (models.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Detail(id uint, detail_data *models.Customer) error {
	args := m.Called(id, detail_data)
	return args.Error(0)
}

func (m *MockCustomerRepository) Create(customerData *models.Customer) error {
	args := m.Called(customerData)
	return args.Error(0)
}

func (m *MockCustomerRepository) Update(customerData *data.CustomerData) (models.Customer, error) {
	args := m.Called(customerData)
	return args.Get(0).(models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) UpdatePatch(customerData *data.CustomerData) (models.Customer, error) {
	args := m.Called(customerData)
	return args.Get(0).(models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCustomerRepository) GetCustomersWithPagination(pageRequest pkg_requests.PageRequest) (pkg_data.PaginateData[data.CustomerData], error) {
	args := m.Called(pageRequest)
	return args.Get(0).(pkg_data.PaginateData[data.CustomerData]), args.Error(1)
}
