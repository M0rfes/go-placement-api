package main

import (
	"log"
	"os"
	"placement/models"
	"placement/routes"
	"placement/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	adminHashService services.HashService
)

func init() {
	adminHashService = services.NewHashService(15)

	if _, err := os.Stat("./public"); os.IsNotExist(err) {
		os.Mkdir("public", os.FileMode(0755))
		if _, err := os.Stat("./public/avatar"); os.IsNotExist(err) {
			os.Mkdir("public/avatar", os.FileMode(0755))
		}
		if _, err := os.Stat("./public.resume"); os.IsNotExist(err) {
			os.Mkdir("public/resume", os.FileMode(0755))
		}
	}
	err := mgm.SetDefaultConfig(nil, "placement", options.Client().ApplyURI("mongodb+srv://faheem:faheem@cluster0.6ezyv.mongodb.net"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	admin := &models.Admin{}
	_ = mgm.Coll(admin).First(bson.M{"username": "admin"}, admin)
	if admin.ID == primitive.NilObjectID {
		admin.Username = "admin"
		admin.Password, _ = adminHashService.HashPassword("admin")
		_ = mgm.Coll(admin).Create(admin)
	}
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Static("/", "./public")
	app.Use(cors.New())
	studentsRouter := app.Group("/students")
	companiesRouter := app.Group("/companies")
	jobsRouter := app.Group("/jobs")
	applicationRouter := app.Group("/applications")
	adminRouter := app.Group("/admin")
	routes.SetupStudentsRoute(studentsRouter)
	routes.SetupCompaniesRoute(companiesRouter)
	routes.SetupJobsRoute(jobsRouter)
	routes.SetupApplicationRoutes(applicationRouter)
	routes.SetupAdminRoute(adminRouter)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(":8080")
}
