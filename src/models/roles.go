package models

import (
	"github.com/bankonly/goginhandlers/src/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ROLE_MODEL_NAME = "roles"
	ROLE_INSTANCE   *mongo.Collection
)

type Role struct {
	DefaultField `bson:",inline"`
	RoleName     string `json:"role_name" bson:"role_name"`
}

func InitRole() {
	ROLE_INSTANCE = configs.DBInstance.Collection(ROLE_MODEL_NAME)
}
