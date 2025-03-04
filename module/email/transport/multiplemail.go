package ginemail

import (
	"email-service/common"
	"email-service/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"

	emailmodel "email-service/module/email/model"
	storagemail "email-service/module/email/storage"
)

func MultipleMail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req emailmodel.MultipleEmailRequest

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err := storagemail.MultipleSendEmail(req.AddressesTo, req.Subject, req.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, emailmodel.ErrSendEmail(err))
		}

		c.JSON(http.StatusOK, emailmodel.SendEmaiSuccess)
	}
}