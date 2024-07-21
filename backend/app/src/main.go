package main

import (
	"fmt"
	"net/http"
)

var mux *http.ServeMux

func main() {
	// Instantiate the database connection pool
	// pool, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Handle the homepage
		fmt.Fprintf(w, "This is the home page")
	})

	// User pages
	handleUsers()

	// Professor pages
	handleProfessors()

	// Course pages
	handleCourses()

	// College pages
	handleColleges()

	// Review Pages
	handleReviews()

	// Search pages
	handleSearch()

	// I am choosing not to implement search by reviews because that seems backwards

	// Run the server
	http.ListenAndServe(":8080", mux)
}
