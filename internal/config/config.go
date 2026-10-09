package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using environment variables")
	}
}

func GetServiceBusConnectionString() string {
	return os.Getenv("SERVICE_BUS_CONNECTION_STRING")
}

func GetQueueName() string {
	return os.Getenv("QUEUE_NAME")
}
