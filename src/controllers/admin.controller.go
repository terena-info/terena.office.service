package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/terena-info/terena.godriver/binding"
	"github.com/terena-info/terena.godriver/response"
	"github.com/terena-info/terena.godriver/utils"
	"terena.office/src/models"
	"terena.office/src/repositories"
)

type _AdminHandlers interface {
	FindById(*gin.Context)
	Login(*gin.Context)
	Profile(*gin.Context)
}

type _AdminAdaptor struct {
	adminRepo repositories.AdminRepositories
}

func (adaptor _AdminAdaptor) FindById(ctx *gin.Context) {
	res := response.New(ctx)
	ID := ctx.Param("id") // Admin ID from Params

	result := adaptor.adminRepo.FindById(utils.StringToObjectId(ID))

	res.Json(response.H{Data: result})
}

func (adaptor _AdminAdaptor) Login(ctx *gin.Context) {
	res := response.New(ctx)

	var body models.Login // Store login body
	ctx.ShouldBind(&body)

	binding.New(body).ValidateStruct().RunError(&binding.RunErrorOption{})

	admin, token := adaptor.adminRepo.Login(body.Email, body.Password)

	res.SetHeader("token", token)
	res.SetCookie("token", token)
	res.Json(response.H{Data: admin})
}

func (adapter _AdminAdaptor) Profile(ctx *gin.Context) {
	res := response.New(ctx)
	admin := adapter.adminRepo.GetAuth(ctx)
	res.Json(response.H{Data: admin})
}

func AdminController() _AdminHandlers {
	var handle _AdminHandlers = _AdminAdaptor{
		adminRepo: repositories.AdminRepository(),
	}
	return handle
}
