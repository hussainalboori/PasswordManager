package handler


import (
	"net/http"
)

func Handleindex(w http.ResponseWriter, r *http.Request) {
	// Check if there is a session
	if sessionExists(r) {
		// Redirect to dashboard
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		// Render index.html
		http.ServeFile(w, r, "templates/index.html")
	}
}

func sessionExists(r *http.Request) bool {
	// Check if the session cookie exists
	_, err := r.Cookie("session")
	return err == nil
}