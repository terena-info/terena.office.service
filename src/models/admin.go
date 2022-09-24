package models

import (
	"github.com/bankonly/goginhandlers/src/configs"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ADMIN_MODEL_NAME = "users"
	ADMIN_INSTANCE   *mongo.Collection
)

type Admin struct {
	DefaultField `bson:",inline"`
	FullName     string `json:"full_name" bson:"full_name"`
	Password     string `json:"password" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

type AdminLogin struct {
	Email    string `json:"email" form:"email" bson:"email"`
	Password string `json:"password" form:"password" bson:"password"`
}

type AdminJWTPayload struct {
	UserId primitive.ObjectID
	jwt.RegisteredClaims
}

func InitAdmin() {
	ADMIN_INSTANCE = configs.DBInstance.Collection(ADMIN_MODEL_NAME)
}
