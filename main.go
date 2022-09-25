package main

import (
	"fmt"
	"net/http"

	"github.com/bankonly/goginhandlers/src/configs"
	"github.com/bankonly/goginhandlers/src/models"
	"github.com/bankonly/goginhandlers/src/routers"
	"github.com/bankonly/goutils/handler"
	"github.com/bankonly/goutils/middlewares"
	"github.com/bankonly/goutils/response"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv() // Load .env file

	configs.InitMongoDB()
	configs.InitAwsConfig()

	models.InitModels() // Load model to variable

	if configs.Env.APP_ENV == configs.Const.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.New()

	app.Use(gin.Logger())

	// Example with mongo transaction
	app.GET("/example-mongo-transaction", handler.SessionHandler(func(hc handler.HandlerContext) {
		// Whole Function Here
		c := hc.Context
		// MongoSessionContext := hc.SessionContext
		res := response.New(c)
		res.Success(response.H{})
	}, configs.DBClient))

	app.Use(gin.CustomRecovery(middlewares.ErrorRecovery))

	app.GET("/", func(ctx *gin.Context) {
		res := response.New(ctx)
		res.Success(response.H{})
	})

	routers.RegisterRouter(app)

	app.Use(func(ctx *gin.Context) {
		res := response.New(ctx)
		res.PanicError(http.StatusNotFound, response.HError{Message: "Api not found"})
	})

	fmt.Printf("Version: %s\n", configs.Env.APP_ENV)
	fmt.Printf("Port: %s\n", configs.Env.PORT)
	app.Run(fmt.Sprintf(":%s", configs.Env.PORT))
}
