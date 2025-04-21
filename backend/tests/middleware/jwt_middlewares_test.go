package middlewares

import (
	"net/http/httptest"
	"testing"
	"yamanmnur/simple-dashboard/internal/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestJwtMiddleware(t *testing.T) {
	// Set up the Fiber app
	app := fiber.New()

	// Mock the secret key
	viper.Set("APP_SECRET_KEY", "test_secret_key")
	secretKey := viper.Get("APP_SECRET_KEY").(string)
	secretKeyByte := []byte(secretKey)

	// Define a route with the middleware
	app.Use(middlewares.JwtMiddleware)
	app.Get("/protected", func(c *fiber.Ctx) error {
		userId := c.Locals("UserId")
		return c.JSON(fiber.Map{"userId": userId})
	})

	t.Run("Missing Authorization Header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer invalid_token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Valid Token", func(t *testing.T) {
		// Create a valid token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "12345",
		})
		tokenString, _ := token.SignedString(secretKeyByte)

		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
