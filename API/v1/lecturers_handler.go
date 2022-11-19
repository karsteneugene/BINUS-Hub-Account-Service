package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var lecturers = []model.Lecturer{{Lecturer_ID: "D6405", Fname: "Ardimas", Lname: "Purwita", Email: "ardimas.purwita@binus.ac.id", PasswordHash: "", Phone_No: "081112345678", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus a ligula.", Profile_Img: "https://www.youtube.com/shorts/AWOyEIuVzzQ"}}

func GetLecturers(c *fiber.Ctx) error {
	return c.JSON(lecturers)
}

func SetLecturer(c *fiber.Ctx) error {
	lecturer := new(model.Lecturer)

	if err := c.BodyParser(lecturer); err != nil {
		return err
	}

	return c.JSON(lecturer)
}
