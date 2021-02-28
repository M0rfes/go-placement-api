package middelware

import (
	"net/http"
	"placement/models"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) error {
	roll := c.Locals("roll")
	if models.AdminRoll != models.Roll(roll.(float64)) {
		err := models.UnAuthorizeError{
			Message: "You don't have access to this entity",
			Status:  http.StatusForbidden,
		}
		return c.Status(err.Status).JSON(err)
	}
	return c.Next()
}
