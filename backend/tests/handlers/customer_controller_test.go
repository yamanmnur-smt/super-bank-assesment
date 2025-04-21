package handlers_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/handlers"
	"yamanmnur/simple-dashboard/internal/middlewares"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) Detail(id uint) (data.CustomerDetailData, error) {
	args := m.Called(id)
	return args.Get(0).(data.CustomerDetailData), args.Error(1)
}

func (m *MockCustomerService) GetCustomersWithPagination(req pkg_requests.PageRequest) (*pkg_data.PaginateResponse[data.CustomerData], error) {
	args := m.Called(req)
	return args.Get(0).(*pkg_data.PaginateResponse[data.CustomerData]), args.Error(1)
}

func (m *MockCustomerService) UpdatePhotoCustomer(req *requests.CustomerPhotoRequest) (data.CustomerData, error) {
	args := m.Called(req)
	return args.Get(0).(data.CustomerData), args.Error(1)
}

func (m *MockCustomerService) UploadPhotoCustomer(baseURL string, file *multipart.FileHeader) (string, error) {
	args := m.Called(baseURL, file)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockCustomerService) GetPhotoCustomer(filename string) ([]byte, error) {
	args := m.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCustomerService) Create(req *requests.CustomerRequest) (data.CustomerDetailData, error) {
	args := m.Called(req)
	return args.Get(0).(data.CustomerDetailData), args.Error(1)
}

func (m *MockCustomerService) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCustomerService) FindById(id uint) (data.CustomerData, error) {
	args := m.Called(id)
	return args.Get(0).(data.CustomerData), args.Error(1)
}

func (m *MockCustomerService) Update(userData *data.CustomerData) (data.CustomerData, error) {
	args := m.Called(userData)
	return args.Get(0).(data.CustomerData), args.Error(1)
}

func TestGetCustomerByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customer/:id", controller.GetCustomerByID)

		mockCustomer := data.CustomerDetailData{Id: 1, Name: "John Doe"}
		mockService.On("Detail", uint(1)).Return(mockCustomer, nil)

		req := httptest.NewRequest(http.MethodGet, "/customer/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "Detail", uint(1))
	})

	t.Run("Invalid Customer ID", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customer/:id", controller.GetCustomerByID)

		req := httptest.NewRequest(http.MethodGet, "/customer/invalid", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Customer Service Return Error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customer/:id", controller.GetCustomerByID)

		mockService.On("Detail", uint(1)).Return(data.CustomerDetailData{}, errors.New("error db"))

		req := httptest.NewRequest(http.MethodGet, "/customer/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertCalled(t, "Detail", uint(1))

	})
}

func TestGetCustomersWithPagination(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customers", controller.GetCustomersWithPagination)

		mockResponse := pkg_data.PaginateResponse[data.CustomerData]{
			Data: []data.CustomerData{{Id: 1, Name: "John Doe"}},
			PageData: pkg_data.PageData{
				Page:      1,
				Limit:     10,
				TotalRows: 1,
			},
		}
		mockService.On("GetCustomersWithPagination", mock.Anything).Return(&mockResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/customers?page_number=0&page_size=10", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "GetCustomersWithPagination", mock.Anything)
	})

	t.Run("Query Param Error", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customers", controller.GetCustomersWithPagination)

		req := httptest.NewRequest(http.MethodGet, "/customers?page_number=1&page_size=invalid_value", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Validate Query Param Error", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customers", controller.GetCustomersWithPagination)

		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Customer Service Return Error", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}
		mockService.On("GetCustomersWithPagination", mock.Anything).Return(&pkg_data.PaginateResponse[data.CustomerData]{}, errors.New("db error"))

		app.Get("/customers", controller.GetCustomersWithPagination)

		req := httptest.NewRequest(http.MethodGet, "/customers?page_number=0&page_size=10", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertCalled(t, "GetCustomersWithPagination", mock.Anything)

	})
}

func TestUploadPhotoCustomer(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer/photo", controller.UploadPhotoCustomer)

		mockService.On("UploadPhotoCustomer", mock.Anything, mock.Anything).Return("http://example.com/photo.jpg", nil)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("fileUpload", "test.png")
		part.Write([]byte("dummy file content"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/customer/photo", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "UploadPhotoCustomer", mock.Anything, mock.Anything)
	})

	t.Run("Success https", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer/photo", controller.UploadPhotoCustomer)

		mockService.On("UploadPhotoCustomer", "https://example.com", mock.Anything).Return("https://example.com/photo.jpg", nil)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("fileUpload", "test.png")
		part.Write([]byte("dummy file content"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/customer/photo", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("X-Forwarded-Proto", "https")
		req.Host = "example.com"
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "UploadPhotoCustomer", "https://example.com", mock.Anything)
	})

	t.Run("Missing File", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer/photo", controller.UploadPhotoCustomer)

		req := httptest.NewRequest(http.MethodPost, "/customer/photo", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("Service Error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer/photo", controller.UploadPhotoCustomer)

		mockService.On("UploadPhotoCustomer", mock.Anything, mock.Anything).Return("", errors.New("upload error"))

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("fileUpload", "test.png")
		part.Write([]byte("dummy file content"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/customer/photo", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertCalled(t, "UploadPhotoCustomer", mock.Anything, mock.Anything)
	})

}

func TestUpdatePhotoCustomer(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Put("/customer/photo", controller.UpdatePhotoCustomer)

		mockCustomer := data.CustomerData{Id: 1, Name: "John Doe"}
		mockService.On("UpdatePhotoCustomer", mock.Anything).Return(mockCustomer, nil)

		body := bytes.NewBufferString(`{"photo_url": "http://example.com/photo.jpg"}`)
		req := httptest.NewRequest(http.MethodPut, "/customer/photo", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "UpdatePhotoCustomer", mock.Anything)
	})

	t.Run("Service Return Error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Put("/customer/photo", controller.UpdatePhotoCustomer)

		mockService.On("UpdatePhotoCustomer", mock.Anything).Return(data.CustomerData{}, errors.New("service error"))

		body := bytes.NewBufferString(`{"photo_url": "http://example.com/photo.jpg"}`)
		req := httptest.NewRequest(http.MethodPut, "/customer/photo", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertCalled(t, "UpdatePhotoCustomer", mock.Anything)
	})

	t.Run("Invalid Requests Payload", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Put("/customer/photo", controller.UpdatePhotoCustomer)

		req := httptest.NewRequest(http.MethodPut, "/customer/photo", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

}

func TestGetPhotoCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customers/get-photo", controller.GetPhotoCustomer)

		mockService.On("GetPhotoCustomer", mock.Anything).Return([]byte("test"), nil)

		req := httptest.NewRequest(http.MethodGet, "/customers/get-photo?filename=photo.png", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "GetPhotoCustomer", mock.Anything)
	})

	t.Run("service error", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Get("/customers/get-photo", controller.GetPhotoCustomer)

		mockService.On("GetPhotoCustomer", mock.Anything).Return([]byte(""), errors.New("service error"))

		req := httptest.NewRequest(http.MethodGet, "/customers/get-photo?filename=photo.png", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertCalled(t, "GetPhotoCustomer", mock.Anything)
	})
}

func TestCreateCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer", controller.Create)

		mockCustomer := data.CustomerDetailData{Id: 1, Name: "John Doe"}
		mockService.On("Create", mock.Anything).Return(mockCustomer, nil)

		body := bytes.NewBufferString(`{
			"photo" : "",
			"name" : "Asep",
			"email" : "asep@mail.co,",
			"phone" : "08823423",
			"address" : "Bandung",
			"gender" : "male",
			"account_purpose" : "",
			"source_of_income" : "",
			"income_per_month" : "",
			"jobs" : "",
			"position" : "",
			"industries" : "",
			"company_name" : "",
			"address_company" : "",
			"username" : "asep"
		}`)
		req := httptest.NewRequest(http.MethodPost, "/customer", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertCalled(t, "Create", mock.Anything)
	})

	t.Run("invalid request payload", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer", controller.Create)

		req := httptest.NewRequest(http.MethodPost, "/customer", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("request validated", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			return middlewares.ErrorHandler(ctx, ctx.Next())
		})
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer", controller.Create)

		body := bytes.NewBufferString(`{
			"photo" : "",
			"name" : "Asep",
			"phone" : "08823423",
			"address" : "Bandung",
			"gender" : "male",
			"account_purpose" : "",
			"source_of_income" : "",
			"income_per_month" : "",
			"jobs" : "",
			"position" : "",
			"industries" : "",
			"company_name" : "",
			"address_company" : "",
			"username" : "asep"
		}`)
		req := httptest.NewRequest(http.MethodPost, "/customer", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockCustomerService)
		controller := handlers.CustomerController{Service: mockService}

		app.Post("/customer", controller.Create)

		mockService.On("Create", mock.Anything).Return(data.CustomerDetailData{}, errors.New("service error"))

		body := bytes.NewBufferString(`{
			"photo" : "",
			"name" : "Asep",
			"email" : "asep@mail.co,",
			"phone" : "08823423",
			"address" : "Bandung",
			"gender" : "male",
			"account_purpose" : "",
			"source_of_income" : "",
			"income_per_month" : "",
			"jobs" : "",
			"position" : "",
			"industries" : "",
			"company_name" : "",
			"address_company" : "",
			"username" : "asep"
		}`)
		req := httptest.NewRequest(http.MethodPost, "/customer", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertCalled(t, "Create", mock.Anything)
	})

}
