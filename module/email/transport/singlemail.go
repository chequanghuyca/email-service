package ginemail

import (
	"email-service/common"
	"email-service/component/appctx"
	"net/http"

	emailmodel "email-service/module/email/model"
	storagemail "email-service/module/email/storage"

	"github.com/gin-gonic/gin"
)

// @Summary Send a single email
// @Description Send an email to a single recipient
// @Tags email
// @Accept json
// @Produce json
// @Param email body emailmodel.EmailRequest true "Email data"
// @Success 200 {string} string "Email sent successfully"
// @Router /api/email/single [post]
func SingleMail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req emailmodel.EmailRequest

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err := storagemail.SingleSendEmail(req.To, req.Subject, req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, emailmodel.ErrSendEmail(err))
			return
		}

		c.JSON(http.StatusOK, emailmodel.SendEmaiSuccess)
	}
}
