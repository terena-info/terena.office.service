package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gerrors "github.com/terena-info/terena.godriver/gerror"
	"github.com/terena-info/terena.godriver/middlewares"
	"github.com/terena-info/terena.godriver/response"
	"terena.office/src/configs"
	router "terena.office/src/routers"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // Set release mode for gin

	cfg := configs.New() // Load default config

	cfg.LoadEnvironments() // Load environment
	cfg.ConnectDatabase()  // Connect database

	app := gin.New() // Start app

	app.Use(gin.CustomRecovery(middlewares.ErrorRecovery)) // Error handle
	app.Use(gin.Logger())                                  // Enable access log

	router.New(app) // Register router

	app.GET("/", func(ctx *gin.Context) {
		res := response.New(ctx)
		res.Json(response.H{Message: "TerenaAPIAction"})
	})

	app.Use(func(ctx *gin.Context) {
		gerrors.Panic(http.StatusNotFound, gerrors.E{Message: "API_NOT_FOUND"})
	})

	app.Run(fmt.Sprintf(":%s", configs.Env.PORT)) // Start port
}
