package ginemail

import (
	"email-service/common"
	"email-service/component/appctx"
	"email-service/helper"
	emailmodel "email-service/module/email/model"
	storagemail "email-service/module/email/storage"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ResponseEmailPortfolio(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		godotenv.Load()

		var req emailmodel.EmailResponsePortfolio

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dataSendContact := helper.MailResponseData{
			Name: req.Name,
			MyPhone: os.Getenv("SYSTEM_PHONE_NUMBER"),
			MyEmail: os.Getenv("SYSTEM_EMAIL"),
		}

		err := storagemail.SingleSendEmail(
			req.Email, 
			helper.GetSubjectMailResponse(),  
			helper.GetBodyMailResponse(dataSendContact),
		)
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, emailmodel.ErrSendEmail(err))
		}

		errSendMe := storagemail.ResponseMeEmail(req.Message)

		if errSendMe != nil {
			c.JSON(http.StatusInternalServerError, emailmodel.ErrSendEmail(errSendMe))
		}

		c.JSON(http.StatusOK, emailmodel.SendEmaiSuccess)
	}
}