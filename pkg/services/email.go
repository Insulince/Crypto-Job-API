package services

import (
	"net/smtp"
	"fmt"
	"os"
	"crypto-jobs/pkg/models/configuration"
)

func SendEmailTo(recipient string, subject string, message string) () {
	config := configuration.GetConfiguration()
	gmailHost := "smtp.gmail.com"
	gmailPort := "587"

	email := "From: " + config.SenderEmailUsername + "\nTo: " + recipient + "\nSubject: " + subject + "\n\n" + message

	err := smtp.SendMail(
		gmailHost+":"+gmailPort,
		smtp.PlainAuth(
			"",
			config.SenderEmailUsername,
			config.SenderEmailPassword,
			gmailHost,
		),
		config.SenderEmailUsername,
		[]string{recipient},
		[]byte(email),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
}
