package routers

import (
	"github.com/bankonly/goginhandlers/src/services"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(app *gin.Engine, prefix string) {
	handler := services.RegisterUserHandler()

	r := app.Group(prefix)

	r.POST("/login", handler.Login)
}
