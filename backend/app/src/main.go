package main

import (
	"net/http"
)

func main() {
	// Instantiate the database connection pool
	// pool, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Handle requests
	mux = handleRequests(mux)

	// Run the server
	http.ListenAndServe(":8080", mux)
}
