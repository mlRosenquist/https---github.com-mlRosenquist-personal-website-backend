package controllers

import (
	"net/http"
	"personal-website-backend/repositories"
	"personal-website-backend/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var imageValidate = validator.New()

type ImageController struct {
	FileRepository repositories.FileRepository
}

// swagger:route POST /image Thing get-thing
//
// # This is the summary for getting a thing by its UUID
//
// This is the description for getting a thing by its UUID. Which can be longer.
//
// responses:
//
//	200: ThingResponse
//	404: ErrorResponse
//	500: ErrorResponse
func (ic *ImageController) CreateImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	result, err := ic.FileRepository.StoreFile(file.Filename, src)

	return c.JSON(http.StatusCreated, responses.ImageResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func (ic *ImageController) GetAllImages(c echo.Context) error {
	images, err := ic.FileRepository.GetAllFiles()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, responses.ImageResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": images}})
}

func (ic *ImageController) DeleteAImage(c echo.Context) error {
	imageId := c.Param("imageId")
	err := ic.FileRepository.DeleteFile(imageId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ImageResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Image with specified ID not found!"}})
	}

	return c.JSON(http.StatusOK, responses.ImageResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Image successfully deleted!"}})
}

func (ic *ImageController) GetAImage(c echo.Context) error {
	imageId := c.Param("imageId")
	content, err := ic.FileRepository.GetFileContent(imageId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	return c.Blob(http.StatusOK, "image/png", content.Bytes())
}

func (ic *ImageController) EditAImage(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, echo.Map{"error": "Method not implemented"})
}
