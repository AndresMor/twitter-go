package bd

import (
	"Documents/Go/twitter-go/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, pag int64) ([]*models.ResponseTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var result []*models.ResponseTweet

	condition := bson.M{
		"userid": ID,
	}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "date", Value: -1}})
	opts.SetSkip((pag - 1) * 20)

	cursor, err := col.Find(ctx, condition, opts)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var record models.ResponseTweet
		err := cursor.Decode(&record)
		if err != nil {
			return result, false
		}
		result = append(result, &record)
	}

	return result, true
}
