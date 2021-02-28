package routes

import (
	"github.com/gofiber/fiber/v2"

	"placement/controllers"
	"placement/middelware"
)

// SetupCompaniesRoute sets routes up related to students
func SetupCompaniesRoute(router fiber.Router) {
	router.Post("/login", controllers.LoginCompany)
	router.Post("/register", controllers.RegisterCompany)
	router.Get("/refresh", middelware.IsRefreshTokenValid, controllers.RefreshToken)
	router.Get("/me/jobs", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.GetMyJobs)
	router.Get("/me/applications", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.GetAllApplicationsForCompany)
	router.Get("/me", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.GetLoggedInCompany)
	router.Get("/", controllers.GetAllCompanies)
	router.Get("/:id", controllers.GetOneCompany)
	router.Put("/", middelware.IsAccessTokenValid, middelware.IsCompany, controllers.UpdateCompany)
	router.Post("/avatar", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.UploadCompanyAvatar)
}
