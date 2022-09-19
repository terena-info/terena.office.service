package main

import (
	"fmt"
	"net/http"

	"github.com/bankonly/goginhandlers/src/configs"
	"github.com/bankonly/goutils/middlewares"
	"github.com/bankonly/goutils/response"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv() // Load .env file

	configs.InitMongoDB()
	configs.InitAwsConfig()

	if configs.Env.APP_ENV == configs.Const.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.New()

	app.Use(gin.Logger())

	app.Use(gin.CustomRecovery(middlewares.ErrorRecovery))

	app.GET("/", func(ctx *gin.Context) {
		res := response.New(ctx)
		res.Success(response.H{})
	})

	app.Use(func(ctx *gin.Context) {
		res := response.New(ctx)
		res.PanicError(http.StatusNotFound, response.HError{Message: "Api not found"})
	})

	fmt.Printf("Version: %s", configs.Env.APP_ENV)
	app.Run(fmt.Sprintf(":%s\n", configs.Env.PORT))
}
