package handler

import (
	"log"
	"net/http"
	"password-manger/data"
	"strconv"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	_, sessionData, exists := getSession(r)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve user ID and username from session data
	userID := sessionData["userID"]
	username := sessionData["username"]
	ID, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("Error converting userID to int: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Retrieve encryption key for the current user
	key, err := data.GetKeyByID(ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error retrieving key: %v", err)
		return
	}

	// Retrieve passwords for the current user
	passwords, err := data.GetPasswordsByUserID(ID, key)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error retrieving passwords: %v", err)
		return
	}

	// Assuming you have a template file named dashboard.html
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))

	// Data to pass to the template
	dashboardData := struct {
		UserID    string
		Username  string
		Passwords []data.Password
	}{
		UserID:    userID,
		Username:  username,
		Passwords: passwords,
	}

	// Execute the template
	err = tmpl.Execute(w, dashboardData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}
