package handler

import (
	"log"
	"net/http"
	"password-manger/data"
	"strconv"
)

func HandleDeletePassword(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	// Get the value of the 'id' parameter from the form
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	// Convert the 'id' parameter to an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}
	delete := data.DeletePasswordByID(idInt)
    if delete != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }else {
        log.Println("Password deleted successfully")
    }
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}
