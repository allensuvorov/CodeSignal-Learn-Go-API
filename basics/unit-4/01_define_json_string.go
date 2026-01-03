package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Base URL for the API
const baseURL = "http://localhost:8000"

func main() {
	// TODO: Define the JSON string containing the new todo data
	jsonBody := `{
        "title": "Learn Go http requests",
        "description": "Complete a course on Go API calls.",
        "done": false
        }`
	// Complete endpoint URL
	endpoint := baseURL + "/todos"

	// TODO: Insert the JSON data into the body of the request
	// Send the POST request directly
	resp, err := http.Post(endpoint, "application/json", bytes.NewBufferString(jsonBody))
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}

	// Check if the request was successful
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
