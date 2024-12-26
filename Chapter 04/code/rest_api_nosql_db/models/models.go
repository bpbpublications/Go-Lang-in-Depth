package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Supplier struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Address string             `json:"address,omitempty" validate:"required"`
	Mobile  string             `json:"mobile,omitempty" validate:"required"`
}
