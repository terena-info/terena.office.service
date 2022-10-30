package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct{}

type DefaultField struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
