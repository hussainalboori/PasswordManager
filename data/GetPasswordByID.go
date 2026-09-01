package data

import (
    "database/sql"
    "log"
)

// PasswordMetadata holds non-sensitive fields for listing
type PasswordMetadata struct {
    Id      int
    Website string
    Username string
}

// GetPasswordsMetadataByUserID returns list of password entries without the decrypted password
func GetPasswordsMetadataByUserID(userID int) ([]PasswordMetadata, error) {
    query := `
        SELECT id, website, username
        FROM passwords
        WHERE user_id = ?;
    `

    db, err := sql.Open("sqlite3", GetDBPath())
    if err != nil {
        log.Printf("Error opening database: %v", err)
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query(query, userID)
    if err != nil {
        log.Printf("Error querying metadata: %v", err)
        return nil, err
    }
    defer rows.Close()

    var list []PasswordMetadata
    for rows.Next() {
        var id int
        var website, username string
        if err := rows.Scan(&id, &website, &username); err != nil {
            log.Printf("Error scanning metadata row: %v", err)
            continue
        }
        list = append(list, PasswordMetadata{Id: id, Website: website, Username: username})
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return list, nil
}

// GetPasswordByID returns the decrypted password for a given password ID and user ID
func GetPasswordByID(passwordID int, userID int) (string, error) {
    query := `SELECT password FROM passwords WHERE id = ? AND user_id = ?;`

    db, err := sql.Open("sqlite3", GetDBPath())
    if err != nil {
        log.Printf("Error opening database: %v", err)
        return "", err
    }
    defer db.Close()

    var encrypted []byte
    row := db.QueryRow(query, passwordID, userID)
    err = row.Scan(&encrypted)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", nil // not found or not authorized
        }
        log.Printf("Error scanning password: %v", err)
        return "", err
    }

    // Need the user's key to decrypt. Fetch key from users table
    key, err := GetKeyByID(userID)
    if err != nil {
        log.Printf("Error fetching key for user: %v", err)
        return "", err
    }
    decrypted, err := Decrypt(encrypted, key)
    if err != nil {
        log.Printf("Error decrypting password: %v", err)
        return "", err
    }
    return decrypted, nil
}
