package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
	Nicknames []string `json:"nicknames,omitempty" validate:"required"`
	Born int32 `json:"born,omitempty" validate:"required"`
}