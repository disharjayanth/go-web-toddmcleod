package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct models user data
type User struct {
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
	ID     primitive.ObjectID `json:"id" bson:"_id"`
}
