package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Base URL for the API
	baseURL := "http://localhost:8000"

	// New todo data structured as JSON string
	jsonBody := `{
        "title": "Learn Go http requests",
        "description": "Complete a course on Go API calls.",
        "done": false
    }`

	// Send POST request
	resp, err := http.Post(baseURL+"/todos", "application/json", bytes.NewBufferString(jsonBody))
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Check if the request was successful
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("New todo added successfully!")
		fmt.Println(string(body))

		// verifyAddition
		id := getTodoId(body)
		verifyAddition(id, baseURL)
	} else {
		// Handle potential errors from the POST request
		fmt.Printf("Failed to add a new todo\nStatus Code: %d\nError Details: %s\n", resp.StatusCode, string(body))
	}
}

func getTodoId(body []byte) string {
	type createTodoResponse struct {
		ID int `json: "id"` // why did not use these tags? why was it seen as a number by unmarshal while we are parsing a string?
	}

	var resp createTodoResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return strconv.Itoa(resp.ID)
}

/*
 */

func verifyAddition(id, baseURL string) {
	// TODO: Perform a GET request
	resp, err := http.Get(baseURL + "/todos/" + id)
	// TODO: Handle potential errors from the GET request
	if err != nil {
		log.Fatalf("Could not get a todo, error: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	// TODO: Check if the GET request was successful and print the new todo item
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Varified new todo item created: %s", body)
	} else {
		fmt.Printf("Unexpected Status Code: %d \n", resp.StatusCode)
		fmt.Printf("Error Details: %s\n", body)
	}

}
