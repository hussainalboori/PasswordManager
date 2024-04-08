package data

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
    "log"
)

func InsertUser(username, email, passwordHash string) error {
    // SQL statement to insert a new user
    insertUserSQL := `
        INSERT INTO users (username, email, password_hash)
        VALUES (?, ?, ?);
    `

    // Open a connection to the SQLite3 database
    db, err := sql.Open("sqlite3", "data.db")
    if err != nil {
        log.Printf("Error opening database: %v", err)
        return err
    }
    defer db.Close()

    // Execute the SQL statement to insert the user
    _, err = db.Exec(insertUserSQL, username, email, passwordHash)
    if err != nil {
        return err
    }

    return nil
}
