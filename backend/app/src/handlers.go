package main

import (
	"fmt"
	"net/http"
)

func handleRequests(mux *http.ServeMux) *http.ServeMux {
	// Register the routes and handlers
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Handle the homepage
		fmt.Fprintf(w, "This is the home page")
	})

	// User pages
	mux = handleUsers(mux)

	// Professor pages
	mux = handleProfessors(mux)

	// Course pages
	mux = handleCourses(mux)

	// College pages
	mux = handleColleges(mux)

	// Review Pages
	mux = handleReviews(mux)

	// Search pages
	mux = handleSearch(mux)
	return mux
}

func handleUsers(mux *http.ServeMux) *http.ServeMux {
	// We could remove this since it is redundant to the search url
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all users
		var users []User
		fmt.Fprintf(w, "This is the users page")
		rows := GetUsers()
		for rows.Next() {
			var user User
			err := rows.Scan(&user.Uid, &user.Name)
			if err != nil {
				fmt.Printf("Error scanning user: %s", err.Error())
			}
			users = append(users, user)
		}
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the user with the given id
		var user User
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the user page for user: %s", id)
		rows := GetUserById(id)
		rows.Next()
		err := rows.Scan(&user.Uid, &user.Name)
		if err != nil {
			fmt.Printf("Error scanning user: %s", err.Error())
		}
	})

	mux.HandleFunc("GET /users/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the reviews for the user with the given id
		var reviews []Review
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the reviews page for user: %s", id)
		rows := GetReviewsByUser(id)
		for rows.Next() {
			var review Review
			err := rows.Scan(&review.Id, &review.Quality, &review.Difficulty, &review.Date, &review.Course, &review.Professor, &review.User)
			if err != nil {
				fmt.Printf("Error scanning review: %s", err.Error())
			}
			reviews = append(reviews, review)
		}
	})
	// this is redundant
	// mux.HandleFunc("GET /users/{id}/reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
	// 	// Return the review with the given id for the user with the given id
	// 	id := r.PathValue("id")
	// 	rid := r.PathValue("rid")
	// 	fmt.Fprintf(w, "This is the review page for user: %s and review: %s", id, rid)
	// })

	// TODO: this needs to be expanded on; i'm not sure how we can create a user without a name input
	mux.HandleFunc("POST /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new user with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new user with id: %s", id)
	})

	// TODO: same here, there needs to be information to update
	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the user with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update user page for user: %s", id)
	})
	return mux
}

func handleColleges(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /colleges", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all colleges
		var colleges []College
		fmt.Fprintf(w, "This is the colleges page")
		rows := GetColleges()
		for rows.Next() {
			var college College
			err := rows.Scan(&college.Id, &college.Name)
			if err != nil {
				fmt.Printf("Error scanning college: %s", err.Error())
			}
			colleges = append(colleges, college)
		}
	})

	mux.HandleFunc("GET /colleges/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the college with the given id
		var college College
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the college page for college: %s", id)
		rows := GetCollegeById(id)
		rows.Next()
		err := rows.Scan(&college.Id, &college.Name)
		if err != nil {
			fmt.Printf("Error scanning college: %s", err.Error())
		}
	})

	mux.HandleFunc("GET /colleges/{id}/courses", func(w http.ResponseWriter, r *http.Request) {
		// Return the courses for the college with the given id
		var courses []Course
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the courses page for college: %s", id)
		rows := GetCoursesByCollege(id)
		for rows.Next() {
			var course Course
			err := rows.Scan(&course.Id, &course.Name, &course.Description, &course.College)
			if err != nil {
				fmt.Printf("Error scanning course: %s", err.Error())
			}
			courses = append(courses, course)
		}
	})

	// this is redundant to get college by id
	// mux.HandleFunc("GET /colleges/{id}/courses/{cid}", func(w http.ResponseWriter, r *http.Request) {
	// 	// Return the course with the given id for the college with the given id

	// 	id := r.PathValue("id")
	// 	cid := r.PathValue("cid")
	// 	fmt.Fprintf(w, "This is the course page for college: %s and course: %s", id, cid)
	// })

	// TODO: expand this to add a name
	mux.HandleFunc("POST /colleges/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new college with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new college with id: %s", id)
	})

	// TODO: expand this to actually add information or disregard
	mux.HandleFunc("PUT /colleges/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the college with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update college page for college: %s", id)
	})
	return mux
}

func handleCourses(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /courses", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all courses
		var courses []Course
		fmt.Fprintf(w, "This is the courses page")
		rows := GetCourses()
		for rows.Next() {
			var course Course
			err := rows.Scan(&course.Id, &course.Name, &course.Description, &course.College)
			if err != nil {
				fmt.Printf("Error scanning course: %s", err.Error())
			}
			courses = append(courses, course)
		}
	})

	mux.HandleFunc("GET /courses/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the course with the given id
		var course Course
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the course page for course: %s", id)
		rows := GetCourseById(id)
		rows.Next()
		err := rows.Scan(&course.Id, &course.Name, &course.Description, &course.College)
		if err != nil {
			fmt.Printf("Error scanning course: %s", err.Error())
		}
	})

	mux.HandleFunc("GET /courses/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the reviews for the course with the given id
		var reviews []Review
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the reviews page for course: %s", id)
		rows := GetReviewsByCourse(id)
		for rows.Next() {
			var review Review
			err := rows.Scan(&review.Id, &review.Quality, &review.Difficulty, &review.Date, &review.Course, &review.Professor, &review.User)
			if err != nil {
				fmt.Printf("Error scanning review: %s", err.Error())
			}
			reviews = append(reviews, review)
		}
	})

	// this is redundant to get review by id
	// mux.HandleFunc("GET /courses/{id}/reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
	// 	// Return the review with the given id for the course with the given id
	// 	id := r.PathValue("id")
	// 	rid := r.PathValue("rid")
	// 	fmt.Fprintf(w, "This is the review page for course: %s and review: %s", id, rid)
	// })

	// TODO: expand this to add a name and description
	mux.HandleFunc("POST /courses/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new course with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new course with id: %s", id)
	})

	// TODO: same here, there needs to be information to update
	mux.HandleFunc("PUT /courses/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the course with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update course page for course: %s", id)
	})
	return mux
}

func handleReviews(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all reviews
		var reviews []Review
		fmt.Fprintf(w, "This is the reviews page")
		rows := GetReviews()
		for rows.Next() {
			var review Review
			err := rows.Scan(&review.Id, &review.Quality, &review.Difficulty, &review.Date, &review.Course, &review.Professor, &review.User)
			if err != nil {
				fmt.Printf("Error scanning review: %s", err.Error())
			}
			reviews = append(reviews, review)
		}
	})

	mux.HandleFunc("GET /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Return the review with the given id
		var review Review
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the review page for review: %s", rid)
		rows := GetReviewById(rid)
		rows.Next()
		err := rows.Scan(&review.Id, &review.Quality, &review.Difficulty, &review.Date, &review.Course, &review.Professor, &review.User)
		if err != nil {
			fmt.Printf("Error scanning review: %s", err.Error())
		}
	})

	// TODO: expand this to add a quality, difficulty, date, course, professor, and user
	mux.HandleFunc("POST /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "Create a new review with id: %s", rid)
	})

	// TODO: same here, there needs to be information to update
	mux.HandleFunc("PUT /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Update the review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the update review page for review: %s", rid)
	})

	mux.HandleFunc("DELETE /reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
		// Delete the review with the given id
		rid := r.PathValue("rid")
		fmt.Fprintf(w, "This is the delete review page for review: %s", rid)
		DeleteReview(rid)

	})
	return mux
}

func handleSearch(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		// Handle the search page
		fmt.Fprintf(w, "This is the search page")
	})

	mux.HandleFunc("GET /search/users", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all users
		var users []User
		fmt.Fprintf(w, "This is the search page for users")
		rows := GetUsers()
		for rows.Next() {
			var user User
			err := rows.Scan(&user.Uid, &user.Name)
			if err != nil {
				fmt.Printf("Error scanning user: %s", err.Error())
			}
			users = append(users, user)
		}
	})

	mux.HandleFunc("GET /search/users/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of users with the given name
		var users []User
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for users with name: %s", name)
		rows := GetUsersByName(name)
		for rows.Next() {
			var user User
			err := rows.Scan(&user.Uid, &user.Name)
			if err != nil {
				fmt.Printf("Error scanning user: %s", err.Error())
			}
			users = append(users, user)
		}
	})

	mux.HandleFunc("GET /search/professors", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all professors
		var professors []Professor
		fmt.Fprintf(w, "This is the search page for professors")
		rows := GetProfessors()
		for rows.Next() {
			var professor Professor
			err := rows.Scan(&professor.Id, &professor.Name)
			if err != nil {
				fmt.Printf("Error scanning professor: %s", err.Error())
			}
			professors = append(professors, professor)
		}
	})

	mux.HandleFunc("GET /search/professors/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of professors with the given name
		var professors []Professor
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for professors with name: %s", name)
		rows := GetProfessorsByName(name)
		for rows.Next() {
			var professor Professor
			err := rows.Scan(&professor.Id, &professor.Name)
			if err != nil {
				fmt.Printf("Error scanning professor: %s", err.Error())
			}
			professors = append(professors, professor)
		}
	})

	mux.HandleFunc("GET /search/courses", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all courses
		var courses []Course
		fmt.Fprintf(w, "This is the search page for courses")
		rows := GetCourses()
		for rows.Next() {
			var course Course
			err := rows.Scan(&course.Id, &course.Name, &course.Description, &course.College)
			if err != nil {
				fmt.Printf("Error scanning course: %s", err.Error())
			}
			courses = append(courses, course)
		}
	})

	mux.HandleFunc("GET /search/courses/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of courses with the given name
		var courses []Course
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for courses with name: %s", name)
		rows := GetCoursesByName(name)
		for rows.Next() {
			var course Course
			err := rows.Scan(&course.Id, &course.Name, &course.Description, &course.College)
			if err != nil {
				fmt.Printf("Error scanning course: %s", err.Error())
			}
			courses = append(courses, course)
		}
	})

	mux.HandleFunc("GET /search/colleges", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all colleges
		var colleges []College
		fmt.Fprintf(w, "This is the search page for colleges")
		rows := GetColleges()
		for rows.Next() {
			var college College
			err := rows.Scan(&college.Id, &college.Name)
			if err != nil {
				fmt.Printf("Error scanning college: %s", err.Error())
			}
			colleges = append(colleges, college)
		}
	})

	mux.HandleFunc("GET /search/colleges/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of colleges with the given name
		var colleges []College
		name := r.PathValue("name")
		fmt.Fprintf(w, "This is the search page for colleges with name: %s", name)
		rows := GetCollegesByName(name)
		for rows.Next() {
			var college College
			err := rows.Scan(&college.Id, &college.Name)
			if err != nil {
				fmt.Printf("Error scanning college: %s", err.Error())
			}
			colleges = append(colleges, college)
		}
	})
	// I am choosing not to implement search by reviews because that seems backwards
	return mux
}

func handleProfessors(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /professors", func(w http.ResponseWriter, r *http.Request) {
		// Return the list of all professors
		var professors []Professor
		fmt.Fprintf(w, "This is the professors page")
		rows := GetProfessors()
		for rows.Next() {
			var professor Professor
			err := rows.Scan(&professor.Id, &professor.Name)
			if err != nil {
				fmt.Printf("Error scanning professor: %s", err.Error())
			}
			professors = append(professors, professor)
		}
	})

	mux.HandleFunc("GET /professors/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Return the professor with the given id
		var professor Professor
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the professor page for professor: %s", id)
		rows := GetProfessorById(id)
		rows.Next()
		err := rows.Scan(&professor.Id, &professor.Name)
		if err != nil {
			fmt.Printf("Error scanning professor: %s", err.Error())
		}
	})

	mux.HandleFunc("GET /professors/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Return the reviews for the professor with the given id
		var reviews []Review
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the reviews page for professor: %s", id)
		rows := GetReviewsByProfessor(id)
		for rows.Next() {
			var review Review
			err := rows.Scan(&review.Id, &review.Quality, &review.Difficulty, &review.Date, &review.Course, &review.Professor, &review.User)
			if err != nil {
				fmt.Printf("Error scanning review: %s", err.Error())
			}
			reviews = append(reviews, review)
		}
	})

	// this is redundant
	// mux.HandleFunc("GET /professors/{id}/reviews/{rid}", func(w http.ResponseWriter, r *http.Request) {
	// 	// Return the review with the given id for the professor with the given id
	// 	id := r.PathValue("id")
	// 	rid := r.PathValue("rid")
	// 	fmt.Fprintf(w, "This is the review page for professor: %s and review: %s", id, rid)
	// })

	// TODO: this needs to be expanded on; i'm not sure how we can create a professor without a name input
	mux.HandleFunc("POST /professors/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Create a new professor with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "Create a new professor with id: %s", id)
	})

	// TODO: same here, there needs to be information to update
	mux.HandleFunc("PUT /professors/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Update the professor with the given id
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is the update professor page for professor: %s", id)
	})
	return mux
}
