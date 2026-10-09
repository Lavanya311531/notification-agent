package service

import (
	"fmt"

	"notification-agent/internal/repository"
)

func StartGRPCServer() {
	fmt.Println("gRPC Server Started on :50051")
}

func GetStatus(notificationID string) {

	status := repository.GetNotificationStatus(
		notificationID,
	)

	fmt.Println(
		"Notification Status:",
		status,
	)
}
