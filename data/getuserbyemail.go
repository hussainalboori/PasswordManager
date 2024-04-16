package data

import (
	"database/sql"
	"log"
)

// User represents a user in the system
type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
}

// GetUserByEmail retrieves a user from the database based on their email
func GetUserByEmail(email string) (*User, error) {
	var user User

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT id, username, email, password_hash FROM users WHERE email = ?", email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil, nil
		}
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}
	return &user, nil
}

func GetUserByID(userID int) (*User, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the database to retrieve the user with the specified ID
	query := "SELECT id, username, email FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	// Initialize a User struct to store the retrieved user data
	user := &User{}

	// Scan the row and populate the User struct with the retrieved data
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		// Handle the error, e.g., user not found
		return nil, err
	}

	return user, nil
}
