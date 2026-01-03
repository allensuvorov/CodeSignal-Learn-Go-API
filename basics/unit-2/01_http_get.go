package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

// TODO: Define the base URL for the API
const baseURL = "http://localhost:8000"

func main() {
    // TODO: Fetch all todos using the Get method
    resp, err := http.Get(baseURL + "/todos")
    if err != nil {
        log.Fatalf("Failed to get todos: %v", err)
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read body: %v", err)
    }
    
    // TODO: Print raw response
    fmt.Println("Raw resonse:")
    fmt.Println(string(body))
}
