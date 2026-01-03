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
        "title": "Learn Go http package",
        "description": "Complete a course on Go API calls.",
        "done": false
    }`

	// Complete endpoint URL
	endpoint := baseURL + "/todos"

	// Create a new GET request (this is incorrect)
	resp, err := http.Post(endpoint, "application/json", bytes.NewBufferString(jsonBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Check if the request was successful (incorrect status code check)
	if resp.StatusCode == http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}
		fmt.Println("New todo added successfully!")
		fmt.Println(string(body))
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read error response body: %v", err)
		}
		fmt.Printf("Failed to add a new todo\nStatus Code: %d\nError Details: %s\n", resp.StatusCode, string(body))
	}
}
