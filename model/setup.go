package model

import (
	// Import mysql

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect : Create a database sql connection
func Connect() (*gorm.DB, error) {

	// Karsten root:DSAccounts#2024
	// db, err := sql.Open("mysql", "root:DSAccounts#2024@tcp(127.0.0.1:3306)/accounts_svc")

	// Darren root:Scorch120403
	// db, err := sql.Open("mysql", "root:Scorch120403@tcp(127.0.0.1:3306)/accounts_svc")

	// Darren MAC root:Scorch1204
	// db, err := sql.Open("mysql", "root:Scorch1204@tcp(127.0.0.1:3306)/accounts_svc")

	// ellyz dbadmin:password
	db, err := sql.Open("mysql", "dbadmin:password@tcp(127.0.0.1:3306)/accounts_svc")

	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dbConn.AutoMigrate(&Class{}, &Lecturer{}, &Student{})
	dbConn.AutoMigrate(&Course{})
	return dbConn, nil

}
