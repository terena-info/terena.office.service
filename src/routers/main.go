package routers

import "github.com/gin-gonic/gin"

func RegisterRouter(app *gin.Engine) {
	RegisterAdminRouter(app, "/admins")
}
