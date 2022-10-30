package router

import "github.com/gin-gonic/gin"

func New(app *gin.Engine) {
	AdminRouter(app, "/admin")
}
