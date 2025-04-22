package middlewares

import (
	"crypto/rand"
	"crypto/rsa"
	"net/http/httptest"
	"testing"
	"time"
	"yamanmnur/simple-dashboard/internal/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func generateRS256Token() (string, error) {
	// Generate RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "fake-user-id",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign with private key
	return token.SignedString(privateKey)
}

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

	t.Run("Unauthorized Signing Method", func(t *testing.T) {
		// Create a token with an invalid signing method
		token, err := generateRS256Token()
		assert.NoError(t, err)

		req := httptest.NewRequest("GET", "/protected", nil)

		req.Header.Set("Authorization", "Bearer "+token)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}
