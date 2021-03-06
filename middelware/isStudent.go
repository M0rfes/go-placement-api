package middelware

import (
	"placement/models"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

// IsStudent middelware checks whether a logged user is student
func IsStudent(c *fiber.Ctx) error {
	roll := c.Locals("roll")
	if models.StudentRoll != models.Roll(roll.(float64)) {
		err := models.UnAuthorizeError{
			Message: "You don't have access to this entity",
			Status:  http.StatusForbidden,
		}
		return c.Status(err.Status).JSON(err)
	}
	return c.Next()
}
