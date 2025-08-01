package storagemail

import (
	"log"
	"os"

	emailmodel "email-service/module/email/model"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SingleSendEmail(to string, subject string, body string) error {
	godotenv.Load()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SYSTEM_EMAIL"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		465,
		os.Getenv("SYSTEM_EMAIL"),
		os.Getenv("SYSTEM_EMAIL_SERVER"),
	)

	log.Println("Sending email to", dialer)

	if err := dialer.DialAndSend(mailer); err != nil {
		emailmodel.ErrSendEmail(err)
	}

	log.Println("Email sent successfully to", to)
	return nil
}
