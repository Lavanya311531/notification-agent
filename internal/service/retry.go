package service

import "fmt"

func RetryNotification(
	notificationID string,
) {

	for i := 1; i <= 3; i++ {

		fmt.Println(
			"Retry Attempt:",
			i,
		)
	}

	MoveToDeadLetterQueue(
		notificationID,
	)
}
