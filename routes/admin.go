package routes

import (
	"placement/controllers"
	"placement/middelware"

	"github.com/gofiber/fiber/v2"
)

// SetupAdminRoute sets up the admin routes.
func SetupAdminRoute(router fiber.Router) {
	router.Post("/login", controllers.LoginAdmin)
	router.Post("/toggle", middelware.IsAccessTokenValid, controllers.ToggleAproven)
}
