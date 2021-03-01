package routes

import (
	"placement/controllers"
	"placement/middelware"

	"github.com/gofiber/fiber/v2"
)

func SetupApplicationRoutes(router fiber.Router) {
	router.Post("/", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.CreateApplication)
	router.Put("/:id", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.UpdateApplication)
	router.Delete("/:id", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.DeleteApplication)
}
