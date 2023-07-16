package repositories

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"personal-website-backend/configs"
	"personal-website-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var imageCollection *mongo.Collection = configs.GetCollection(configs.CLIENT, "images")

type MongoFileRepository struct {
}

func (m *MongoFileRepository) StoreFile(fileName string, fileContent io.Reader) (*models.Image, error) {
	uniqueFileName := fmt.Sprintf("%s-%s", time.Now().Format("20060102150405"), fileName)

	bucket, _ := gridfs.NewBucket(configs.DB)
	uploadStream, err := bucket.OpenUploadStream(uniqueFileName)
	if err != nil {
		return nil, err
	}
	defer uploadStream.Close()

	if _, err = io.Copy(uploadStream, fileContent); err != nil {
		return nil, err
	}

	img := &models.Image{
		FileName: uniqueFileName,
	}

	result, err := imageCollection.InsertOne(context.Background(), img)
	if err != nil {
		return nil, err
	}

	img.Id = result.InsertedID.(primitive.ObjectID)
	return img, nil
}

func (m *MongoFileRepository) GetAllFiles() ([]models.Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var images []models.Image
	defer cancel()

	results, err := imageCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleImage models.Image
		if err = results.Decode(&singleImage); err != nil {
			return nil, err
		}

		images = append(images, singleImage)
	}

	return images, nil
}

func (m *MongoFileRepository) GetFileContent(imageId string) (bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var image models.Image
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(imageId)

	err := imageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&image)

	if err != nil {
		return *bytes.NewBufferString(""), err
	}

	bucket, _ := gridfs.NewBucket(configs.DB)

	var buf bytes.Buffer

	_, err = bucket.DownloadToStreamByName(image.FileName, &buf)

	if err != nil {
		return *bytes.NewBufferString(""), err
	}

	return buf, nil
}

func (m *MongoFileRepository) DeleteFile(imageId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(imageId)

	// First, find the image in the collection
	var img models.Image
	err := imageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&img)
	if err != nil {
		return err
	}

	// Get a handle to the fs.files collection
	fsFiles := configs.DB.Collection("fs.files")

	// Find the GridFS file by filename
	var file bson.M
	err = fsFiles.FindOne(ctx, bson.M{"filename": img.FileName}).Decode(&file)
	if err != nil {
		return err
	}

	// Next, delete the file from GridFS
	bucket, _ := gridfs.NewBucket(configs.DB)
	err = bucket.Delete(file["_id"].(primitive.ObjectID))
	if err != nil {
		return err
	}

	// Finally, delete the document from the collection
	result, err := imageCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return err
	}

	return nil
}
