package controllers

import (
	"encoding/json"
	"net/http"
	"placement/models"

	"placement/services"

	"github.com/gofiber/fiber/v2"
)

var (
	jwtService     services.JwtService
	studentService services.StudentService
)

func init() {
	jwtService = services.NewJwtService()
	studentService = services.NewStudentService()
}

// LoginStudent controller to login student
func LoginStudent(c *fiber.Ctx) error {
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
	student, err := studentService.LoginStudent(body.Email, body.Password)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		}
		return c.Status(error.Status).JSON(error)
	}
	accessToken, _ := jwtService.GenerateAccessToken(student.ID.Hex(), models.StudentRoll)
	refreshToken, _ := jwtService.GenerateRefreshToken(student.ID.Hex(), models.StudentRoll)
	tokenResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(tokenResponse)
}

// GetLoggedInStudent get the logged in user
func GetLoggedInStudent(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	student, err := studentService.FindStudentByID(userID)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: "someting went wrong",
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(student)
}

// RefreshToken refreshes token
func RefreshToken(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	roll := c.Locals("roll")
	accessToken, _ := jwtService.GenerateAccessToken(userID.(string), models.Roll(roll.(float64)))
	refreshToken, _ := jwtService.GenerateRefreshToken(userID.(string), models.Roll(roll.(float64)))
	tokenResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(tokenResponse)
}

// RegisterStudent to register a new student
func RegisterStudent(c *fiber.Ctx) error {
	var body *models.Student
	json.Unmarshal(c.Body(), &body)
	if body == nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "body cant be empty",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.FirstName == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "first name cant be empty",
			Key:     "fistName",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.LastName == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "last name cant be empty",
			Key:     "lastName",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.UINNumber == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "UIN number cant be empty",
			Key:     "UINNumber",
		}

		return c.Status(error.Status).JSON(error)
	}
	if len(body.PhoneNumber) != 10 {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid Phone number",
			Key:     "phoneNumber",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Gender == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "gender cant be empty",
			Key:     "gender",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Email == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "email cant be empty",
			Key:     "email",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Department == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Department cant be empty",
			Key:     "department",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Program == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "program cant be empty",
			Key:     "program",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.HomeAddress == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "home address cant be empty",
			Key:     "homeAddress",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.CurrentAddress == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "current address cant be empty",
			Key:     "currentAddress",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Password == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "password cant be empty",
			Key:     "password",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.ConfirmPassword == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "confirm password cant be empty",
			Key:     "confirmPassword",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Password != body.ConfirmPassword {
		error := models.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "password and confirm password must be the same",
			Key:     "[password confirmPassword]",
		}
		return c.Status(error.Status).JSON(error)
	}
	student, err := studentService.Register(body)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(error.Status).JSON(error)
	}
	accessToken, _ := jwtService.GenerateAccessToken(student.ID.Hex(), models.StudentRoll)
	refreshToken, _ := jwtService.GenerateRefreshToken(student.ID.Hex(), models.StudentRoll)
	tokenResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(tokenResponse)
}

func GetAllStudents(c *fiber.Ctx) error {
	students := studentService.GetAllStudents()
	return c.JSON(students)
}
