package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(req *requests.Login) (data.JwtToken, error) {
	args := m.Called(req)
	return args.Get(0).(data.JwtToken), args.Error(1)
}

func (m *MockAuthService) Profile(userID uint) (data.UserProfileData, error) {
	args := m.Called(userID)
	return args.Get(0).(data.UserProfileData), args.Error(1)
}

func (m *MockAuthService) GenerateToken(userId string) (string, error) {
	args := m.Called(userId)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockAuthService) Register(request *requests.Register) (data.JwtToken, error) {
	args := m.Called(request)
	return args.Get(0).(data.JwtToken), args.Error(1)
}

func TestAuthController_Login(t *testing.T) {

	mockService := new(MockAuthService)
	controller := handlers.AuthController{Service: mockService}
	app := fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		return middlewares.ErrorHandler(ctx, ctx.Next())
	})
	t.Run("success", func(t *testing.T) {

		app.Post("/login", controller.Login)

		mockRequest := requests.Login{Username: "username", Password: "password"}
		mockResponse := data.JwtToken{
			User: data.UserProfileData{
				Id:       1,
				Username: "username",
				Name:     "name",
			},
			Token: "mocktoken",
		}
		mockService.On("Login", &mockRequest).Return(mockResponse, nil)
		jsonBytes, _ := json.Marshal(mockRequest)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(string(jsonBytes)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("validate request body", func(t *testing.T) {

		app.Post("/login", controller.Login)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{"usernae":"123"}`))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("invalid request body", func(t *testing.T) {

		app.Post("/login", controller.Login)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`notvalidreqbody`))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service error", func(t *testing.T) {
		app := fiber.New(fiber.Config{
			ErrorHandler: middlewares.ErrorHandler,
		})
		app.Post("/login", controller.Login)
		mockRequest := requests.Login{Username: "test@example.com", Password: "password"}
		mockService.On("Login", &mockRequest).Return(data.JwtToken{}, errors.New("service error"))
		jsonBytes, _ := json.Marshal(mockRequest)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(string(jsonBytes)))

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestAuthController_Profile(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockAuthService)
		controller := handlers.AuthController{Service: mockService}
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		app.Get("/profile", func(ctx *fiber.Ctx) error {
			ctx.Locals("UserId", "1")
			return controller.Profile(ctx)
		})

		mockResponse := data.UserProfileData{
			Id:       1,
			Username: "username",
			Name:     "name",
		}
		mockService.On("Profile", uint(1)).Return(mockResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("missing user ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockAuthService)
		controller := handlers.AuthController{Service: mockService}

		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		app.Get("/profile", func(ctx *fiber.Ctx) error {
			ctx.Locals("UserId", "")
			return controller.Profile(ctx)
		})
		req := httptest.NewRequest(http.MethodGet, "/profile", nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockAuthService)
		controller := handlers.AuthController{Service: mockService}
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		app.Get("/profile", func(ctx *fiber.Ctx) error {
			ctx.Locals("UserId", "userid")
			return controller.Profile(ctx)
		})
		req := httptest.NewRequest(http.MethodGet, "/profile", nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockAuthService)
		controller := handlers.AuthController{Service: mockService}
		app.Get("/profile", func(ctx *fiber.Ctx) error {
			ctx.Locals("UserId", "1")
			return controller.Profile(ctx)
		})
		mockService.On("Profile", uint(1)).Return(data.UserProfileData{}, errors.New("service error"))

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
