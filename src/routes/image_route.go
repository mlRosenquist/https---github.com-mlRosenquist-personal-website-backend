package routes

import (
	"personal-website-backend/controllers"

	"github.com/labstack/echo/v4"
)

func ImageRoute(e *echo.Echo) {
	e.POST("/image", controllers.CreateImage)
	// e.GET("/image/:imageId", controllers.GetAImage)
	// e.PUT("/image/:imageId", controllers.EditAImage)
	// e.DELETE("/image/:imageId", controllers.DeleteAImage)
	// e.GET("/images", controllers.GetAllUsers)
}
