package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gerrors "github.com/terena-info/terena.godriver/gerror"
	"terena.office/src/repositories"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("authorization")
		token = strings.Replace(token, "Bearer ", "", 1)

		adminRepp := repositories.AdminRepository()

		adminData, err := adminRepp.VerifyAccessToken(token)
		if err != nil {
			gerrors.Panic(http.StatusUnauthorized, gerrors.E{Message: err.Error()})
		}

		ctx.Set("auth", adminData) // Set auth data
		ctx.Next()

	}
}
