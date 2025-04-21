package mocks_test

import (
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindById(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(username string) (models.User, error) {
	args := m.Called(username)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Create(userData data.UserData) (models.User, error) {
	args := m.Called(userData)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Update(userData data.UserData) (models.User, error) {
	args := m.Called(userData)
	return args.Get(0).(models.User), args.Error(1)
}
