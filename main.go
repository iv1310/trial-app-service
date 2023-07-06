package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func storeTimestampToDB(timestamp time.Time) error {
	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Create the database connection string
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create the table if it doesn't exist
	createTableStmt := `
		CREATE TABLE IF NOT EXISTS access_logs (
			id INT AUTO_INCREMENT PRIMARY KEY,
			access_time DATETIME
		)`
	_, err = db.Exec(createTableStmt)
	if err != nil {
		return err
	}

	// Insert the timestamp into the database
	insertStmt := "INSERT INTO access_logs (access_time) VALUES (?)"
	_, err = db.Exec(insertStmt, timestamp)
	if err != nil {
		return err
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	osVersion := runtime.GOOS
	trialStartDate := time.Date(2023, time.July, 05, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()

	// Store timestamp in MySQL database
	err := storeTimestampToDB(currentDate)
	if err != nil {
		log.Printf("Error storing timestamp: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Current OS Version: %s", osVersion)
	fmt.Fprintf(w, "Xendit - Trial - Candidate Name: Ivan Fransiskus Simatupang\n")
	fmt.Fprintf(w, "Trial Start Date: %s\n", trialStartDate.Format("2006-01-02"))
	fmt.Fprintf(w, "Current Date: %s\n", currentDate.Format("2006-01-02"))
	log.Printf("User accessed the application. Remote Addr: %s\n", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Web app is running. Access it at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
