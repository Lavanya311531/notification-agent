package service

import (
	"fmt"

	"notification-agent/internal/repository"
)

func MoveToDeadLetterQueue(
	notificationID string,
) {

	fmt.Println(
		"Moved To DLQ:",
		notificationID,
	)

	repository.UpdateNotificationStatus(
		notificationID,
		"DLQ",
	)

	fmt.Println(
		"Notification ID:",
		notificationID,
	)

	fmt.Println(
		"Status updated to DLQ",
	)
}
