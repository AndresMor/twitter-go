package models

type RequestTweet struct {
	Message string `bson:"message" json:"message"`
}
