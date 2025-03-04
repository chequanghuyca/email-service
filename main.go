package main

import (
	"email-service/component/appctx"
	"email-service/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()

	appContext := appctx.NewAppContext( os.Getenv("SYSTEM_SECRET_KEY"))

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.Static("/static", "./static")
	middleware.ApiServices(appContext, r)

	r.Run()
}