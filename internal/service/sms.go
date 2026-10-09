package service

import (
	"fmt"
	"os"

	twilio "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(recipient string, message string) {

	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: os.Getenv("TWILIO_ACCOUNT_SID"),
			Password: os.Getenv("TWILIO_AUTH_TOKEN"),
		},
	)

	params := &api.CreateMessageParams{}

	params.SetTo(recipient)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))

	// Twilio trial approved template
	params.SetBody("sms_appointment_reminders")

	resp, err := client.Api.CreateMessage(params)

	if err != nil {
		fmt.Println("SMS Error:", err)
		return
	}

	fmt.Println("SMS Sent Successfully")
	fmt.Println("Message SID:", *resp.Sid)
}
