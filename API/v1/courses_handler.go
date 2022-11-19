package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var courses = []model.Course{{ID: 1, Course_ID: "COMP6705001", Course_Name: "Distributed System", Class_ID: "L5AC", Lecturer_ID: "D6405"}}

func GetCourses(c *fiber.Ctx) error {
	return c.JSON(courses)
}

func SetCourses(c *fiber.Ctx) error {
	course := new(model.Course)

	if err := c.BodyParser(course); err != nil {
		return err
	}

	return c.JSON(course)
}
