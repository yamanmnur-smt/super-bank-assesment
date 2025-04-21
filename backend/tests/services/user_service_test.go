package services_test

import (
	"errors"
	"testing"

	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/internal/services"
	mocks_test "yamanmnur/simple-dashboard/tests/services/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	mockUserModel := models.User{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "password",
	}
	mockUserModel.ID = 1

	expectedUserData := data.UserData{
		Id:       1,
		Name:     "John Doe",
		Username: "johndoe",
		Password: "password",
	}

	mockRepo.On("FindById", uint(1)).Return(mockUserModel, nil)

	result, err := service.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserData, result)
	mockRepo.AssertExpectations(t)
}

func TestFindById_Error(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	mockRepo.On("FindById", uint(1)).Return(models.User{}, errors.New("user not found"))

	result, err := service.FindById(1)

	assert.Error(t, err)
	assert.Equal(t, data.UserData{}, result)
	mockRepo.AssertExpectations(t)
}

func TestFindByUsername(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	mockUserModel := models.User{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "password",
	}
	mockUserModel.ID = 1

	mockRepo.On("FindByUsername", "johndoe").Return(mockUserModel, nil)

	expectedUser := data.UserData{Id: 1, Name: "John Doe", Username: "johndoe", Password: "password"}

	result, err := service.FindByUsername("johndoe")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestFindByUsername_Error(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	mockRepo.On("FindByUsername", "johndoe").Return(models.User{}, errors.New("user not found"))

	result, err := service.FindByUsername("johndoe")

	assert.Error(t, err)
	assert.Equal(t, data.UserData{}, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	// Mocking the input and expected output
	mockUserModel := models.User{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "password",
	}
	mockUserModel.ID = 1

	inputUser := data.UserData{Name: "John Doe", Username: "johndoe", Password: "password"}

	mockRepo.On("Create", inputUser).Return(mockUserModel, nil)

	expectedUser := data.UserData{Id: 1, Name: "John Doe", Username: "johndoe", Password: "password"}

	result, err := service.Create(inputUser)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate_Error(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	inputUser := data.UserData{Name: "John Doe", Username: "johndoe", Password: "password"}
	mockRepo.On("Create", inputUser).Return(models.User{}, errors.New("failed to create user"))

	result, err := service.Create(inputUser)

	assert.Error(t, err)
	assert.Equal(t, data.UserData{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	// Mocking the input and expected output
	mockUserModel := models.User{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "newpassword",
	}
	mockUserModel.ID = 1

	inputUser := data.UserData{Id: 1, Name: "John Doe", Username: "johndoe", Password: "newpassword"}
	mockRepo.On("Update", inputUser).Return(mockUserModel, nil)

	expectedUser := data.UserData{Id: 1, Name: "John Doe", Username: "johndoe", Password: "newpassword"}

	result, err := service.Update(inputUser)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_Error(t *testing.T) {
	mockRepo := new(mocks_test.MockUserRepository)
	service := services.UserService{Repository: mockRepo}

	inputUser := data.UserData{Id: 1, Name: "John Doe", Username: "johndoe", Password: "newpassword"}
	mockRepo.On("Update", inputUser).Return(models.User{}, errors.New("failed to update user"))

	result, err := service.Update(inputUser)

	assert.Error(t, err)
	assert.Equal(t, data.UserData{}, result)
	mockRepo.AssertExpectations(t)
}
