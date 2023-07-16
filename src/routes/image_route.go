package routes

import (
	"personal-website-backend/controllers"

	"github.com/labstack/echo/v4"
)

func ImageRoute(e *echo.Echo, controller controllers.ImageController) {

	e.POST("/image", controller.CreateImage)
	e.GET("/image/:imageId", controller.GetAImage)
	e.PUT("/image/:imageId", controller.EditAImage)
	e.DELETE("/image/:imageId", controller.DeleteAImage)
	e.GET("/images", controller.GetAllImages)
}
