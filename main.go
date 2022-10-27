package main

import (
	"fmt"
	"log"

	api "github.com/karsteneugene/BINUS-Hub-Account-Service/API"

	"github.com/karsteneugene/BINUS-Hub-Account-Service/setting"
)

var swaggerparam setting.Swagger

func init() {
	// prepare swagger (testing using hardcoded values)
	swaggerparam.SCHEMA = "http"
	swaggerparam.PORT = "8080"
	swaggerparam.SERVER_NAME = "localhost"
	swaggerparam.BASE_PATH = "api/v1"
}

// Setting up an association between the function and endpoints
func main() {
	app := api.Api()
	app.Static("/", "./swagger-ui-4.15.1/dist")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", swaggerparam.PORT)))

}
