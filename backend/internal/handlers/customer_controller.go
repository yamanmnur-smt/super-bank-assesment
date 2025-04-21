package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/services"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"
	pkg_response "yamanmnur/simple-dashboard/pkg/responses"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	Service services.ICustomerService
}

func GetBaseURL(c *fiber.Ctx) string {
	scheme := "http"
	if c.Protocol() == "https" || c.Secure() {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, c.Hostname())
}

func (c *CustomerController) GetCustomerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid customer ID")
	}

	customer, err := c.Service.Detail(uint(idUint))
	if err != nil {
		return err
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Get Data",
			Code:    "200",
		},
		Data: customer,
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}

func (c *CustomerController) GetCustomersWithPagination(ctx *fiber.Ctx) error {
	var pageRequest pkg_requests.PageRequest

	if err := ctx.QueryParser(&pageRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	if err := pageRequest.Validate(); err != nil {
		return pkg_data.InvalidReqPayloadError{Message: err.Error()}
	}

	result, err := c.Service.GetCustomersWithPagination(pageRequest)
	if err != nil {
		response := pkg_response.BasicResponse{
			MetaData: pkg_response.MetaData{
				Status:  "error",
				Message: err.Error(),
				Code:    "500",
			},
		}
		ctx.Status(http.StatusInternalServerError).JSON(response)
		return err
	}

	response := pkg_response.PaginateResponse[data.CustomerData]{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Retrieve Data",
			Code:    "200",
		},
		Data: result.Data,

		PageData: pkg_data.PageData{
			Page:      int(pageRequest.PageNumber),
			Limit:     int(pageRequest.PageSize),
			TotalRows: result.PageData.TotalRows,
		},
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}

func (c *CustomerController) UpdatePhotoCustomer(ctx *fiber.Ctx) error {
	var requestBody requests.CustomerPhotoRequest

	if err := ctx.BodyParser(&requestBody); err != nil {
		return pkg_data.InvalidReqPayloadError{Message: "missing payload body"}
	}

	customer, err := c.Service.UpdatePhotoCustomer(&requestBody)
	if err != nil {
		return err
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Update Photo Customer",
			Code:    "200",
		},
		Data: customer,
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}

func (c *CustomerController) UploadPhotoCustomer(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("fileUpload")

	if err != nil {
		return err
	}

	fileinfo, err := c.Service.UploadPhotoCustomer(GetBaseURL(ctx), file)
	if err != nil {
		return err
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Status:  "success",
			Message: "Success To Update Photo Customer",
			Code:    "200",
		},
		Data: fileinfo,
	}

	ctx.Status(http.StatusOK).JSON(response)
	return nil
}

func (c *CustomerController) GetPhotoCustomer(ctx *fiber.Ctx) error {
	filename := ctx.Query("filename")

	file, err := c.Service.GetPhotoCustomer(filename)
	if err != nil {
		return err
	}

	ctx.Response().Header.Set("Content-Type", "image/png")
	ctx.Send(file)

	return nil
}

func (c *CustomerController) Create(ctx *fiber.Ctx) error {
	var requestBody requests.CustomerRequest
	if err := ctx.BodyParser(&requestBody); err != nil {
		return pkg_data.InvalidReqPayloadError{Message: "invalid request payload"}
	}

	if err := requestBody.Validate(); err != nil {
		return pkg_data.InvalidReqPayloadError{Message: err.Error()}
	}

	res, err := c.Service.Create(&requestBody)

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
