package routes

import (
	"placement/controllers"
	"placement/middelware"

	"github.com/gofiber/fiber/v2"
)

func SetupApplicationRoutes(router fiber.Router) {
	router.Post("/", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.CreateApplication)
}
