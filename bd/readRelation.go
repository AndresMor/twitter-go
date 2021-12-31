package bd

import (
	"Documents/Go/twitter-go/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ReadRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relationship")

	condition := bson.M{
		"userId":         t.UserID,
		"userRelationId": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
