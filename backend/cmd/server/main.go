package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

func main() {
    fmt.Println("Backend server started")
    // Get database connection info from environment variables
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    fmt.Println("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable")
    // Build the connection string
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
            dbHost, dbPort, dbUser, dbPassword, dbName)

    // Connect to the PostgreSQL database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
            log.Fatalf("Error opening database connection: %v", err)
    }
    defer db.Close()

    // Ping the database to check the connection
    err = db.Ping()
    if err != nil {
            log.Fatalf("Error connecting to the database: %v", err)
    }

    fmt.Println("Successfully connected to the PostgreSQL database!")
}
