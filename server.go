package main

import (
	"log"
	"net/http"
)


func main() {
	port := ":8080"
	log.Printf("Connect to our website through http://localhost%s", port)
	http.ListenAndServe(port, nil)
}