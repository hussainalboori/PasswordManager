package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HandleNewPassword(w http.ResponseWriter, r *http.Request) {
	_, sessionData, exists := getSession(r)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	data := struct {
		UserID   string
		Username string
	}{
		UserID:   sessionData["userID"],
		Username: sessionData["username"],
	}

	// Load the template file
	tmpl, err := template.ParseFiles("templates/new.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error loading template file: %v\n", err)
		return
	}

	errsen := tmpl.Execute(w, data)
	if errsen != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing template: %v\n", err)
	}

	// Check if the request method is POST
	if r.Method == http.MethodPost {
		// Get the form values
		userID := r.PostForm.Get("user_id")
		website := r.PostForm.Get("website")
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		// userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Printf("Error converting userID to integer: %v\n", err)
			return
		}

		log.Printf("user: %v\n", userID)
		log.Printf("website: %v\n", website)
		log.Printf("username:%v\n", username)
		log.Printf("password: %v\n", password)

		// Insert the new password into the database
		// err = data.InsertPassword(userIDInt, website, username, password)
		// if err != nil {
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	fmt.Printf("Error inserting password into database: %v\n", err)
		// 	return
		// }

		// Redirect the user to a success page
		http.Redirect(w, r, "/success", http.StatusSeeOther)
		return
	}

	// Execute the template with empty data
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing template: %v\n", err)
	}
}
