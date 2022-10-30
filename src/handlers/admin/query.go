package adminHandle

import (
	"github.com/gin-gonic/gin"
	"github.com/terena-info/terena.godriver/response"
	"github.com/terena-info/terena.godriver/utils"
)

func (adaptor _Adator) FindById(ctx *gin.Context) {
	res := response.New(ctx)
	ID := ctx.Param("id") // Admin ID from Params

	result := adaptor.adminRepo.FindById(utils.StringToObjectId(ID))

	res.Json(response.H{Data: result})
}
