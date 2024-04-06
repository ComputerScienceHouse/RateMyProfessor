package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// function to format JSON data
func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

func createUser(name, job string) {
	fmt.Println("Creating user...")

	apiUrl := "https://reqres.in/api/users"
	userData := []byte(`{"name":"` + name + `","job":"` + job + `"}`)

	// create new http request
	request, error := http.NewRequest("POST", apiUrl, bytes.NewBuffer(userData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// send the request
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := formatJSON(responseBody)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
}

func getUser(id string) {
	fmt.Println("Getting user by ID...")

	// make GET request to API to get user by ID
	apiUrl := "https://reqres.in/api/users/" + id
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := formatJSON(responseBody)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
}
