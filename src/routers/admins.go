package routers

import (
	"github.com/bankonly/goginhandlers/src/middlewares"
	"github.com/bankonly/goginhandlers/src/services"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(app *gin.Engine, prefix string) {
	handler := services.RegisterAdminHandler()

	r := app.Group(prefix)

	r.POST("/login", handler.Login)
	r.GET("/seed-admin", handler.SeedAdmin)
	r.GET("/profile", middlewares.Auth, handler.Profile)
}
