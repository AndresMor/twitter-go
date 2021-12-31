package bd

import (
	"Documents/Go/twitter-go/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(ID string, page int64) ([]models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relationship")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelationId",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursos, err := col.Aggregate(ctx, conditions)
	var result []models.GetTweets
	err = cursos.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
