package model

import (
	"fmt"
	// Import mysql
	// "github.com/go-sql-driver/mysql"
	// "database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect : Create a database sql connection
func Connect() {

	// Darren root:Scorch120403
	db, err := gorm.Open(mysql.Open("root:DSAccounts#2024@tcp(127.0.0.1:3306)/accounts_svc"))

	// ellyz dbadmin:password
	//db, err := gorm.Open(mysql.Open("dbadmin:password@tcp(127.0.0.1:3306)/accounts_svc"))

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&Class{}, &Lecturer{}, &Student{})

	dbConn = db
}
