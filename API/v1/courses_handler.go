package apiv1

import (
	"fmt"
	"strconv"

	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

func GetCourses(c *fiber.Ctx) error {
	var courses []model.Course
	db.Find(&courses)
	return c.JSON(courses)
}

func SetCourse(c *fiber.Ctx) error {
	course := new(model.Course)

	if err := c.BodyParser(course); err != nil {
		return err
	}
	db.Create(&course)

	return c.JSON(course)
}

func UpdateCourse(c *fiber.Ctx) error {
	course := new(model.Course)

	if err := c.BodyParser(course); err != nil {
		return err
	}

	db.Model(&model.Course{}).Where("id = ?", course.ID).Updates(model.Course{Course_ID: course.Course_ID, Course_Name: course.Course_Name, Class_ID: course.Class_ID, Lecturer_ID: course.Lecturer_ID})
	return c.JSON(course)
}

func DeleteCourse(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	var course model.Course
	course.ID = uint(id)
	db.First(&course)
	db.Delete(&course)
	return c.JSON(course)
}
