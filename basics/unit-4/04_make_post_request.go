package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: Import necessary packages

func main() {
	// Base URL for the API
	baseURL := "http://localhost:8000"

	// TODO: Define the new todo data as a JSON string
	jsonBody := `{
        "title": "Master Go http requests",
        "description": "Champion GET and POST request",
        "done": false
    }`

	// TODO: Use http.NewRequest to create a POST request to the endpoint {baseURL}/todos
	// - Use the JSON data as the body of your request
	r, err := http.NewRequest("POST", baseURL+"/todos", bytes.NewBufferString(jsonBody))
	if err != nil {
		log.Fatalf("Failed to create a new request: %v", err)
	}
	r.Header.Add("Content-Type", "application/json")
	// TODO: Send the POST request and handle the response
	// - Use http.Client to send the request
	resp, err := new(http.Client).Do(r)
	if err != nil {
		log.Fatalf("Failed to do a post request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	// - Check if the response status code is http.StatusCreated
	if resp.StatusCode == http.StatusCreated {
		//   - If it is, print a success message and the response body
		fmt.Println("Success making an http request")
		fmt.Printf("Response body: %s\n", body)
	} else {
		//   - If not, print an error message with the status code and response body
		fmt.Println("Failed making an http request with code: %v", resp.StatusCode)
		fmt.Printf("Error details: %s\n", body)
	}
}
