package models

type Notification struct {
	NotificationID string `json:"notification_id"`
	Recipient      string `json:"recipient"`
	Channel        string `json:"channel"`
	Message        string `json:"message"`
}
