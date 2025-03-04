package ginemail

import (
	"email-service/common"
	"email-service/component/appctx"
	"net/http"

	emailmodel "email-service/module/email/model"
	storagemail "email-service/module/email/storage"

	"github.com/gin-gonic/gin"
)

func SingleMail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req emailmodel.EmailRequest

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err := storagemail.SingleSendEmail(req.To, req.Subject, req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, emailmodel.ErrSendEmail(err))
		}

		c.JSON(http.StatusOK, emailmodel.SendEmaiSuccess)
	}
}