package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

func GetLecturers(c *fiber.Ctx) error {
	var lecturers []model.Lecturer
	db.Find(&lecturers)
	return c.JSON(lecturers)
}

func SetLecturer(c *fiber.Ctx) error {
	lecturer := new(model.Lecturer)

	if err := c.BodyParser(lecturer); err != nil {
		return err
	}

	db.Create(&lecturer)
	return c.JSON(lecturer)
}

func UpdateLecturer(c *fiber.Ctx) error {
	lecturer := new(model.Lecturer)

	if err := c.BodyParser(lecturer); err != nil {
		return err
	}

	db.Model(&model.Lecturer{}).Where("lecturer_id = ?", lecturer.Lecturer_ID).Updates(model.Lecturer{Fname: lecturer.Fname, Lname: lecturer.Lname, Email: lecturer.Email, Phone_No: lecturer.Phone_No, Description: lecturer.Description, Profile_Img: lecturer.Profile_Img})
	return c.JSON(lecturer)
}
