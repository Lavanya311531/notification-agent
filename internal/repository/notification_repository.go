package repository

import (
	"fmt"
)

func GetNotification() (
	string,
	string,
	string,
	string,
) {

	var notificationID string
	var channel string
	var recipient string
	var message string

	query := `
	SELECT notification_id,
	       channel,
	       recipient,
	       message
	FROM notifications
	WHERE status = 'PENDING'
	LIMIT 1
	`

	err := DB.QueryRow(query).Scan(
		&notificationID,
		&channel,
		&recipient,
		&message,
	)

	if err != nil {
		fmt.Println("Database Error:", err)
		return "", "", "", ""
	}

	return notificationID,
		channel,
		recipient,
		message
}

func GetNotificationStatus(
	notificationID string,
) string {

	var status string

	query := `
	SELECT status
	FROM notifications
	WHERE notification_id = $1
	`

	err := DB.QueryRow(
		query,
		notificationID,
	).Scan(&status)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return status
}

func UpdateNotificationStatus(
	notificationID string,
	status string,
) {

	query := `
	UPDATE notifications
	SET status = $1
	WHERE notification_id = $2
	`

	_, err := DB.Exec(
		query,
		status,
		notificationID,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(
		"Notification ID:",
		notificationID,
	)

	fmt.Println(
		"Status updated to",
		status,
	)
}

func SaveMessageLog(
	notificationID string,
	azureMessageID string,
	status string,
) {

	query := `
	INSERT INTO notification_logs
	(
		notification_id,
		azure_message_id,
		status
	)
	VALUES
	(
		$1,
		$2,
		$3
	)
	`

	_, err := DB.Exec(
		query,
		notificationID,
		azureMessageID,
		status,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(
		"Message Log Saved Successfully",
	)

	fmt.Println(
		"Azure Message ID:",
		azureMessageID,
	)
}

func GetStatus(
	notificationID string,
) string {

	var status string

	query := `
	SELECT status
	FROM notifications
	WHERE notification_id = $1
	`

	err := DB.QueryRow(
		query,
		notificationID,
	).Scan(
		&status,
	)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return status
}
