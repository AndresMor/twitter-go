package models

type Relation struct {
	UserID         string `bson:"userId" json:"userId"`
	UserRelationID string `bson:"userRelationId" json:"userRelationId"`
}
