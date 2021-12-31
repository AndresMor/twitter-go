package bd

import (
	"Documents/Go/twitter-go/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllRelations(ID string, page int64, search string, types string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var results []*models.User
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var finded, includ bool
	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		includ = false
		finded, err = ReadRelation(r)
		if types == "new" && finded == false {
			includ = true
		}

		if types == "follow" && finded == true {
			includ = true
		}

		if r.UserRelationID == ID {
			includ = false
		}

		if includ == true {
			s.Password = ""
			s.Bio = ""
			s.Website = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
