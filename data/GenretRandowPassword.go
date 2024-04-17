package data

import (
	"math/rand"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	numbers          = "0123456789"
	symbols          = "!@#$%^&*()-=_+[]{}|;':\"<>,.?/~`"
)

func GenerateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Create a combined string of characters to choose from
	allCharacters := uppercaseLetters + lowercaseLetters + numbers + symbols

	// Create a byte slice with the length of the desired password
	password := make([]byte, length)

	// Loop through each position in the password and randomly choose a character
	for i := 0; i < length; i++ {
		password[i] = allCharacters[rand.Intn(len(allCharacters))]
	}

	// Convert the byte slice to a string and return the password
	return string(password)
}
