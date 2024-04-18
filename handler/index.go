package handler

import (
	"log"
	"net/http"
	"text/template"
)

func Handleindex(w http.ResponseWriter, r *http.Request) {
	_, _, exists := getSession(r)
	if !exists {
		template.Must(template.ParseFiles("templates/index.html")).Execute(w, nil)
		log.Println("index page rendered")
		return
	} else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		log.Println("already loged in redirected to dashboard")
	}
}
