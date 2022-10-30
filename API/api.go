package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/karsteneugene/BINUS-Hub-Account-Service/model"
	apiv1 "github.com/karsteneugene/BINUS-Hub-Account-Service/API/v1"
)


func middleware(c *fiber.Ctx) error {
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendString(c.Path())
}


func Api() *fiber.App {
	
	model.Connect()

	// Webserver
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// Define Root API path
	api := app.Group("/api", middleware)

	// Group v1 API path
	v1 := api.Group("/v1", middleware)

	// Check health of v1 path
	v1.Get("/accounts", apiv1.GetAccounts)
	v1.Post("/accounts", apiv1.SetAccount)
	v1.Get("/courses", apiv1.GetCourses)
	v1.Post("/courses", apiv1.SetCourses)
	v1.Get("/classes", apiv1.GetClasses)
	v1.Post("/classes", apiv1.SetClasses)

	return app
}
