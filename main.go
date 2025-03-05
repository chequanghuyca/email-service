package main

import (
	"email-service/component/appctx"
	"email-service/middleware"
	"os"

	_ "email-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Email Service Golang
// @version 1.0
// @description This is an email service API built with Golang and Gin.
func main() {
    godotenv.Load()

	appContext := appctx.NewAppContext(os.Getenv("SYSTEM_SECRET_KEY"))

	r := gin.Default()

	r.Use(middleware.Recover(appContext))
	r.Use(middleware.CorsConnect())

	r.Static("/static", "./static")
	
	middleware.ApiServices(appContext, r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}