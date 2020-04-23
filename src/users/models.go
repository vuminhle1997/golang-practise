package users

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string `bson:"firstname,omitempty"`
	Lastname string `bson:"lastname,omitempty"`
	Email string `bson:"email,omitempty"`
	DOB Date `bson:"dob,omitempty"`
	Password string `bson:"password,omitempty"`
	Active bool `bson:"active,omitempty"`
}