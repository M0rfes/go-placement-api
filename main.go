package main

import (
	"log"
	"placement/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "placement", options.Client().ApplyURI("mongodb+srv://faheem:faheem@cluster0.6ezyv.mongodb.net"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()
	studentsRouter := app.Group("/students")
	routes.SetupStudentsRoute(studentsRouter)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(":8080")
}
