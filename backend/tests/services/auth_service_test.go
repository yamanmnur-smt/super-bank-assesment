package services_test

import (
	"errors"
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(user data.UserData) (data.UserData, error) {
	args := m.Called(user)
	return args.Get(0).(data.UserData), args.Error(1)
}

func (m *MockUserService) FindByUsername(username string) (data.UserData, error) {
	args := m.Called(username)
	return args.Get(0).(data.UserData), args.Error(1)
}

func (m *MockUserService) FindById(userId uint) (data.UserData, error) {
	args := m.Called(userId)
	return args.Get(0).(data.UserData), args.Error(1)
}

func (m *MockUserService) Update(userData data.UserData) (data.UserData, error) {
	args := m.Called(userData)
	return args.Get(0).(data.UserData), args.Error(1)
}

func TestAuthService_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserService := new(MockUserService)
		authService := &services.AuthService{UserService: mockUserService}
		viper.Set("APP_SECRET_KEY", "testsecretkey") // Ensure APP_SECRET_KEY is set for the test

		mockUser := data.UserData{
			Id:       1,
			Name:     "John Doe",
			Username: "johndoe",
			Password: "hashedpassword",
		}

		mockUserService.On("Create", mock.Anything).Return(mockUser, nil)

		request := &requests.Register{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "password123",
		}

		token, err := authService.Register(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, token.Token)
		assert.Equal(t, mockUser.Id, token.User.Id)
		assert.Equal(t, mockUser.Name, token.User.Name)
		assert.Equal(t, mockUser.Username, token.User.Username)
	})

	t.Run("user service error", func(t *testing.T) {
		mockUserService := new(MockUserService)
		authService := &services.AuthService{UserService: mockUserService}
		viper.Set("APP_SECRET_KEY", "testsecretkey")

		mockUserService.On("Create", mock.Anything).Return(data.UserData{}, errors.New("service error"))

		request := &requests.Register{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "password123",
		}

		token, err := authService.Register(request)

		assert.Error(t, err)
		assert.Empty(t, token.Token)
	})

}

func TestAuthService_Login(t *testing.T) {

	t.Run("Success Login", func(t *testing.T) {
		mockUserService := new(MockUserService)

		authService := &services.AuthService{UserService: mockUserService}

		viper.Set("APP_SECRET_KEY", "testsecretkey") // Ensure APP_SECRET_KEY is set for the test

		mockUser := data.UserData{
			Id:       1,
			Name:     "John Doe",
			Username: "johndoe",
			Password: "$2a$10$HfKwaDheE9I8tVRI/jWXOuXqgJf64IA6MTkd.kRtJflMXMAB0jwyu", // bcrypt hash for "password123"
		}

		mockUserService.On("FindByUsername", "johndoe").Return(mockUser, nil)

		request := &requests.Login{
			Username: "johndoe",
			Password: "password",
		}

		token, err := authService.Login(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, token.Token)
		assert.Equal(t, mockUser.Id, token.User.Id)
		assert.Equal(t, mockUser.Name, token.User.Name)
		assert.Equal(t, mockUser.Username, token.User.Username)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUserService := new(MockUserService)

		authService := &services.AuthService{UserService: mockUserService}

		mockUser := data.UserData{}

		mockUserService.On("FindByUsername", "johndoe").Return(mockUser, gorm.ErrRecordNotFound)

		request := &requests.Login{
			Username: "johndoe",
			Password: "password",
		}

		token, err := authService.Login(request)
		assert.Error(t, err)
		assert.Equal(t, data.JwtToken{}, token)

	})

	t.Run("Password Not Valid", func(t *testing.T) {
		mockUserService := new(MockUserService)
		authService := &services.AuthService{UserService: mockUserService}

		viper.Set("APP_SECRET_KEY", "testsecretkey") // Ensure APP_SECRET_KEY is set for the test

		mockUser := data.UserData{
			Id:       1,
			Name:     "John Doe",
			Username: "johndoe",
			Password: "$2a$10$HfKwaDheE9I8tVRI/jWXOuXqgJf64IA6MTkd.kRtJflMXMAB0jwyu", // bcrypt hash for "password123"
		}

		mockUserService.On("FindByUsername", "johndoe").Return(mockUser, nil)

		request := &requests.Login{
			Username: "johndoe",
			Password: "password23",
		}

		token, err := authService.Login(request)

		assert.Error(t, err)
		assert.Equal(t, "credential wrong", err.Error())
		assert.Equal(t, data.JwtToken{}, token)

	})
}

func TestAuthService_Profile(t *testing.T) {
	t.Run("Success Found User", func(t *testing.T) {
		mockUserService := new(MockUserService)
		authService := &services.AuthService{UserService: mockUserService}

		mockUser := data.UserData{
			Id:       1,
			Name:     "John Doe",
			Username: "johndoe",
		}

		mockUserService.On("FindById", uint(1)).Return(mockUser, nil)

		profile, err := authService.Profile(1)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Id, profile.Id)
		assert.Equal(t, mockUser.Name, profile.Name)
		assert.Equal(t, mockUser.Username, profile.Username)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUserService := new(MockUserService)

		authService := &services.AuthService{UserService: mockUserService}

		mockUserService.On("FindById", uint(1)).Return(data.UserData{}, gorm.ErrRecordNotFound)

		profile, err := authService.Profile(1)

		assert.Error(t, err)
		assert.Equal(t, data.UserProfileData{}, profile)
	})
}

func TestAuthService_GenerateToken(t *testing.T) {
	authService := &services.AuthService{}

	viper.Set("APP_SECRET_KEY", "testsecretkey")

	token, err := authService.GenerateToken("1")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("testsecretkey"), nil
	})

	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, "1", claims["sub"])
	assert.NotNil(t, claims["exp"])
	assert.Equal(t, "jwtservice:3321", claims["iss"])
}
