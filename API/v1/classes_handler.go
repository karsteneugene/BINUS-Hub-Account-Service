package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var classes = []model.Class{{ID: "L5AC", Class_Desc: "L4AC"}}

func GetClasses(c *fiber.Ctx) error {
	return c.JSON(classes)
}

func SetClasses(c *fiber.Ctx) error {
	NewClass := new(model.Class)

	c.BodyParser(NewClass)

	NewClass.ID = "L5AC"
	NewClass.Class_Desc = "RICK ROLL"

	classes = append(classes, *NewClass)
	return c.JSON(NewClass)
}
