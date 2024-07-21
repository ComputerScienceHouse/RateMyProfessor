package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var pool *sql.DB // this should be defined in main.main() with secrets

func CreateUser(uid string, name string) *sql.Rows {
	rows, err := pool.Query("INSERT INTO 'users' (uid, name) VALUES ('%s', '%s')", uid, name)
	if err != nil {
		fmt.Printf("Error creating user with uid %s and name %s: %s", uid, name, err.Error())
	}
	return rows
}

func GetUserById(uid string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'users' WHERE uid = '%s'", uid)
	if err != nil {
		fmt.Printf("Error getting user from uid %s: %s", uid, err.Error())
	}
	return rows
}

func GetUsersByName(name string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'users' WHERE name LIKE %'%s'%", name)
	if err != nil {
		fmt.Printf("Error getting users from name %s: %s", name, err.Error())
	}
	return rows
}

func GetUsers() *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'users'")
	if err != nil {
		fmt.Printf("Error getting all users: %s", err.Error())
	}
	return rows
}

func DeleteUser(uid string) {
	_, err := pool.Query("DELETE FROM 'users' WHERE uid = '%s'", uid)
	if err != nil {
		fmt.Printf("Error deleting user with uid %s: %s", uid, err.Error())
	}
}

func CreateCollege(id string, name string) *sql.Rows {
	rows, err := pool.Query("INSERT INTO 'colleges' (id, name) VALUES ('%s', '%s')", id, name)
	if err != nil {
		fmt.Printf("Error creating college with name %s: %s", name, err.Error())
	}
	return rows
}

func GetCollegeById(id string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'colleges' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error getting college from id %s: %s", id, err.Error())
	}
	return rows
}

func GetCollegesByName(name string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'colleges' WHERE name LIKE %'%s'%", name)
	if err != nil {
		fmt.Printf("Error getting colleges from name %s: %s", name, err.Error())
	}
	return rows
}

func GetColleges() *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'colleges'")
	if err != nil {
		fmt.Printf("Error getting all colleges: %s", err.Error())
	}
	return rows
}

func DeleteCollege(id string) {
	_, err := pool.Query("DELETE FROM 'colleges' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error deleting college with id %s: %s", id, err.Error())
	}
}

func CreateCourse(id string, name string, description string, collegeId string) *sql.Rows {
	rows, err := pool.Query("INSERT INTO 'courses' (id, name, description, college) VALUES ('%s', '%s', '%s', '%s')", id, name, description, collegeId)
	if err != nil {
		fmt.Printf("Error creating course with name %s: %s", name, err.Error())
	}
	return rows
}

func GetCourseById(id string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'courses' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error getting course from id %s: %s", id, err.Error())
	}
	return rows
}

func GetCoursesByName(name string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'courses' WHERE name LIKE %'%s'%", name)
	if err != nil {
		fmt.Printf("Error getting courses from name %s: %s", name, err.Error())
	}
	return rows
}

func GetCoursesByCollege(collegeId string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'courses' WHERE college = '%s'", collegeId)
	if err != nil {
		fmt.Printf("Error getting courses from college id %s: %s", collegeId, err.Error())
	}
	return rows
}

func GetCourses() *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'courses'")
	if err != nil {
		fmt.Printf("Error getting all courses: %s", err.Error())
	}
	return rows
}

func DeleteCourse(id string) {
	_, err := pool.Query("DELETE FROM 'courses' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error deleting course with id %s: %s", id, err.Error())
	}
}

func CreateReview(id string, quality float32, difficulty float32, date int, courseId string, professorId string, userId string) *sql.Rows {
	rows, err := pool.Query("INSERT INTO 'reviews' (id, quality, difficulty, date, course, professor, user) VALUES ('%s', '%f', '%f', '%d', '%s', '%s', '%s')", id, quality, difficulty, date, courseId, professorId, userId)
	if err != nil {
		fmt.Printf("Error creating review with id %s: %s", id, err.Error())
	}
	return rows
}

func GetReviewById(id string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'reviews' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error getting review from id %s: %s", id, err.Error())
	}
	return rows
}

func GetReviewsByCourse(courseId string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'reviews' WHERE course = '%s'", courseId)
	if err != nil {
		fmt.Printf("Error getting reviews from course id %s: %s", courseId, err.Error())
	}
	return rows
}

func GetReviewsByProfessor(professorId string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'reviews' WHERE professor = '%s'", professorId)
	if err != nil {
		fmt.Printf("Error getting reviews from professor id %s: %s", professorId, err.Error())
	}
	return rows
}

func GetReviewsByUser(userId string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'reviews' WHERE user = '%s'", userId)
	if err != nil {
		fmt.Printf("Error getting reviews from user id %s: %s", userId, err.Error())
	}
	return rows
}

func GetReviews() *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'reviews'")
	if err != nil {
		fmt.Printf("Error getting all reviews: %s", err.Error())
	}
	return rows
}

func DeleteReview(id string) {
	_, err := pool.Query("DELETE FROM 'reviews' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error deleting review with id %s: %s", id, err.Error())
	}

}

func CreateProfessor(id string, name string) *sql.Rows {
	rows, err := pool.Query("INSERT INTO 'professors' (id, name) VALUES ('%s', '%s')", id, name)
	if err != nil {
		fmt.Printf("Error creating professor with name %s: %s", name, err.Error())
	}
	return rows
}

func GetProfessorById(id string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'professors' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error getting professor from id %s: %s", id, err.Error())
	}
	return rows
}

func GetProfessorsByName(name string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'professors' WHERE name LIKE %'%s'%", name)
	if err != nil {
		fmt.Printf("Error getting professors from name %s: %s", name, err.Error())
	}
	return rows
}

func GetProfessors() *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'professors'")
	if err != nil {
		fmt.Printf("Error getting all professors: %s", err.Error())
	}
	return rows
}

func DeleteProfessor(id string) {
	_, err := pool.Query("DELETE FROM 'professors' WHERE id = '%s'", id)
	if err != nil {
		fmt.Printf("Error deleting professor with id %s: %s", id, err.Error())
	}
}
