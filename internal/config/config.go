package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func GetServiceBusConnectionString() string {
	return os.Getenv("SERVICE_BUS_CONNECTION_STRING")
}

func GetQueueName() string {
	return os.Getenv("QUEUE_NAME")
}
