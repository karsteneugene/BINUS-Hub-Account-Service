package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var courses = []model.Course{{ID: 0, Course_Name: "Distributed System"}}

func GetCourses(c *fiber.Ctx) error {
	return c.JSON(courses)
}

func SetCourses(c *fiber.Ctx) error {
	NewCourse := new(model.Course)

	c.BodyParser(NewCourse)

	NewCourse.ID = 1
	NewCourse.Course_Name = "Intelligent Systems"

	courses = append(courses, *NewCourse)
	return c.JSON(NewCourse)
}
