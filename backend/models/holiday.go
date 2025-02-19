package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Holiday struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Date string             `json:"date" bson:"date"`
	Name string             `json:"name" bson:"name"`
}
