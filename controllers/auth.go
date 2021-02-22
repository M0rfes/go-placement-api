package controllers

import (
	"placement/models"

	"github.com/gofiber/fiber/v2"
)

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
