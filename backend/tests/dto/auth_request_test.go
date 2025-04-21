package dto_test

import (
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/requests"

	"github.com/stretchr/testify/assert"
)

func TestLoginDTO(t *testing.T) {
	t.Run("should return error sort when User Name And Password is empty", func(t *testing.T) {
		request := &requests.Login{}

		err := request.Validate()
		assert.Error(t, err, "User Name And Password is empty")

	})

	t.Run("should return error sort when User Name is empty", func(t *testing.T) {
		request := &requests.Login{
			Username: "",
			Password: "password",
		}
		err := request.Validate()
		assert.Error(t, err, "User Name is empty")
	})

	t.Run("should return error sort when Password is empty", func(t *testing.T) {
		request := &requests.Login{
			Username: "username",
			Password: "",
		}
		err := request.Validate()
		assert.Error(t, err, "Password is empty")
	})
}

func TestRegisterDTO(t *testing.T) {
	t.Run("should return error sort when User Name And Password is empty", func(t *testing.T) {
		request := &requests.Register{}

		err := request.Validate()
		assert.Error(t, err, "User Name And Password is empty")

	})

	t.Run("should return error sort when User Name is empty", func(t *testing.T) {
		request := &requests.Register{
			Username: "",
			Password: "password",
		}
		err := request.Validate()
		assert.Error(t, err, "User Name is empty")
	})

	t.Run("should return error sort when Password is empty", func(t *testing.T) {
		request := &requests.Register{
			Username: "username",
			Password: "",
		}
		err := request.Validate()
		assert.Error(t, err, "Password is empty")
	})
}
