package model

import (
	// Import mysql

	"database/sql"
	"fmt"

	"github.com/karsteneugene/BINUS-Hub-Account-Service/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect : Create a database sql connection
func Connect(params *setting.MysqlConn) (*gorm.DB, error) {

	// Karsten root:DSAccounts#2024
	// db, err := sql.Open("mysql", "root:DSAccounts#2024@tcp(127.0.0.1:3306)/accounts_svc")

	// Darren root:Scorch120403
	// db, err := sql.Open("mysql", "root:Scorch120403@tcp(127.0.0.1:3306)/accounts_svc")

	// Darren MAC root:Scorch1204
	// db, err := sql.Open("mysql", "root:Scorch1204@tcp(127.0.0.1:3306)/accounts_svc")

	// ellyz dbadmin:password
	// db, err := sql.Open("mysql", "dbadmin:password@tcp(127.0.0.1:3306)/accounts_svc")

	// Docker accounts_db:accounts_svc
	// db, err := sql.Open("mysql", "root:accounts_svc@tcp("+os.Getenv("DB_HOST")+":3306)/accounts_svc")
	// fmt.Println("TCP CONNECT: " + "root:accounts_svc@tcp(" + os.Getenv("DB_HOST") + ":3306)/accounts_svc")

	// params.MYSQL_USERNAME+":"+params.MYSQL_PASSWORD+"@tcp("+params.MYSQL_ADDRESS+":3306)/"+params.MYSQL_DB_NAME
	db, err := sql.Open("mysql", params.MYSQL_USERNAME+":"+params.MYSQL_PASSWORD+"@tcp("+params.MYSQL_ADDRESS+":3306)/"+params.MYSQL_DB_NAME)
	fmt.Println("TCP CONNECT: " + params.MYSQL_USERNAME + ":" + params.MYSQL_PASSWORD + "@tcp(" + params.MYSQL_ADDRESS + ":3306)/" + params.MYSQL_DB_NAME)

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
