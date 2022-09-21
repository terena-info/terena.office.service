package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type _UserServiceInterface interface {
	Login(*gin.Context)
}

type _UserServices struct{}

func (h _UserServices) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World")
}

func RegisterUserHandler() _UserServiceInterface {
	var handler _UserServiceInterface = _UserServices{}
	return handler
}
