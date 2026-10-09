package service

import (
	"fmt"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(recipient string, message string) {

	from := mail.NewEmail(
		"Lavanya R",
		os.Getenv("SENDER_EMAIL"),
	)

	to := mail.NewEmail(
		"Lavanya",
		recipient,
	)

	subject := "LAVANYA TEST EMAIL - SENDGRID"

	plainTextContent := `
Hello Lavanya,

This is a test email from Notification Agent.

Notification ID: NOTIF-005

Thanks,
Notification Agent
`
	fmt.Println("Sending email to:", recipient)
	email := mail.NewSingleEmail(
		from,
		subject,
		to,
		plainTextContent,
		plainTextContent,
	)

	client := sendgrid.NewSendClient(
		os.Getenv("SENDGRID_API_KEY"),
	)

	response, err := client.Send(email)

	if err != nil {
		fmt.Println("Email Error:", err)
		return
	}

	fmt.Println("Email Sent Successfully")
	fmt.Println("Status Code:", response.StatusCode)
}
