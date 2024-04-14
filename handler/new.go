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
	if r.Method == "GET" {
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
		data := struct {
			UserID   int
			Username string
		}{
			UserID:   userID,
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
			return
		}
		id = userID
		fmt.Printf("User ID 49 : %v\n", id)
	}

	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// 	log.Printf("Method Not Allowed:gg %v\n", r.Method)
	// 	return
	// }
	fmt.Printf("User ID 56 : %v\n", id)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error parsing form: %v\n", err)
	}
	if id == 0 {
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
		data := struct {
			UserID   int
		}{
			UserID:   userID,
		}

		id = data.UserID
	}

	// Get the form values
	website := r.PostForm.Get("website")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	// if website == "" || username == "" || password == "" || id == 0 {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }
	err = data.InsertPassword(id , website, username, password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error inserting password: %v\n", err)
		return
	} else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		log.Printf("Password inserted successfully\n")
	}
}
