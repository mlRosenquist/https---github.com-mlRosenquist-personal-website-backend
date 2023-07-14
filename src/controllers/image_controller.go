package controllers

import (
	"io"
	"net/http"
	"personal-website-backend/configs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var imageCollection *mongo.Collection = configs.GetCollection(configs.CLIENT, "images")
var imageValidate = validator.New()

func CreateImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	bucket, _ := gridfs.NewBucket(configs.DB)
	uploadStream, err := bucket.OpenUploadStream(file.Filename)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	if _, err = io.Copy(uploadStream, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"result": "success"})
}
