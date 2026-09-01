package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"password-manger/data"
	"password-manger/handler"
	"password-manger/static"
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

	portStr := os.Getenv("PORT")
	var listener net.Listener

	if portStr != "" {
		if portStr[0] != ':' {
			portStr = ":" + portStr
		}
		var err error
		listener, err = net.Listen("tcp", portStr)
		if err != nil {
			log.Fatalf("Failed to listen on specified PORT %s: %v", portStr, err)
		}
	} else {
		portsToTry := []string{":8080", ":8081", ":8082", ":8083"}
		for _, p := range portsToTry {
			l, err := net.Listen("tcp", p)
			if err == nil {
				listener = l
				portStr = p
				break
			}
		}
		if listener == nil {
			log.Fatalf("Could not find an available port (tried 8080-8083). Please set PORT env variable.")
		}
	}

	http.HandleFunc("/", handler.Handleindex)
	http.HandleFunc("/signup", handler.Signup)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/dashboard", handler.Dashboard)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/dashboard/password", handler.HandleGetPassword)
	http.HandleFunc("/dashboard/delete", handler.HandleDeletePassword)
	http.HandleFunc("/dashboard/new", handler.HandleNewPassword)
	http.Handle("/static/", http.StripPrefix("/static", static.Handler()))

	log.Printf("Connect to our website through http://localhost%s", portStr)
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}



