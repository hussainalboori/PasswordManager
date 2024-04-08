package handler

import (
	"html/template"
	"log"
	"net/http"
	"password-manger/data"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
		return // Return to exit the function after rendering the template
	}

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error parsing form: %v", err)
		return
	}

	// Get the form values
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	// Validate the form values
	if username == "" || email == "" || password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Hash the password
	passwordHash, err := data.HashPassword(password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error hashing password: %v", err)
		return
	}

	// Insert the user into the database
	err = data.InsertUser(username, email, passwordHash)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error inserting user into database: %v", err)
		return
	}

	// Redirect the user to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
