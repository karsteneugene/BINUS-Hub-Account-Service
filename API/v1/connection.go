package apiv1

import (
	"fmt"
	"log"
	"os"
	"strconv"

	model "github.com/karsteneugene/BINUS-Hub-Account-Service/model"
	"github.com/karsteneugene/BINUS-Hub-Account-Service/setting"
	"gorm.io/gorm"
)

var mysqlparam setting.MysqlConn
var db *gorm.DB

func init() {

	getEnvVar := func(input string, output *int) {
		tmpval, err := strconv.Atoi(os.Getenv(input))
		*output = tmpval
		if err != nil {
			panic(fmt.Sprintf("Error in conversing %s with err of %s", input, err))
		}
	}

	// setting environment variables for both mysql database and swagger schema
	if len(os.Args) == 1 { // no additional arg
		os.Setenv("MYSQL_USERNAME", "root")
		os.Setenv("MYSQL_PASSWORD", "accounts_svc")
		os.Setenv("MYSQL_ADDRESS", "localhost")
		os.Setenv("MYSQL_DB_NAME", "accounts_svc")
		os.Setenv("MYSQL_MaxIdleConns", "10")
		os.Setenv("MYSQL_MaxOpenConns", "100")
		os.Setenv("MYSQL_ConnMaxIdleTime", "5")
		os.Setenv("MYSQL_ConnMaxLifetime", "60")

		os.Setenv("SCHEMA", "http")
		os.Setenv("SERVER_NAME", "localhost")
		os.Setenv("PORT", "8080")
		os.Setenv("BASE_PATH", "api/v1")
	}

	// initialized mysql parameter
	mysqlparam.MYSQL_USERNAME = os.Getenv("MYSQL_USERNAME")
	mysqlparam.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	mysqlparam.MYSQL_ADDRESS = os.Getenv("MYSQL_ADDRESS")
	mysqlparam.MYSQL_DB_NAME = os.Getenv("MYSQL_DB_NAME")
	getEnvVar("MYSQL_MaxIdleConns", &mysqlparam.MYSQL_MaxIdleConns)
	getEnvVar("MYSQL_MaxOpenConns", &mysqlparam.MYSQL_MaxOpenConns)
	getEnvVar("MYSQL_ConnMaxIdleTime", &mysqlparam.MYSQL_ConnMaxIdleTime)
	getEnvVar("MYSQL_ConnMaxLifetime", &mysqlparam.MYSQL_ConnMaxLifetime)

	var err error
	db, err = model.Connect(&mysqlparam)
	if err != nil {
		log.Panicln("Unable to connect to database: ", err.Error())
	}

}
