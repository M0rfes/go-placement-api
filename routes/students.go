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
	router.Get("/me/applications", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.GetAllApplicationsForStudent)
	router.Delete("/me", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.DeleteStudent)
	router.Get("/me", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.GetLoggedInStudent)
	router.Get("/", controllers.GetAllStudents)
	router.Get("/approved", controllers.GetAllApprovedStudent)
	router.Get("/unapproved", controllers.GetAllUnApprovedStudent)
	router.Get("/:id", controllers.GetOneStudent)
	router.Delete("/:id", middelware.IsAccessTokenValid, middelware.IsAdmin, controllers.DeleteStudentById)
	router.Put("/", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.UpdateStudent)
	router.Post("/avatar", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.UploadStudentAvatar)
	router.Post("/resume", middelware.IsAccessTokenValid, middelware.IsStudent, controllers.UploadStudentResume)
}
