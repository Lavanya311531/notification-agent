package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const connStr = "host=localhost port=5432 user=postgres password=Lavanya@123 dbname=notificationdb sslmode=disable"

func GetPendingNotification() (string, string, string) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	defer db.Close()

	var channel string
	var recipient string
	var message string

	err = db.QueryRow(`
        SELECT channel, recipient, message
        FROM  public.notifications
        WHERE status='PENDING'
        LIMIT 1
    `).Scan(&channel, &recipient, &message)

	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}

	return channel, recipient, message
}
func UpdateNotificationStatus(notificationID string, status string) {

	fmt.Println("Notification ID:", notificationID)
	fmt.Println("Status updated to", status)
}
