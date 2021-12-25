package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	LastName string             `bson:"lastname" json:"lastname,omitempty"`
	Birthday time.Time          `bson:"birthdate" json:"birthday,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Bio      string             `bson:"bio" json:"bio,omitempty"`
	Website  string             `bson:"website" json:"webSite,omitempty"`
}
