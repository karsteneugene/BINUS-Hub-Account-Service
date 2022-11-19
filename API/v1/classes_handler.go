package apiv1

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"
)

var db, err = model.Connect()

func GetClasses(c *fiber.Ctx) error {
	var class []model.Class
	db.Find(&class)
	return c.JSON(class)
}

func SetClasses(c *fiber.Ctx) error {
	class := new(model.Class)

	if err := c.BodyParser(class); err != nil {
		return err
	}

	db.Create(&class)

	return c.JSON(class)
}
