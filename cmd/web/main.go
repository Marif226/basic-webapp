package main

import (
	"fmt"
	"net/http"
	"github.com/marif226/basic-webapp/pkg/handlers"
)

const portNumber = ":8080"

// Main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}