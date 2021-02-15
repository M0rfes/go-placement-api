package routes

import (
	"github.com/gofiber/fiber/v2"

	"placement/controllers"
	"placement/middelware"
)

// SetupStudentsRoute sets routes up related to students
func SetupStudentsRoute(router fiber.Router) {
	router.Post("/login", controllers.LoginStudent)
	router.Post("/register", controllers.RegisterStudent)
	router.Get("/refresh", middelware.IsRefreshTokenValid, controllers.RefreshToken)
	router.Get("/me", middelware.IsAccessTokenValid, controllers.GetLoggedInUser)
}
