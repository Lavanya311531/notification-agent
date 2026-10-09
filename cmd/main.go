package main

import (
	"notification-agent/internal/config"
	"notification-agent/internal/repository"
	"notification-agent/internal/service"
	"notification-agent/internal/servicebus"
)

func main() {

	// Load Environment Variables
	config.LoadEnv()

	// Connect PostgreSQL
	repository.ConnectDB()

	// Start gRPC Server
	service.StartGRPCServer()

	// Load Secret
	service.GetSecret("SendGridApiKey")

	// Create Notification
	service.CreateNotification()

	// Send Service Bus Test Message
	service.SendMessage("Hello Azure Service Bus")

	/*service.GetStatus(
		"NOTIF-003",
	)

	)*/

	// fmt.Println(
	//     "Notification Status:",
	//     service.GetStatus("NOTIF-003"),
	// )

	// service.RetryNotification("NOTIF-003")

	// Start Consumer
	err := servicebus.ReceiveMessage(

		// service.RetryNotification("NOTIF-003")

		// Start Consumer

		config.GetServiceBusConnectionString(),
		config.GetQueueName(),
	)

	if err != nil {
		panic(err)
	}
}
