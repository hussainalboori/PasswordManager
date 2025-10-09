package handler

import (
	"encoding/json"
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

	// Retrieve password metadata for the current user (no decrypted password)
	metadata, err := data.GetPasswordsMetadataByUserID(ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error retrieving password metadata: %v", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))

	// Data to pass to the template
	dashboardData := struct {
		UserID   string
		Username string
		Passwords []data.PasswordMetadata
	}{
		UserID:   userID,
		Username: username,
		Passwords: metadata,
	}

	// Execute the template
	err = tmpl.Execute(w, dashboardData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}

// HandleGetPassword returns a single decrypted password as JSON for an authenticated user
func HandleGetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	_, sessionData, exists := getSession(r)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := sessionData["userID"]
	ID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// parse id param
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	pid, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	pwd, err := data.GetPasswordByID(pid, ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if pwd == "" {
		// Not found or not authorized
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	resp := map[string]string{"password": pwd}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
