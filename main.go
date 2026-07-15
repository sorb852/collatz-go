package main

import (
	"collatz-go/cli"
	"collatz-go/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetPrefix("[collatz] ")

	log.Println("Getting user arguments...")
	args := cli.HandleArguments()
	log.Printf("Using arguments %v", args)

	log.Println("Setting up the server...")

	log.Println("Setting server routes...")
	http.HandleFunc("/collatz", routes.CollatzHandler)

	// Start that baby UP
	log.Printf("Staring the server localhost:%v", args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", args.Port), nil))
}
