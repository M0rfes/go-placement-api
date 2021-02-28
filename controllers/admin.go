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
	body, err = adminService.LoginAdmin(body.Password)
	if err != nil {
		error := models.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
		return c.Status(error.Status).JSON(error)
	}
	accessToken, _ := jwtService.GenerateAccessToken(body.ID.Hex(), models.AdminRoll)
	refreshToken, _ := jwtService.GenerateRefreshToken(body.ID.Hex(), models.AdminRoll)
	response := &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(response)
}

func ToggleAproven(c *fiber.Ctx) error {
	var body *models.ToggleAproven
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "body cant be empty",
		}
		return c.Status(error.Status).JSON(error)
	}
	err = adminService.ToggleAproven(body.StudentID)
	if err != nil {
		error := models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "someting went wrong",
		}
		return c.Status(error.Status).JSON(error)
	}
	return c.JSON(map[string]uint8{
		"status": http.StatusOK,
	})
}
