package controllers

import (
	"fmt"
	"net/http"
	"os"
	"placement/models"
	"placement/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	companyService services.CompanyService
)

func init() {
	companyService = services.NewCompanyService()
}

func RegisterCompany(c *fiber.Ctx) error {
	var body = &models.Company{
		Name:               c.FormValue("name"),
		Email:              c.FormValue("email"),
		RegistrationNumber: c.FormValue("registrationNumber"),
		GSTNumber:          c.FormValue("gstNumber"),
		WebSiteURL:         c.FormValue("webSiteUrl"),
		PhoneNumber:        c.FormValue("phoneNumber"),
		Address:            c.FormValue("address"),
		Password:           c.FormValue("password"),
		ConfirmPassword:    c.FormValue("confirmPassword"),
	}
	if body.Name == "" {
		error := models.ErrorResponse{
			Message: "name cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "name",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Email == "" {
		error := models.ErrorResponse{
			Message: "email cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "email",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.RegistrationNumber == "" {
		error := models.ErrorResponse{
			Message: "registerion number cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "registerionNumber",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.GSTNumber == "" {
		error := models.ErrorResponse{
			Message: "gst number cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "gstNumber",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.WebSiteURL == "" {
		error := models.ErrorResponse{
			Message: "web siteurl cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "webSiteURL",
		}
		return c.Status(error.Status).JSON(error)
	}
	if len(body.PhoneNumber) != 10 {
		error := models.ErrorResponse{
			Message: "phone number should be 10 digits long",
			Status:  http.StatusBadRequest,
			Key:     "phoneNumber",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Address == "" {
		error := models.ErrorResponse{
			Message: "addres can't be empty",
			Status:  http.StatusBadRequest,
			Key:     "address",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Password == "" {
		error := models.ErrorResponse{
			Message: "password can't be empty",
			Status:  http.StatusBadRequest,
			Key:     "password",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.ConfirmPassword == "" {
		error := models.ErrorResponse{
			Message: "confirm password can't be empty",
			Status:  http.StatusBadRequest,
			Key:     "confirmPassword",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Password != body.ConfirmPassword {
		error := models.ErrorResponse{
			Message: "password and confirm password should be same",
			Status:  http.StatusBadRequest,
			Key:     "password,confirmPassword",
		}
		return c.Status(error.Status).JSON(error)
	}
	body, err := companyService.RegisterCompany(body)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	avatar, err := c.FormFile("avatar")
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong while extracting avatar",
			Status:  http.StatusInternalServerError,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	nameArray := strings.Split(avatar.Filename, ".")
	avatarExt := nameArray[len(nameArray)-1]
	if avatarExt != "jpg" && avatarExt != "png" && avatarExt != "jpeg" {
		error := models.ErrorResponse{
			Message: "only jpg png and jpeg file formates are supported",
			Status:  http.StatusBadRequest,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	body, err = companyService.RegisterCompany(body)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	path, err := os.Getwd()
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(error.Status).JSON(error)
	}
	avatarPath := fmt.Sprintf("%s/public/avatar/%s.%s", path, body.ID.Hex(), avatarExt)

	if err = c.SaveFile(avatar, avatarPath); err == nil {
		body.Avatar = fmt.Sprintf("/avatar/%s.%s", body.ID.Hex(), avatarExt)
		// companyService
	}
	return c.JSON(body)
}
