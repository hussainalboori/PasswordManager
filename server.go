package main

import (
	"database/sql"
	"log"
	"net/http"
	"password-manger/data"
	"password-manger/handler"
)

func main() {
	dbFilePath := "data.db" // SQLite database file path

	err := data.CreateDatabaseIfNotExists(dbFilePath) // Create the database if it does not exist
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
		return
	}
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()
	port := ":8080" // Port to run the server on
	http.HandleFunc("/signup", handler.Signup)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/dashboard", handler.Dashboard)
	http.HandleFunc("/logout", handler.Logout)
	log.Printf("Connect to our website through http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
