package apiv1

import (
	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"

	"github.com/gofiber/fiber/v2"
)

func GetStudents(c *fiber.Ctx) error {
	var students []model.Student
	db.Find(&students)
	return c.JSON(students)
}

func SetStudent(c *fiber.Ctx) error {
	student := new(model.Student)

	if err := c.BodyParser(student); err != nil {
		return err
	}

	student.PasswordHash, _ = HashPassword(student.PasswordHash)

	db.Create(&student)
	return c.JSON(student)
}

func UpdateStudent(c *fiber.Ctx) error {
	student := new(model.Student)

	if err := c.BodyParser(student); err != nil {
		return err
	}

	student.PasswordHash, _ = HashPassword(student.PasswordHash)

	db.Model(&model.Student{}).Where("binusian_id = ?", student.Binusian_ID).Updates(model.Student{Fname: student.Fname, Lname: student.Lname, Email: student.Email, PasswordHash: student.PasswordHash, Phone_No: student.Phone_No, Description: student.Description, Profile_Img: student.Profile_Img})
	return c.JSON(student)
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
