package controllers

import (
	"encoding/json"
	"net/http"
	"placement/models"
	"placement/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	jobService services.JobService
)

func init() {
	jobService = services.NewJobService()
}

func AddJob(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var job *models.Job
	err := json.Unmarshal(c.Body(), &job)
	if err != nil {
		error := &models.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	if job == nil {
		error := models.ErrorResponse{
			Message: "body cant be empty",
			Status:  http.StatusBadRequest,
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.Title == "" {
		error := models.ErrorResponse{
			Message: "title cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "title",
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.Description == "" {
		error := models.ErrorResponse{
			Message: "description cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "description",
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.Openings == 0 {
		error := models.ErrorResponse{
			Message: "openings cant be 0",
			Status:  http.StatusBadRequest,
			Key:     "openings",
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.Type == "" {
		error := models.ErrorResponse{
			Message: "type cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "type",
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.Location == "" {
		error := models.ErrorResponse{
			Message: "location cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "location",
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.Position == "" {
		error := models.ErrorResponse{
			Message: "position cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "position",
		}
		return c.Status(error.Status).JSON(error)
	}
	if job.LastDayOfSummission == primitive.DateTime(0) {
		error := models.ErrorResponse{
			Message: "last day of summission cant be empty",
			Status:  http.StatusBadRequest,
			Key:     "lastDayOfSummission",
		}
		return c.Status(error.Status).JSON(error)
	}

	job.Company, _ = primitive.ObjectIDFromHex(userID.(string))
	job, err = jobService.CreateJob(job)
	if err != nil {
		error := &models.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(job)
}
