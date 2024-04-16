package handler

import (
	"net/http"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	_, sessionData, exists := getSession(r)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Assuming you have a template file named dashboard.html
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))

	data := struct {
		UserID   string
		Username string
	}{
		UserID:   sessionData["userID"],
		Username: sessionData["username"],
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

