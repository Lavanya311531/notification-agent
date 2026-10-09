package service

import (
	"fmt"

	"notification-agent/internal/models"
	"notification-agent/internal/repository"
)

func ProcessNotification(
	notification models.Notification,
) error {

	fmt.Println(
		"Notification ID:",
		notification.NotificationID,
	)

	fmt.Println(
		"Recipient:",
		notification.Recipient,
	)

	fmt.Println(
		"Channel:",
		notification.Channel,
	)

	fmt.Println(
		"Message:",
		notification.Message,
	)

	switch notification.Channel {

	case "EMAIL":

		SendEmail(
			notification.Recipient,
			notification.Message,
		)

	case "SMS":

		SendSMS(
			notification.Recipient,
			notification.Message,
		)

	default:

		fmt.Println(
			"Invalid Channel:",
			notification.Channel,
		)

		return fmt.Errorf(
			"invalid channel",
		)
	}

	repository.UpdateNotificationStatus(
		notification.NotificationID,
		"SUCCESS",
	)

	fmt.Println(
		"Notification ID:",
		notification.NotificationID,
	)

	fmt.Println(
		"Status updated to SUCCESS",
	)

	fmt.Println(
		"Application Insights:",
		notification.Channel+"NotificationProcessed",
	)

	fmt.Println(
		"Audit Log:",
		notification.Channel+
			" Notification Sent Successfully",
	)

	return nil
}
