package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")

	fmt.Println("Connection String:")
	fmt.Println(connStr)

	db, err := sql.Open(
		"postgres",
		connStr,
	)

	if err != nil {
		fmt.Println("Database Error:", err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Database Error:", err)
		return
	}

	DB = db

	var dbName string
	var schema string

	err = DB.QueryRow(
		"SELECT current_database(), current_schema()",
	).Scan(&dbName, &schema)

	if err != nil {
		fmt.Println("Debug Error:", err)
	} else {
		fmt.Println("Connected Database:", dbName)
		var serverIP string

		err = DB.QueryRow(
			"SELECT inet_server_addr()::text",
		).Scan(&serverIP)

		if err == nil {
			fmt.Println("Server IP:", serverIP)
		}

		fmt.Println("Connected Schema:", schema)
	}

	fmt.Println("PostgreSQL Connected Successfully")
}
