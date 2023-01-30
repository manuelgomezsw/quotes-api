package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	Author      string             `json:"author" bson:"author"`
	Message     string             `json:"message" bson:"message"`
	DateCreated primitive.DateTime `json:"date_created" bson:"dateCreated"`
	Tags        []string           `json:"tags" bson:"tags"`
}
