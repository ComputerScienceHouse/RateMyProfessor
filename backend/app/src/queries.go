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
		fmt.Printf("Error creating user with uid %s and name%s: %s", uid, name, err.Error())
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

func SearchUsersByName(name string) *sql.Rows {
	rows, err := pool.Query("SELECT * FROM 'users' WHERE name = '%s'", name)
	if err != nil {
		fmt.Printf("Error getting users from name %s: %s", name, err.Error())
	}
	return rows
}
