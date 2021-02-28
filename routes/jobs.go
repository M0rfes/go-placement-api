package routes

import (
	"placement/controllers"
	"placement/middelware"

	"github.com/gofiber/fiber/v2"
)

// SetupJobsRoute to setup jobs routes.
func SetupJobsRoute(router fiber.Router) {
	router.Post("/", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.AddJob)
	router.Get("/", controllers.GetAllJobs)
	router.Get("/:id/applications", controllers.GetAllApplicationsForJob)
	router.Get("/:id", controllers.GetJobByID)
	router.Put("/:id", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.UpdateJob)
}
