package router

import (
	"github.com/gin-gonic/gin"
	adminHandle "terena.office/src/handlers/admin"
)

func AdminRouter(app *gin.Engine, prefix string) {
	r := app.Group(prefix)

	handles := adminHandle.New()

	r.GET("/:id", handles.FindById)
}
