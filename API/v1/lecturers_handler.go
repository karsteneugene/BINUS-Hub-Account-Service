package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var lecturers = []model.Lecturer{{Lecturer_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", PasswordHash: "", Phone_No: "081219201007", Description: "Lorem ipsum", Profile_Img: "<path of image>"}}

func GetLecturers(c *fiber.Ctx) error {
	return c.JSON(lecturers)
}

func SetLecturer(c *fiber.Ctx) error {
	NewLecturer := new(model.Lecturer)

	c.BodyParser(NewLecturer)

	NewLecturer.Lecturer_ID = "2440035463"
	NewLecturer.Fname = "Karsten"
	NewLecturer.Lname = "Lie"
	NewLecturer.Email = "karsten.lie@binus.ac.id"
	NewLecturer.PasswordHash = ""
	NewLecturer.Phone_No = "082139311197"
	NewLecturer.Description = "Lorem ipsum"
	NewLecturer.Profile_Img = "https://www.youtube.com/shorts/AWOyEIuVzzQ"

	// accounts = make([]model.Account, len(accounts)+1)
	lecturers = append(lecturers, *NewLecturer)
	return c.JSON(NewLecturer)
}
