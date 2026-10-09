package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")

	log.Println("Attempting PostgreSQL connection...")
	log.Println("Connection String:")
	log.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to open PostgreSQL connection:", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping PostgreSQL:", err)
		panic(err)
	}

	DB = db

	log.Println("✅ PostgreSQL connected successfully")
}
