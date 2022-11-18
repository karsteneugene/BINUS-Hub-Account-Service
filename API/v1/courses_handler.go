package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var courses = []model.Course{{Course_ID: "0", Course_Name: "Distributed System", Class_ID: "L1AC", Lecturer_ID: "D6059"}}

func GetCourses(c *fiber.Ctx) error {
	return c.JSON(courses)
}

func SetCourses(c *fiber.Ctx) error {
	NewCourse := new(model.Course)

	c.BodyParser(NewCourse)

	NewCourse.Course_ID = "1"
	NewCourse.Course_Name = "Intelligent Systems"
	NewCourse.Class_ID = "Intelligent Systems"
	NewCourse.Lecturer_ID = "D6058"

	courses = append(courses, *NewCourse)
	return c.JSON(NewCourse)
}
