package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	// "encoding/base64"
	"fmt"
	"io"
)

// GenerateRandomKey generates a random key of specified length
func GenerateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Encrypt encrypts a plaintext with the secret key using AES
func Encrypt(plaintext []byte, secretKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// Decrypt decrypts a ciphertext with the secret key using AES
func Decrypt(ciphertext []byte, secretKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

// func main() {
// 	keyLength := 32 // Length of the AES key (in bytes)
// 	key, err := GenerateRandomKey(keyLength)
// 	if err != nil {
// 		fmt.Println("Error generating key:", err)
// 		return
// 	}

// 	fmt.Println("Key:", key)

// 	plaintext := []byte("password123")
// 	encrypted, err := Encrypt(plaintext, key)
// 	if err != nil {
// 		fmt.Println("Encryption error:", err)
// 		return
// 	}

// 	// Store 'encrypted' in the database

// 	decrypted, err := Decrypt(encrypted, key)
// 	if err != nil {
// 		fmt.Println("Decryption error:", err)
// 		return
// 	}

// 	fmt.Println("Original:", string(plaintext))
// 	fmt.Println("Decrypted:", string(decrypted))
// }
