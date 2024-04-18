package data

import (
	"database/sql"
	"log"
)

// Password represents a password entry in the database.
type Password struct {
	Id 	 int
	Website  string
	Username string
	Password string // Stored as byte slice
}

// GetPasswordsByUserID retrieves passwords associated with a user ID from the database.
func GetPasswordsByUserID(userID int, key []byte) ([]Password, error) {
	// SQL query to select passwords by user ID
	query := `
		SELECT id, website, username, password
		FROM passwords
		WHERE user_id = ?;
	`

	// Open a connection to the SQLite3 database
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}
	defer db.Close()

	// Execute the SQL query
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the query results and store passwords in a slice
	var passwords []Password
	for rows.Next() {
		var id int
		var website, username string
		var encryptedPassword []byte
		err := rows.Scan(&id, &website, &username, &encryptedPassword)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		// Decrypt the password
		password, err := Decrypt(encryptedPassword, key)
		if err != nil {
			log.Printf("Error decrypting password: %v", err)
			continue
		}
		passwords = append(passwords, Password{
			Id : id,
			Website:  website,
			Username: username,
			Password: password,
		})
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return passwords, nil
}

