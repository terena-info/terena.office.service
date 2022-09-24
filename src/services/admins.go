package services

import (
	"net/http"

	"github.com/bankonly/goginhandlers/src/models"
	"github.com/bankonly/goginhandlers/src/repositories"
	"github.com/bankonly/goutils/response"
	"github.com/bankonly/goutils/utils"
	"github.com/gin-gonic/gin"
)

type _AdminServiceInterface interface {
	Login(*gin.Context)
	SeedAdmin(*gin.Context)
	Profile(*gin.Context)
}

type _AdminServices struct {
	AdminRepo repositories.AdminInterface
}

func (s _AdminServices) Profile(c *gin.Context) {
	res := response.New(c)
	auth := s.AdminRepo.GetAuth(c)
	auth.Password = ""
	res.Success(response.H{Data: auth})
}

func (s _AdminServices) SeedAdmin(c *gin.Context) {
	res := response.New(c)
	msg := s.AdminRepo.SeedAdmin()
	res.Success(response.H{Message: msg})
}

func (s _AdminServices) Login(c *gin.Context) {
	res := response.New(c)

	var body models.AdminLogin
	c.ShouldBind(&body)

	validateErr := utils.ValidateStruct(body)
	if validateErr != nil {
		utils.Panic(http.StatusBadRequest, utils.POption{Message: validateErr.Error()})
	}

	token := s.AdminRepo.Login(body.Email, body.Password)

	res.Success(response.H{Data: token})
}

func RegisterAdminHandler() _AdminServiceInterface {
	var handler _AdminServiceInterface = _AdminServices{
		AdminRepo: repositories.RegisterAdminRepository(),
	}
	return handler
}
