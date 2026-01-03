package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
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

		// TODO: Perform a GET request

		// resp2, err := http.Get(baseURL + "/todos/" + )
		// TODO: Check if the GET request was successful and print the new todo item
		// TODO: Handle potential errors from the GET request
	} else {
		// Handle potential errors from the POST request
		fmt.Printf("Failed to add a new todo\nStatus Code: %d\nError Details: %s\n", resp.StatusCode, string(body))
	}
}
