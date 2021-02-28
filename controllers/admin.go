package controllers

import (
	"encoding/json"
	"net/http"
	"placement/models"
	"placement/services"

	"github.com/gofiber/fiber/v2"
)

var (
	adminService services.AdminService
)

func init() {
	adminService = services.NewAdminService()
}

func LoginAdmin(c *fiber.Ctx) error {
	var body *models.Admin
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "body cant be empty",
		}
		return c.Status(error.Status).JSON(error)
	}
}
