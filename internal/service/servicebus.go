package service

import "fmt"

func SendMessage(message string) {
	fmt.Println("Message sent to Azure Service Bus:", message)
}

func ReceiveMessage() (string, string, string) {

	fmt.Println("Fetching notification from PostgreSQL")

	channel, recipient, message := GetPendingNotification()

	return channel, recipient, message
}
