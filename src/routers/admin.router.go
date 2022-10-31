package router

import (
	"github.com/gin-gonic/gin"
	"terena.office/src/controllers"
	"terena.office/src/middlewares"
	"terena.office/src/validations"
)

func AdminRouter(app *gin.Engine, prefix string) {
	r := app.Group(prefix)

	handles := controllers.AdminController()

	r.GET("/:id", validations.ValidateParamObjectId("id"), handles.FindById)
	r.POST("/login", handles.Login)
	r.GET("/profile", middlewares.AuthMiddleware(), handles.Profile)
}
