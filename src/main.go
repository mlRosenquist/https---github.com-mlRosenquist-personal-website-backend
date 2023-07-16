package main

import (
	"personal-website-backend/configs"
	"personal-website-backend/controllers"
	"personal-website-backend/repositories"
	"personal-website-backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mlRosenquist/personal-website-backend/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.ImageRoute(e, controllers.ImageController{
		FileRepository: &repositories.MongoFileRepository{},
	})

	e.Logger.Fatal(e.Start(":8000"))
}
