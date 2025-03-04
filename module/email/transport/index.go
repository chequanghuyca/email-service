package ginemail

import (
	"email-service/component/appctx"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(appCtx appctx.AppContext, router *gin.RouterGroup) {
	email := router.Group("/email")

	email.POST("/single", SingleMail(appCtx))
	email.POST("/multiple", MultipleMail(appCtx))
}