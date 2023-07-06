package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mackerelio/go-osstat/cpu"
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

func getMemoryUsage() float64 {
	memFile := "/sys/fs/cgroup/memory/memory.usage_in_bytes"

	// Read the value from the file
	data, err := os.ReadFile(memFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse the value as an integer
	value, err := parseMemoryUsage(string(data))
	if err != nil {
		log.Fatalf("Error parsing value: %v", err)
	}

	// Convert the value to megabytes
	memUsage := float64(value) / (1024 * 1024)

	return memUsage
}

func getCPUUsage() float64 {
	initialStats, err := cpu.Get()
	if err != nil {
		log.Fatalf("Error getting initial CPU stats: %v", err)
	}

	time.Sleep(1 * time.Second)

	newStats, err := cpu.Get()
	if err != nil {
		log.Fatalf("Error getting new CPU stats: %v", err)
	}

	totalTimeDiff := newStats.Total - initialStats.Total
	idleTimeDiff := newStats.Idle - initialStats.Idle

	cpuUsage := 100.0 * (1.0 - float64(idleTimeDiff)/float64(totalTimeDiff))

	return cpuUsage
}

func parseMemoryUsage(data string) (int64, error) {
	data = strings.TrimSpace(data)
	value, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
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

	cpuUsage := getCPUUsage()
	memUsage := getMemoryUsage()

	fmt.Fprintf(w, "Current OS Version: %s\n", osVersion)
	fmt.Fprintf(w, "Xendit - Trial - Candidate Name: Ivan Fransiskus Simatupang\n")
	fmt.Fprintf(w, "Trial Start Date: %s\n", trialStartDate.Format("2006-01-02"))
	fmt.Fprintf(w, "Current Date: %s\n", currentDate.Format("2006-01-02"))
	fmt.Fprintf(w, "CPU Usage: %.2f%%\n", cpuUsage)
	fmt.Fprintf(w, "Memory Usage: %.2f MB\n", memUsage)

	log.Printf("User accessed the application. Remote Addr: %s\n", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Web app is running. Access it at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
