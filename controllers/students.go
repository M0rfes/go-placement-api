package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"placement/models"
	"strconv"
	"strings"

	"placement/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
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
		Roll:         models.StudentRoll,
	}
	return c.JSON(tokenResponse)
}

// GetLoggedInStudent get the logged in user
func GetLoggedInStudent(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	student, err := studentService.FindStudentByID(userID, &options.FindOneOptions{
		Projection: bson.M{"password": false},
	})
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
			Key:     "uinNumber",
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
	student, err := studentService.RegisterStudent(body)
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
		Roll:         models.StudentRoll,
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

// UpdateStudent updates the logged in student
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

// UploadStudentAvatar uploads and updates avatar for logged in student
func UploadStudentAvatar(c *fiber.Ctx) error {
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
			Status:  http.StatusBadRequest,
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
			Status:  http.StatusBadRequest,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	student, err := studentService.FindStudentByID(userID.(string))
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	student.Avatar = fmt.Sprintf("/avatar/%s.%s", userID, ext)
	err = studentService.UpdateLoggedInStudent(student)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusBadRequest,
			Key:     "avatar",
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(student)
}

// UploadStudentResume uploads and updates resume for logged in Student
func UploadStudentResume(c *fiber.Ctx) error {
	file, err := c.FormFile("resume")
	if err != nil {
		error := models.ErrorResponse{
			Message: "can't retrive file form body, Try again",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	if file == nil {
		error := models.ErrorResponse{
			Message: "empty file",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	path, err := os.Getwd()
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	userID := c.Locals("userID")
	nameArray := strings.Split(file.Filename, ".")
	ext := nameArray[len(nameArray)-1]
	if ext != "pdf" {
		error := models.ErrorResponse{
			Message: "only PDF formate are supported",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	filePath := fmt.Sprintf("%s/public/resume/%s.%s", path, userID, ext)
	err = c.SaveFile(file, filePath)
	if err != nil {
		error := models.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	student, err := studentService.FindStudentByID(userID.(string))
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	student.Resume = fmt.Sprintf("/resume/%s.%s", userID, ext)
	err = studentService.UpdateLoggedInStudent(student)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusBadRequest,
			Key:     "resume",
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(student)
}
