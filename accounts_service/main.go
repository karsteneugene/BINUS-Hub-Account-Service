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

// getClasses responds with the list of all classes in JSON format
func getClasses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, classes)
}

// Setting up an association between the function and endpoints
func main() {
	router := gin.Default()
	router.GET("/accounts", getAccounts)
	router.GET("/courses", getCourses)
	router.GET("/classes", getClasses)
	router.Run("localhost:8080")

}
