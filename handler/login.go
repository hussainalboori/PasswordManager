package handler

import (
    "encoding/base64"
    "math/rand"
    "net/http"
    "password-manger/data"
    "strconv"
    "text/template"
)

var sessions = make(map[string]map[string]string)

// generateSessionID creates a new, random session ID
func generateSessionID() string {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return ""
    }
    return base64.URLEncoding.EncodeToString(b)
}

// startSession initializes a new session and sets a cookie with the session ID
func startSession(w http.ResponseWriter) string {
    sessionID := generateSessionID()
    sessions[sessionID] = make(map[string]string)
    cookie := http.Cookie{
        Name:     "session_token",
        Value:    sessionID,
        Path:     "/",
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)
    return sessionID
}

// getSession retrieves the session data using the session ID from the cookie
func getSession(r *http.Request) (string, map[string]string, bool) {
    cookie, err := r.Cookie("session_token")
    if err != nil {
        return "", nil, false
    }
    sessionData, exists := sessions[cookie.Value]
    return cookie.Value, sessionData, exists
}

// addToSession adds a key-value pair to a specific session's data
func addToSession(sessionID string, key string, value string) {
    if sessionData, exists := sessions[sessionID]; exists {
        sessionData[key] = value
    }
}

// deleteSession deletes a session and clears the client's session cookie
func deleteSession(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if err != nil {
        // Handle error - no session token found
        return
    }
    sessionID := cookie.Value
    delete(sessions, sessionID)
    cookie.Value = ""
    cookie.Path = "/"
    cookie.HttpOnly = true
    cookie.MaxAge = -1 // Instantly delete the cookie
    http.SetCookie(w, cookie)
}

// Dashboard handler
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

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl := template.Must(template.ParseFiles("templates/login.html"))
        err := tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        return
    }

    // Parse the form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // Get the form values
    email := r.PostForm.Get("email")
    password := r.PostForm.Get("password")

    // Validate the form values
    if email == "" || password == "" {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // Check if the user exists and the password is correct
    user, err := data.GetUserByEmail(email)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    if user == nil || !data.CheckPasswordHash(password, user.PasswordHash) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Start a new session
    sessionID := startSession(w)

    // Add user ID to the session
    addToSession(sessionID, "userID", strconv.Itoa(user.ID))

    // Add username to the session
    addToSession(sessionID, "username", user.Username)

    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
