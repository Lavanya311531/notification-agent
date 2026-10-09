package main

import (
	"log"

	"notification-agent/internal/config"
	"notification-agent/internal/repository"
	"notification-agent/internal/service"
	"notification-agent/internal/servicebus"
)

func main() {

	log.Println("Notification Agent Started")

	config.LoadEnv()

	log.Println("Connecting to PostgreSQL...")
	repository.ConnectDB()

	log.Println("Starting gRPC Server...")
	service.StartGRPCServer()

	log.Println("Loading secrets...")
	service.GetSecret("SendGridApiKey")

	log.Println("Creating notification...")
	service.CreateNotification()

	log.Println("Sending test message...")
	service.SendMessage("Hello Azure Service Bus")

	log.Println("Starting Service Bus consumer...")

	err := servicebus.ReceiveMessage(
		config.GetServiceBusConnectionString(),
		config.GetQueueName(),
	)

	if err != nil {
		log.Println("Consumer Error:", err)
		panic(err)
	}
}
