package handler

import (
	"database/sql"
	"log"
	"net/http"
	"password-manger/data"
	"password-manger/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dbFilePath := data.GetDBPath()
	err := data.CreateDatabaseIfNotExists(dbFilePath)
	if err != nil {
		log.Printf("Error creating database: %v", err)
	} else {
		db, err := sql.Open("sqlite3", dbFilePath)
		if err != nil {
			log.Printf("Error opening database: %v", err)
		} else {
			db.Close()
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handleindex)
	mux.HandleFunc("/signup", handler.Signup)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/dashboard", handler.Dashboard)
	mux.HandleFunc("/logout", handler.Logout)
	mux.HandleFunc("/dashboard/password", handler.HandleGetPassword)
	mux.HandleFunc("/dashboard/delete", handler.HandleDeletePassword)
	mux.HandleFunc("/dashboard/new", handler.HandleNewPassword)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	mux.ServeHTTP(w, r)
}
