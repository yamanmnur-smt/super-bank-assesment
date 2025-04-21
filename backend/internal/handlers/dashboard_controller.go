package handlers

import (
	"net/http"
	"yamanmnur/simple-dashboard/internal/services"
	pkg_response "yamanmnur/simple-dashboard/pkg/responses"

	"github.com/gofiber/fiber/v2"
)

type DashboardController struct {
	Service services.IDashboardService
}

func (c *DashboardController) GetDashboard(ctx *fiber.Ctx) error {

	result, err := c.Service.GetDashboard()
	if err != nil {
		return err
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Get Data",
			Code:    "200",
		},
		Data: result,
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}
