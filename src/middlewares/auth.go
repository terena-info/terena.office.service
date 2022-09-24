package middlewares

import (
	"net/http"
	"strings"

	"github.com/bankonly/goginhandlers/src/repositories"
	"github.com/bankonly/goutils/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	adminRepo := repositories.RegisterAdminRepository()

	token := c.GetHeader("authorization")
	token = strings.Replace(token, "Bearer ", "", 1)

	println(token)

	admin, err := adminRepo.VerifyAccessToken(token)
	if err != nil {
		utils.Panic(http.StatusUnauthorized, utils.POption{Message: err.Error()})
	}

	c.Set("auth", admin)

	c.Next()
}
