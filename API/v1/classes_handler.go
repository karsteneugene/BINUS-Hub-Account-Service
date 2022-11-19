package apiv1

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"
)

var classes = []model.Class{{ID: "L5AC", Class_Desc: "Computer Science 5th Semester Class A"}}

func GetClasses(c *fiber.Ctx) error {
	return c.JSON(classes)
}

func SetClasses(c *fiber.Ctx) error {
	class := new(model.Class)

	if err := c.BodyParser(class); err != nil {
		return err
	}

	return c.JSON(class)
}
