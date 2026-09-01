package handler

import (
	"log"
	"net/http"
)

func Handleindex(w http.ResponseWriter, r *http.Request) {
	_, _, exists := getSession(r)
	if !exists {
		renderTemplateFile(w, "templates/index.html", nil)
		log.Println("index page rendered")
		return
	} else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		log.Println("already loged in redirected to dashboard")
	}
}
