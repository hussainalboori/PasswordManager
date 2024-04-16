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
	log.Println("New User Data Had Been Recived")

	// Validate the form values
	if username == "" || email == "" || password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Generate a random key
	randomKey, err := data.GenerateRandomKey(32) // Adjust the key length as needed
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error generating random key: %v", err)
		return
	} else {
		log.Println("encrption Key has been created")
	}

	// Hash the password
	passwordHash, err := data.HashPassword(password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error hashing password: %v", err)
		return
	} else {
		log.Println("hased Password Has Benn Created")
	}

	// Convert random key to base64 string

	// Insert the user into the database with the random key
	err = data.InsertUser(username, email, passwordHash, randomKey)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error inserting user into database: %v", err)
		return
	} else {
		log.Println("New User Has Been Added To Database")
		log.Printf("User Name: %v", username)
		log.Printf("email: %v", email)
		log.Println("Hashed Password Has Been added")
		log.Println("encrption Key has Been added")
	}

	// Redirect the user to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
