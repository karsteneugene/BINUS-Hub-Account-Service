package apiv1

import (
	"fmt"
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

	lecturer.PasswordHash, _ = HashPassword(lecturer.PasswordHash)

	db.Create(&lecturer)
	return c.JSON(lecturer)
}

// func UpdateLecturer(c *fiber.Ctx) error {
// 	lecturer := new(model.Lecturer)

// 	if err := c.BodyParser(lecturer); err != nil {
// 		return err
// 	}

// 	lecturer.PasswordHash, _ = HashPassword(lecturer.PasswordHash)

// 	db.Model(&model.Lecturer{}).Where("lecturer_id = ?", lecturer.Lecturer_ID).Updates(model.Lecturer{Fname: lecturer.Fname, Lname: lecturer.Lname, Email: lecturer.Email, PasswordHash: lecturer.PasswordHash, Phone_No: lecturer.Phone_No, Description: lecturer.Description, Profile_Img: lecturer.Profile_Img})
// 	return c.JSON(lecturer)
// }

func UpdateLecturer(c *fiber.Ctx) error {
	id := c.Params("id")
	var lecturer model.Lecturer
	db.First(&lecturer, id)
	if err := c.BodyParser(&lecturer); err != nil {
		return c.SendStatus(503)
	}
	fmt.Println(id)
	db.Where("lecturer_id = ?", id).Updates(&lecturer)
	return c.JSON(lecturer)
}

func DeleteLecturer(c *fiber.Ctx) error {
	id := c.Params("id")
	var lecturer model.Lecturer
	lecturer.Lecturer_ID = id
	db.First(&lecturer)
	db.Delete(&lecturer)
	return c.JSON(lecturer)
}
