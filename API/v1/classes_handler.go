package apiv1

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"
)

func GetClasses(c *fiber.Ctx) error {
	var class []model.Class
	db.Find(&class)
	return c.JSON(class)
}

func SetClass(c *fiber.Ctx) error {
	class := new(model.Class)

	if errClass := c.BodyParser(class); errClass != nil {
		return errClass
	}
	db.Create(&class)
	return c.JSON(class)
}

func UpdateClass(c *fiber.Ctx) error {
	class := new(model.Class)

	if err := c.BodyParser(class); err != nil {
		return err
	}

	db.Model(&model.Class{}).Where("id = ?", class.ID).Update("class_desc", class.Class_Desc)
	return c.JSON(class)
}


// func UpdateClass(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	class := new(model.Class)
// 	db.First(&class, id)
// 	if err := c.BodyParser(&class); err != nil {
// 		return c.SendStatus(503)

// 	}

// 	if class.ID != "" {
// 		db.Update("class_desc", class.Class_Desc)

// 	}
// 	return c.JSON(class)
// }
