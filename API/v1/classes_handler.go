package apiv1

import (
	model "BINUS-HUB-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var classes = []model.Class{{ID: 0, Class_Name: "L4AC"}}

func GetClasses(c *fiber.Ctx) error {
	return c.JSON(classes)
}

func SetClasses(c *fiber.Ctx) error {
	NewClass := new(model.Class)

	c.BodyParser(NewClass)

	NewClass.ID = 1
	NewClass.Class_Name = "L4BC"

	classes = append(classes, *NewClass)
	return c.JSON(NewClass)
}