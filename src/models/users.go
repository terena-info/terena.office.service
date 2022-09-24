package models

import (
	"github.com/bankonly/goginhandlers/src/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	USER_MODEL_NAME = "users"
	USER_INSTANCE   *mongo.Collection
)

type User struct {
	DefaultField `bson:",inline"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

type UserOnCreate struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

type Login struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func InitUser() {
	USER_INSTANCE = configs.DBInstance.Collection(USER_MODEL_NAME)
}
