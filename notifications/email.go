package notifications

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendEmailNotification sends an email notification
// to an admin. It will do nothin if emails are not enabled in .env.
func SendEmailNotification(body string) {

	if os.Getenv("EMAILS_ENABLED") == "" {
		return
	}
	
	// For authentication.
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	address := host + ":" + port
	user := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	appName := os.Getenv("APP_NAME")

	// Mail data.
	from := os.Getenv("MAIL_FROM_ADDRESS")
	to := os.Getenv("MAIL_TO_ADDRESS")
	subject := fmt.Sprintf("Subject: DB Backup - Error while backing up %s\n\n", appName)
	msg := []byte(subject + body)

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(address, auth, from, []string{to}, []byte(msg))
	if err != nil {
		panic(err)
	}

}