package data

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "os"
)

func CreateDatabaseIfNotExists(dbFilePath string) error {
	// Check if the database file exists
	_, err := os.Stat(dbFilePath)
	if os.IsNotExist(err) {
		// If the file does not exist, create a new database
		fmt.Printf("Database file '%s' does not exist. Creating a new database...\n", dbFilePath)

		// Create a new database file
		db, err := sql.Open("sqlite3", dbFilePath)
		if err != nil {
			return err
		}
		defer db.Close()

		// Create necessary tables in the new database
		err = createTables(db)
		if err != nil {
			return err
		}

		fmt.Println("Database created successfully.")
	} else if err != nil {
		// If there was an error checking the file existence, return the error
		return err
	} else {
		fmt.Printf("Database file '%s' already exists.\n", dbFilePath)
	}

	return nil
}

func createTables(db *sql.DB) error {
	// SQL statements to create tables
	createUserTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password_hash TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
	createPasswordTableSQL := `
        CREATE TABLE IF NOT EXISTS passwords (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            website TEXT NOT NULL,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
    `

	// Execute SQL statements to create tables
	_, err := db.Exec(createUserTableSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(createPasswordTableSQL)
	if err != nil {
		return err
	}

	return nil
}
