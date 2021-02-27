package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"placement/models"
	"placement/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
