package middlewares

import (
	"errors"
	"net/http"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_response "yamanmnur/simple-dashboard/pkg/responses"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {

		var statusCode int
		var message string
		var invalidReqPayloadError pkg_data.InvalidReqPayloadError
		invalidUserId := errors.New("invalid user id")
		missionUserId := errors.New("missing user id")
		missingPayload := errors.New("missing payload body")
		credentialWrong := errors.New("credential wrong")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			message = "Record not found"
		} else if errors.As(err, &invalidReqPayloadError) {
			statusCode = http.StatusBadRequest
			message = err.Error()
		} else if err.Error() == missionUserId.Error() {
			statusCode = http.StatusUnauthorized
			message = err.Error()
		} else if err.Error() == invalidUserId.Error() {
			statusCode = http.StatusBadRequest
			message = err.Error()
		} else if err.Error() == missingPayload.Error() {
			statusCode = http.StatusBadRequest
			message = err.Error()
		} else if err.Error() == credentialWrong.Error() {
			statusCode = http.StatusBadRequest
			message = "credential wrong"
		} else {
			statusCode = http.StatusInternalServerError
			message = err.Error()
		}

		response := pkg_response.BasicResponse{
			MetaData: pkg_response.MetaData{
				Status:  "error",
				Message: message,
				Code:    http.StatusText(statusCode),
			},
		}

		return ctx.Status(statusCode).JSON(response)
	}

	return nil
}
