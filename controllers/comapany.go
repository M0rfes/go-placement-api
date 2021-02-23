package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"placement/models"
	"placement/services"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	companyService services.CompanyService
)

func init() {
	companyService = services.NewCompanyService()
}

// RegisterCompany handler to register a company.
func RegisterCompany(c *fiber.Ctx) error {
	var body = &models.Company{
		Name:               c.FormValue("name"),
		Email:              c.FormValue("email"),
		RegistrationNumber: c.FormValue("registrationNumber"),
		GSTNumber:          c.FormValue("gstNumber"),
		WebSiteURL:         c.FormValue("webSiteURL"),
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
			Key:     "[password,confirmPassword]",
		}
		return c.Status(error.Status).JSON(error)
	}
	body, err := companyService.RegisterCompany(body)
	if err != nil {
		error := models.ErrorResponse{
			Message: err.Error(),
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
		companyService.UpdateCompany(body)
	}
	accessToken, _ := jwtService.GenerateAccessToken(body.ID.Hex(), models.CompanyRoll)
	refreshToken, _ := jwtService.GenerateRefreshToken(body.ID.Hex(), models.CompanyRoll)
	tokenResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Roll:         models.StudentRoll,
	}
	return c.JSON(tokenResponse)
}

// LoginCompany handler to login a company.
func LoginCompany(c *fiber.Ctx) error {
	var body *models.EmailAndPassword
	json.Unmarshal(c.Body(), &body)
	if body == nil {
		error := models.ErrorResponse{
			Status:  400,
			Message: "body cant be empty",
			Key:     "email,password",
		}
		return c.Status(400).JSON(error)
	}
	if body.Email == "" {
		error := models.ErrorResponse{
			Status:  400,
			Message: "email cant be empty",
			Key:     "email",
		}
		return c.Status(400).JSON(error)
	}
	if body.Password == "" {
		error := models.ErrorResponse{
			Status:  400,
			Message: "password cant be empty",
			Key:     "password",
		}
		return c.Status(400).JSON(error)
	}
	company, err := companyService.LoginCompany(body.Email, body.Password)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		}
		return c.Status(error.Status).JSON(error)
	}
	accessToken, _ := jwtService.GenerateAccessToken(company.ID.Hex(), models.CompanyRoll)
	refreshToken, _ := jwtService.GenerateRefreshToken(company.ID.Hex(), models.CompanyRoll)
	tokenResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Roll:         models.StudentRoll,
	}
	return c.JSON(tokenResponse)
}

// GetLoggedInCompany handler to get the logged in company.
func GetLoggedInCompany(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	company, err := companyService.FindCompanyByID(userID.(string))
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(company)
}

// GetAllCompanies handler to get all companies.
func GetAllCompanies(c *fiber.Ctx) error {
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 30
	}
	skip, err := strconv.ParseInt(c.Query("skip"), 10, 64)
	if err != nil {
		skip = 0
	}
	companies := companyService.GetAllCompanies(&limit, &skip)

	return c.JSON(companies)
}

// GetOneCompany handler to get a company by its ID.
func GetOneCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := companyService.FindCompanyByID(id)
	if err != nil {
		error := &models.ErrorResponse{
			Message: fmt.Sprintf("company with id: %s not found", id),
			Status:  http.StatusNotFound,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(company)
}

// UpdateCompany handler to updates a company.
func UpdateCompany(c *fiber.Ctx) error {
	var body *models.Company
	id := c.Locals("userID").(string)
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong",
		}
		return c.Status(error.Status).JSON(error)
	}
	company, err := companyService.FindCompanyByID(id)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("cant find student with id %s", id),
		}
		return c.Status(error.Status).JSON(error)
	}
	if name := body.Name; name != "" {
		company.Name = name
	}
	if email := body.Email; email != "" {
		company.Email = email
	}
	if registrationNumber := body.RegistrationNumber; registrationNumber != "" {
		company.RegistrationNumber = registrationNumber
	}
	if gstNumber := body.GSTNumber; gstNumber != "" {
		company.GSTNumber = gstNumber
	}
	if webSiteURL := body.WebSiteURL; webSiteURL != "" {
		company.WebSiteURL = webSiteURL
	}
	if phoneNumber := body.PhoneNumber; phoneNumber != "" {
		company.PhoneNumber = phoneNumber
	}
	if address := body.Address; address != "" {
		company.Address = address
	}
	err = companyService.UpdateCompany(company)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "someting went wrong",
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(company)
}

func UploadCompanyAvatar(c *fiber.Ctx) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		error := models.ErrorResponse{
			Message: "can't retrive file form body, Try again",
			Status:  http.StatusBadRequest,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	if file == nil {
		error := models.ErrorResponse{
			Message: "empty file",
			Status:  http.StatusBadRequest,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	path, err := os.Getwd()
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	userID := c.Locals("userID")
	nameArray := strings.Split(file.Filename, ".")
	ext := nameArray[len(nameArray)-1]
	if ext != "jpg" && ext != "png" && ext != "jpeg" {
		error := models.ErrorResponse{
			Message: "only jpg png and jpeg file formates are supported",
			Status:  http.StatusBadRequest,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	filePath := fmt.Sprintf("%s/public/avatar/%s.%s", path, userID, ext)
	err = c.SaveFile(file, filePath)
	if err != nil {
		error := models.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}

	company, err := companyService.FindCompanyByID(userID.(string))
	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("company with ID %s not fount", userID),
			Status:  http.StatusNotFound,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	company.Avatar = fmt.Sprintf("/avatar/%s.%s", userID, ext)
	err = companyService.UpdateCompany(company)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	company.Password = ""
	return c.JSON(company)
}
