package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/terena-info/terena.godriver/middlewares"
	"github.com/terena-info/terena.godriver/response"
	"terena.office/src/configs"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // Set release mode for gin

	cfg := configs.New() // Load default config

	cfg.LoadEnvironments() // Load environment
	cfg.ConnectDatabase()  // Connect database

	app := gin.New() // Start app

	app.Use(gin.CustomRecovery(middlewares.ErrorRecovery)) // Error handle
	app.Use(gin.Logger())                                  // Enable access log

	app.GET("/", func(ctx *gin.Context) {
		res := response.New(ctx)
		res.Json(response.H{Message: "TerenaAPI"})
	})

	app.Run(fmt.Sprintf(":%s", configs.Env.PORT)) // Start port
}
