package main

type Professor struct {
	Id   int
	Name string
}

type Course struct {
	Id          int
	Name        string
	Description string
	College     College
}

type Review struct {
	Id        int
	Rating    int
	Date      string
	Course    Course
	Professor Professor
	User      User
}

type User struct {
	Id   int
	Name string // This could be Username or we can have both Name and CSH Username
	// There's gonna be more I just don't know what
}

type College struct {
	Id   int
	Name string
}
