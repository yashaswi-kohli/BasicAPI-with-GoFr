package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	ISBN      string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Publisher string             `json:"publisher,omitempty" bson:"publisher,omitempty"`
}
