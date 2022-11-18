package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

var accounts = []model.Student{{Binusian_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", PasswordHash: "", Phone_No: "081219201007", Description: "Lorem ipsum", Profile_Img: "<path of image>"}}

func GetAccounts(c *fiber.Ctx) error {
	return c.JSON(accounts)
}

func SetAccount(c *fiber.Ctx) error {
	NewAccount := new(model.Student)

	c.BodyParser(NewAccount)

	NewAccount.Binusian_ID = "2440035463"
	NewAccount.Fname = "Karsten"
	NewAccount.Lname = "Lie"
	NewAccount.Email = "karsten.lie@binus.ac.id"
	NewAccount.PasswordHash = ""
	NewAccount.Phone_No = "082139311197"
	NewAccount.Description = "Lorem ipsum"
	NewAccount.Profile_Img = "https://www.youtube.com/shorts/AWOyEIuVzzQ"

	// accounts = make([]model.Account, len(accounts)+1)
	accounts = append(accounts, *NewAccount)
	return c.JSON(NewAccount)
}

// // Array holding account struct records

// var accounts = []model.Account{{ID: 0, Binusian_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", PasswordHash: "", Phone_No: "081219201007", Role: "student", Description: "Lorem ipsum", Profile_Img: "<path of image>"}}
// var courses = []model.Course{{ID: 0, Course_Name: "Distributed System"}}
// var classes = []model.Class{{ID: 0, Class_Name: "L4AC"}}

// // getAccounts responds with the list of all accounts in JSON format
// func getAccounts(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, accounts)
// }

// // getCourses responds with the list of all courses in JSON format
// func getCourses(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, courses)
// }

// // createCourse creates a new course
// func createCourse(c *gin.Context) {
// 	var newCourse model.Course

// 	if err := c.BindJSON(&newCourse); err != nil {
// 		return
// 	}

// 	courses = append(courses, newCourse)
// 	c.IndentedJSON(http.StatusCreated, newCourse)
// }

// // getClasses responds with the list of all classes in JSON format
// func getClasses(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, classes)
// }

// // createClass creates a new class
// func createClass(c *gin.Context) {
// 	var newClass model.Class

// 	if err := c.BindJSON(&newClass); err != nil {
// 		return
// 	}

// 	classes = append(classes, newClass)
// 	c.IndentedJSON(http.StatusCreated, newClass)
// }
