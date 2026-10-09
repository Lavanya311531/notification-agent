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

	log.Println("Loading Key Vault Secret...")
	service.GetSecret("SendGridApiKey")

	log.Println("Creating Notification...")
	service.CreateNotification()

	log.Println("Sending Service Bus Message...")
	service.SendMessage("Hello Azure Service Bus")

	log.Println("Starting Service Bus Consumer...")
	err := servicebus.ReceiveMessage(
		config.GetServiceBusConnectionString(),
		config.GetQueueName(),
	)

	if err != nil {
		log.Printf("Service Bus Error: %v", err)
		panic(err)
	}
}
