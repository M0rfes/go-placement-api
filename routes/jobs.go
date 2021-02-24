package routes

import (
	"placement/controllers"
	"placement/middelware"

	"github.com/gofiber/fiber/v2"
)

func SetupJobsRoute(router fiber.Router) {
	router.Post("/", middelware.IsAccessTokenValid, middelware.ISCompany, controllers.AddJob)
}
