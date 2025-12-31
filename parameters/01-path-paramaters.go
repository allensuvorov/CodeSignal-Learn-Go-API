package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Base URL for the API
const baseURL = "http://localhost:8000"

func main() {
    // TODO: Define the ID of the todo item you want to retrieve
    itemID := "2"
    url, _ := url.JoinPath(baseURL, "todos", itemID)
    // TODO: Make a GET request using the path parameter for the specific todo item
    resp, err := http.Get(url)
    // TODO: Check if the request was successful
    if err != nil {
        log.Fatalf("Error fetching the todo with path parameter: %v", err)
    }
    
    // - If successful (status code 200), print the todo item
    if resp.StatusCode == http.StatusOK {
        data, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Error reading body: %v", err)
        }
        fmt.Printf("Todo Item: %s\n", data)
    // - If there's an error, print a message, the status code and the body with the error details
    } else {
        log.Printf("Failed to retch todo. StatusCode: %d", resp.StatusCode)
        data, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Error reading error details: %v", err)
        }
        fmt.Printf("Error details: %s\n", data)
    }
}
