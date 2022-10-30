package adminRepo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/terena-info/terena.godriver/gomgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"terena.office/src/models"
)

type AdminRepositories interface {
	FindById(primitive.ObjectID) interface{}
	FetchAllAdmins(*gin.Context, interface{})
}

type _Adator struct {
	adminOrm gomgo.OrmInterface
}

func New() AdminRepositories {
	var adator AdminRepositories = &_Adator{
		adminOrm: gomgo.New(context.TODO(), models.AdminModelName),
	}
	return adator
}
