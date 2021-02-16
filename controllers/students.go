package controllers

import (
	"encoding/json"
	"net/http"
	"placement/models"
	"strconv"

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
			Message: "body can't be empty",
			Key:     "email,password",
		}
		return c.Status(400).JSON(error)
	}
	if body.Email == "" {
		error := models.ErrorResponse{
			Status:  400,
			Message: "email can't be empty",
			Key:     "email",
		}
		return c.Status(400).JSON(error)
	}
	if body.Password == "" {
		error := models.ErrorResponse{
			Status:  400,
			Message: "password can't be empty",
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
			Message: "something went wrong",
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
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: "something went wrong",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body == nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "body can't be empty",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.FirstName == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "First name can't be empty",
			Key:     "fistName",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.LastName == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Last name can't be empty",
			Key:     "lastName",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.UINNumber == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "UIN number can't be empty",
			Key:     "uinNumber",
		}

		return c.Status(error.Status).JSON(error)
	}
	if len(body.PhoneNumber) != 10 {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Phone number",
			Key:     "phoneNumber",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Gender == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Gender can't be empty",
			Key:     "gender",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Email == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Email can't be empty",
			Key:     "email",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Department == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Department can't be empty",
			Key:     "department",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Program == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Program can't be empty",
			Key:     "program",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.HomeAddress == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Home address can't be empty",
			Key:     "homeAddress",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.CurrentAddress == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Current address can't be empty",
			Key:     "currentAddress",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Password == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Password can't be empty",
			Key:     "password",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.ConfirmPassword == "" {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Confirm password can't be empty",
			Key:     "confirmPassword",
		}
		return c.Status(error.Status).JSON(error)
	}
	if body.Password != body.ConfirmPassword {
		error := models.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "Password and confirm password must be the same",
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

// GetAllStudents gets a list of all students
func GetAllStudents(c *fiber.Ctx) error {
	limit, err := strconv.ParseInt(c.Query("limit", "30"), 10, 64)
	if err != nil {
		limit = 30
	}
	skip, err := strconv.ParseInt(c.Query("skip", "0"), 10, 64)
	if err != nil {
		skip = 0
	}
	students := studentService.GetAllStudents(&limit, &skip)
	return c.JSON(students)
}

func UpdateStudent(c *fiber.Ctx) error {
	var body *models.Student
	id := c.Params("id")
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: "something went wrong",
		}
		return c.Status(error.Status).JSON(error)
	}
	student, err := studentService.FindStudentByID(id)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: err.Error(),
		}
		return c.Status(error.Status).JSON(error)
	}
	if firstName := body.FirstName; firstName != "" {
		student.FirstName = firstName
	}
	if lastName := body.LastName; lastName != "" {
		student.LastName = lastName
	}
	if uinNumber := body.UINNumber; uinNumber != "" {
		student.UINNumber = uinNumber
	}
	if phoneNumber := body.PhoneNumber; phoneNumber != "" {
		student.PhoneNumber = phoneNumber
	}
	if gender := body.Gender; gender != "" {
		student.Gender = gender
	}
	if email := body.Email; email != "" {
		student.Email = email
	}
	if department := body.Department; department != "" {
		student.Department = department
	}
	if program := body.Program; program != "" {
		student.Program = program
	}
	if currentAddress := body.CurrentAddress; currentAddress != "" {
		student.CurrentAddress = currentAddress
	}
	if homeAddress := body.HomeAddress; homeAddress != "" {
		student.HomeAddress = homeAddress
	}
	err = studentService.UpdateLoggedInStudent(student)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: err.Error(),
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(student)
}
