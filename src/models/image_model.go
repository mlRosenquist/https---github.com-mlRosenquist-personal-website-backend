package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Image
type Image struct {
	// The UUID of the image
	// example: 6204037c-30e6-408b-8aaa-dd8219860b4b
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	// The name of the image
	// example: dog.jpg
	FileName string `json:"filename,omitempty" bson:"filename,omitempty" validate:"required"`
}
