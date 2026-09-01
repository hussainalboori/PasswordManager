package handler

import (
	"database/sql"
	"log"
	"net/http"
	"password-manger/data"
	appHandler "password-manger/handler"
	_ "modernc.org/sqlite"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dbFilePath := data.GetDBPath()
	err := data.CreateDatabaseIfNotExists(dbFilePath)
	if err != nil {
		log.Printf("Error creating database: %v", err)
	} else {
		db, err := sql.Open("sqlite", dbFilePath)
		if err != nil {
			log.Printf("Error opening database: %v", err)
		} else {
			db.Close()
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", appHandler.Handleindex)
	mux.HandleFunc("/signup", appHandler.Signup)
	mux.HandleFunc("/login", appHandler.Login)
	mux.HandleFunc("/dashboard", appHandler.Dashboard)
	mux.HandleFunc("/logout", appHandler.Logout)
	mux.HandleFunc("/dashboard/password", appHandler.HandleGetPassword)
	mux.HandleFunc("/dashboard/delete", appHandler.HandleDeletePassword)
	mux.HandleFunc("/dashboard/new", appHandler.HandleNewPassword)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	mux.ServeHTTP(w, r)
}
