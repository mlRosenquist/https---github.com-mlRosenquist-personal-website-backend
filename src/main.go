package main

import (
	"personal-website-backend/configs"
	"personal-website-backend/routes"

	"github.com/labstack/echo/v4"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}
