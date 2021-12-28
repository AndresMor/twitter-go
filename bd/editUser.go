package bd

import (
	"Documents/Go/twitter-go/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	record := make(map[string]interface{})
	if len(user.Name) > 0 {
		record["name"] = user.Name
	}
	if len(user.LastName) > 0 {
		record["lastname"] = user.LastName
	}
	record["birthdate"] = user.Birthday
	if len(user.Email) > 0 {
		record["email"] = user.Email
	}
	if len(user.Password) > 0 {
		record["password"] = user.Password
	}
	if len(user.Avatar) > 0 {
		record["avatar"] = user.Avatar
	}
	if len(user.Bio) > 0 {
		record["password"] = user.Bio
	}
	if len(user.Website) > 0 {
		record["website"] = user.Website
	}

	updateStr := bson.M{
		"$set": record,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := col.UpdateOne(ctx, condition, updateStr)
	if err != nil {
		fmt.Println("Error al actualizar regristro " + err.Error())
		return false, err
	}

	return true, nil
}
