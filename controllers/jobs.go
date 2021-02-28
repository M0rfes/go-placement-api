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
	jobService services.JobService
)

func init() {
	jobService = services.NewJobService()
}

// AddJob handler to add a job.
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
	if job.CTC == 0 {
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

	job.CompanyID, _ = primitive.ObjectIDFromHex(userID.(string))
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

// GetAllJobs handler to get all jobs.
func GetAllJobs(c *fiber.Ctx) error {
	jobs := jobService.GetAllJobs()
	return c.JSON(jobs)
}

// GetJobByID handler to get one job by ID
func GetJobByID(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := jobService.GetJobByID(id)
	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("job with id %s not found", id),
			Status:  http.StatusNotFound,
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(job)
}

// UpdateJob handler to update a job.
func UpdateJob(c *fiber.Ctx) error {
	id := c.Params("id")
	var body *models.Job
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Message: "something went wrong",
			Status:  http.StatusInternalServerError,
		}
		return c.Status(error.Status).JSON(error)
	}
	job, err := jobService.GetJobByID(id)
	if job.CompanyID.Hex() != id {
		error := models.ErrorResponse{
			Message: "you don't own this entity",
			Status:  http.StatusForbidden,
		}
		return c.Status(error.Status).JSON(error)
	}
	if ctc := body.CTC; ctc != 0 {
		job.CTC = ctc
	}
	if description := body.Description; description != "" {
		job.Description = description
	}
	if openings := body.Openings; openings != 0 {
		job.Openings = openings
	}
	if jobType := body.Type; jobType != "" {
		job.Type = jobType
	}
	if location := body.Location; location != "" {
		job.Location = location
	}
	if position := body.Position; position != "" {
		job.Position = position
	}
	if ls := body.LastDayOfSummission; ls != primitive.DateTime(0) {
		job.LastDayOfSummission = ls
	}

	if err != nil {
		error := models.ErrorResponse{
			Message: fmt.Sprintf("job with id %s not found", id),
			Status:  http.StatusNotFound,
		}
		return c.Status(error.Status).JSON(error)
	}
	err = jobService.UpdateJob(job)
	return c.JSON(job)
}

// GetAllApplicationsForJob handler to get all applications for a job.
func GetAllApplicationsForJob(c *fiber.Ctx) error {
	id := c.Params("id")
	applications := applicationService.GetAllApplicationsForAJob(id)
	return c.JSON(applications)
}
