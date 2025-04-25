package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Retry for DB to be ready
	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err == nil && db.Ping() == nil {
			break
		}
		fmt.Println("Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")

	// Just a simple query
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS hello (id SERIAL PRIMARY KEY, text TEXT);`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully.")
}

