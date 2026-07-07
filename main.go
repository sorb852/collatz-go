package main

import (
	"log"
	"net/http"
	"collatz-go/routes"
)

func main() {
	log.SetPrefix("[collatz] ")
	log.Println("Setting up the server...")

	http.HandleFunc("/collatz", routes.CollatzHandler)

	// Start that baby UP
	log.Println("Staring the server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
