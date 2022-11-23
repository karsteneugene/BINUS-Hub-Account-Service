package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

func GetCourses(c *fiber.Ctx) error {
	var courses []model.Course
	db.Find(&courses)
	return c.JSON(courses)
}

func SetCourses(c *fiber.Ctx) error {
	course := new(model.Course)

	if err := c.BodyParser(course); err != nil {
		return err
	}
	db.Create(&course)

	return c.JSON(course)
}
