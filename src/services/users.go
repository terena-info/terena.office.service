package services

import (
	"net/http"

	"github.com/bankonly/goginhandlers/src/repositories"
	"github.com/gin-gonic/gin"
)

type _UserServiceInterface interface {
	Login(*gin.Context)
}

type _UserServices struct {
	UserRepo repositories.UserInterface
}

func (h _UserServices) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World")
}

func RegisterUserHandler() _UserServiceInterface {
	var handler _UserServiceInterface = _UserServices{
		UserRepo: repositories.RegisterUserRepository(),
	}
	return handler
}
