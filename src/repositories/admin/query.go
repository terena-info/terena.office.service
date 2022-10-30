package adminRepo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"terena.office/src/configs"
)

func (adator *_Adator) FindById(id primitive.ObjectID) interface{} {
	var output []bson.M
	adator.adminOrm.FindById(id).Decode(&output).ErrorMessage(configs.ADMIN_NOT_FOUND)
	return output
}

func (adator *_Adator) FetchAllAdmins(ctx *gin.Context, output interface{}) {

}
