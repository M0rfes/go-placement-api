package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"placement/models"
	"placement/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	b "gopkg.in/mgo.v2/bson"
)

var (
	applicationService services.ApplicationService
)

func init() {
	applicationService = services.NewApplicationService()
}

func CreateApplication(c *fiber.Ctx) error {
	var application *models.Application
	err := json.Unmarshal(c.Body(), &application)
	if err != nil {
		error := models.ErrorResponse{
			Message: "body can't be empty",
			Status:  http.StatusBadRequest,
		}
		return c.Status(error.Status).JSON(error)
	}
	if application.JobID == primitive.NilObjectID {
		error := models.ErrorResponse{
			Message: "jobId cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "jobId",
		}
		return c.Status(error.Status).JSON(error)
	}
	application.Status = "pending"
	userID := c.Locals("userID").(string)
	application.StudentID, _ = primitive.ObjectIDFromHex(userID)
	job, err := jobService.GetJobByID(application.JobID.Hex())
	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("cant find job by ID %s", application.JobID),
			Status:  http.StatusNotFound,
			Key:     "jobId",
		}
		return c.Status(error.Status).JSON(error)
	}
	application.CompanyID = job.Company.ID
	err = applicationService.CreateApplication(application)
	if err != nil {
		error := models.ErrorResponse{
			Message: "internal server error",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(application)
}

func GetAllApplications(c *fiber.Ctx) error {
	applications := applicationService.GetAllApplications()
	return c.JSON(applications)
}

func GetApplicationById(c *fiber.Ctx) error {
	id := c.Params("id")
	application, err := applicationService.GetApplicationById(id)
	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("application with id %s not found", id),
			Status:  http.StatusNotFound,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(application)
}

func UpdateApplication(c *fiber.Ctx) error {
	var body *models.Application
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Message: "body cant be empty",
			Status:  http.StatusBadRequest,
		}
		return c.Status(error.Status).JSON(error)
	}
	userID := c.Locals("userID").(string)

	id := c.Params("id")
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("application with id %s not found", id),
			Status:  http.StatusNotFound,
		}
		return c.Status(error.Status).JSON(error)
	}
	application, err := applicationService.FindOneApplication(&b.M{"_id": pid})
	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("application with id %s not found", id),
			Status:  http.StatusNotFound,
		}
		return c.Status(error.Status).JSON(error)
	}
	if userID != application.CompanyID.Hex() {
		error := models.ErrorResponse{
			Message: "you don't have access to this entity",
			Status:  http.StatusForbidden,
		}
		return c.Status(error.Status).JSON(error)
	}
	if status := body.Status; status != "" {
		application.Status = status
	}
	err = applicationService.UpdateApplication(application)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(application)
}

// DeleteApplication
func DeleteApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	err := applicationService.DeleteApplication(id)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(map[string]uint{"status": http.StatusOK})
}
