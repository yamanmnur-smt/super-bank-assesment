package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/services"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_response "yamanmnur/simple-dashboard/pkg/responses"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service services.IAuthService
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var requestBody requests.Login

	if err := ctx.BodyParser(&requestBody); err != nil {
		return pkg_data.InvalidReqPayloadError{Message: "invalid request payload"}

	}

	if err := requestBody.Validate(); err != nil {
		return pkg_data.InvalidReqPayloadError{Message: err.Error()}
	}

	res, err := c.Service.Login(&requestBody)

	if err != nil {
		return err
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Login User",
			Code:    "200",
		},
		Data: res,
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}

func (c *AuthController) Profile(ctx *fiber.Ctx) error {

	stringId := fmt.Sprintf("%v", ctx.Locals("UserId"))

	if stringId == "" {
		return errors.New("missing user id")
	}

	idUint, err := strconv.ParseUint(stringId, 10, 32)

	if err != nil {
		return errors.New("invalid user id")
	}

	res, err := c.Service.Profile(uint(idUint))

	if err != nil {
		return err
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Get Profile User",
			Code:    "200",
		},
		Data: res,
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}
