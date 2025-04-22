package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"yamanmnur/simple-dashboard/internal/middlewares"
	pkg_response "yamanmnur/simple-dashboard/pkg/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestErrorHandler_RecordNotFound(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	// Simulate a route that triggers the error
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return gorm.ErrRecordNotFound
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var response pkg_response.BasicResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.MetaData.Status)
	assert.Equal(t, "Record not found", response.MetaData.Message)
	assert.Equal(t, http.StatusText(http.StatusNotFound), response.MetaData.Code)
}

func TestErrorHandler_InternalServerError(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	// Simulate a route that triggers a generic error
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return errors.New("something went wrong")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var response pkg_response.BasicResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.MetaData.Status)
	assert.Equal(t, "something went wrong", response.MetaData.Message)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), response.MetaData.Code)
}

func TestErrorHandler_MissingPayload(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	// Simulate a route that triggers a generic error
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return errors.New("missing payload body")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response pkg_response.BasicResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.MetaData.Status)
	assert.Equal(t, "missing payload body", response.MetaData.Message)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), response.MetaData.Code)
}

func TestErrorHandler_CredentialWrong(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	// Simulate a route that triggers a generic error
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return errors.New("credential wrong")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response pkg_response.BasicResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.MetaData.Status)
	assert.Equal(t, "credential wrong", response.MetaData.Message)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), response.MetaData.Code)
}
