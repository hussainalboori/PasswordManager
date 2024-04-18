package handler

import (
	"fmt"
	"log"
	"net/http"
	"password-manger/data"
	"strconv"
	"text/template"
)

func HandleNewPassword(w http.ResponseWriter, r *http.Request) {
	var id int
	_, sessionData, exists := getSession(r)
	userIDstr := sessionData["userID"]
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		log.Printf("Error converting userID to int: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userdata := struct {
		UserID   int
		Username string
	}{
		UserID:   userID,
		Username: sessionData["username"],
	}
	id = userID
	if r.Method == "GET" {
		// Load the template file
		tmpl, err := template.ParseFiles("templates/new.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Printf("Error loading template file: %v\n", err)
			return
		}

		errsen := tmpl.Execute(w, userdata)
		if errsen != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Printf("Error executing template: %v\n", err)
			return
		}
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error parsing form: %v\n", err)
	}

	// Get the form values
	website := r.PostForm.Get("website")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if website == "" || username == "" || password == "" {
		log.Printf("Empty form values\n")
		return
	}

	key, err := data.GetKeyByID(id)
	if err != nil {
		log.Printf("Error getting key by ID: %v\n", err)
		return
	}

	encryptedPassword, err := data.Encrypt([]byte(password), key)
	if err != nil {
		log.Printf("Error encrypting password: %v\n", err)
		return
	}

	
	err = data.InsertPassword(id, website, username, encryptedPassword)
	if err != nil {
		log.Printf("Error inserting password: %v\n", err)
		return
	}else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
	
}
