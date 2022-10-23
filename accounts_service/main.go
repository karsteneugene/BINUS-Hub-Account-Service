package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct holding the account information
type account struct {
	ID          string `json:"id"`
	Binusian_ID string `json:"binusian_id"`
	Fname       string `json:"first_name"`
	Lname       string `json:"last_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

// Struct holding the course information
type course struct {
	ID string `json:"id"`
	Course_Name string `json:"course_name"`
}

// Struct holding the class information
type class struct {
	ID string `json:"id"`
	Class_Name string `json:"class_name"`
}

// Array holding account struct records
var accounts = []account{{ID: "0", Binusian_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", Role: "student"}}
var courses = []course{{ID: "0", Course_Name: "Distributed System"}}
var classes = []class{{ID: "0", Class_Name: "L4AC"}}

// getAccounts responds with the list of all accounts in JSON format
func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

// getCourses responds with the list of all courses in JSON format
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

// createCourse creates a new course
func createCourse(c *gin.Context) {
	var newCourse course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

// getClasses responds with the list of all classes in JSON format
func getClasses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, classes)
}

// createClass creates a new class
func createClass(c *gin.Context) {
	var newClass class

	if err := c.BindJSON(&newClass); err != nil {
		return
	}

	classes = append(classes, newClass)
	c.IndentedJSON(http.StatusCreated, newClass)
}

// Setting up an association between the function and endpoints
func main() {
	router := gin.Default()
	// REST API Account
	router.GET("/accounts", getAccounts)

	// REST API Course
	router.GET("/courses", getCourses)
	router.POST("/courses", createCourse)

	// REST API Class
	router.GET("/classes", getClasses)
	router.POST("/classes", createClass)

	router.Run("localhost:8080")

}
