package main

import (
	"fmt"
	"net/http"
)

func handleUsers() {
	// We could remove this since it is redundant to the search url
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all users
		fmt.Fprintf(w, "This is the users page")
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the user with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the user page for user: %s", id)
	})

	mux.HandleFunc("GET /users/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the reviews for the user with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the reviews page for user: %s", id)
	})

	mux.HandleFunc("GET /users/{id}/reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Return the review with the given id for the user with the given id
		id := r.PathValue("id")
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the review page for user: %s and review: %s", id, rid)
	})

	mux.HandleFunc("POST /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new user with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new user with id: %s", id)
	})

	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the user with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update user page for user: %s", id)
	})
}

func handleColleges() {
	mux.HandleFunc("GET /colleges", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all colleges
		fmt.Fprintf(w, "This is the colleges page")
	})

	mux.HandleFunc("GET /colleges/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the college with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the college page for college: %s", id)
	})

	mux.HandleFunc("GET /colleges/{id}/courses", func(w http.ResponseWriter, r *http.Request) {
		// Return the courses for the college with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the courses page for college: %s", id)
	})

	mux.HandleFunc("GET /colleges/{id}/courses/{cid}", func(w http.ResponseWriter, r *http.Request) {
		// Return the course with the given id for the college with the given id
		id := r.PathValue("id")
		cid := r.PathValue("cid")
		fmt.Fprintf(w, "This is the course page for college: %s and course: %s", id, cid)
	})

	mux.HandleFunc("POST /colleges/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new college with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new college with id: %s", id)
	})

	mux.HandleFunc("PUT /colleges/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the college with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update college page for college: %s", id)
	})
}

func handleCourses() {
	mux.HandleFunc("GET /courses", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all courses
		fmt.Fprintf(w, "This is the courses page")
	})

	mux.HandleFunc("GET /courses/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the course with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the course page for course: %s", id)
	})

	mux.HandleFunc("GET /courses/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the reviews for the course with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the reviews page for course: %s", id)
	})

	mux.HandleFunc("GET /courses/{id}/reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Return the review with the given id for the course with the given id
		id := r.PathValue("id")
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the review page for course: %s and review: %s", id, rid)
	})

	mux.HandleFunc("POST /courses/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new course with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new course with id: %s", id)
	})

	mux.HandleFunc("PUT /courses/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the course with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update course page for course: %s", id)
	})
}

func handleReviews() {
	mux.HandleFunc("GET /reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all reviews
		fmt.Fprintf(w, "This is the reviews page")
	})

	mux.HandleFunc("GET /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Return the review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the review page for review: %s", rid)
	})

	mux.HandleFunc("POST /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "Create a new review with id: %s", rid)
	})

	mux.HandleFunc("PUT /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Update the review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the update review page for review: %s", rid)
	})

	mux.HandleFunc("DELETE /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Delete the review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the delete review page for review: %s", rid)
	})
}

func handleSearch() {
	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		// Handle the search page
		fmt.Fprintf(w, "This is the search page")
	})

	mux.HandleFunc("GET /search/users", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all users
		fmt.Fprintf(w, "This is the search page for users")
	})

	mux.HandleFunc("GET /search/users/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of users with the given name
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for users with name: %s", name)
	})

	mux.HandleFunc("GET /search/professors", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all professors
		fmt.Fprintf(w, "This is the search page for professors")
	})

	mux.HandleFunc("GET /search/professors/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of professors with the given name
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for professors with name: %s", name)
	})

	mux.HandleFunc("GET /search/courses", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all courses
		fmt.Fprintf(w, "This is the search page for courses")
	})

	mux.HandleFunc("GET /search/courses/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of courses with the given name
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for courses with name: %s", name)
	})

	mux.HandleFunc("GET /search/colleges", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all colleges
		fmt.Fprintf(w, "This is the search page for colleges")
	})

	mux.HandleFunc("GET /search/colleges/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of colleges with the given name
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for colleges with name: %s", name)
	})
}

func handleProfessors() {
	mux.HandleFunc("GET /professors", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all professors
		fmt.Fprintf(w, "This is the professors page")
	})

	mux.HandleFunc("GET /professors/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the professor with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the professor page for professor: %s", id)
	})

	mux.HandleFunc("GET /professors/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the reviews for the professor with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the reviews page for professor: %s", id)
	})

	mux.HandleFunc("GET /professors/{id}/reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Return the review with the given id for the professor with the given id
		id := r.PathValue("id")
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the review page for professor: %s and review: %s", id, rid)
	})

	mux.HandleFunc("POST /professors/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new professor with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new professor with id: %s", id)
	})

	mux.HandleFunc("PUT /professors/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the professor with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update professor page for professor: %s", id)
	})
}
