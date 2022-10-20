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

// Array holding account struct records
var accounts = []account{{ID: "0", Binusian_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", Role: "student"}}

// getAccounts responds with the list of all accounts in JSON format
func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

// Setting up an association between the function and endpoints
func main() {
	router := gin.Default()
	router.GET("/accounts", getAccounts)

	router.Run("localhost:8080")

}
