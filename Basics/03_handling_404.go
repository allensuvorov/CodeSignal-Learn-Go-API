package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

const baseURL = "http://localhost:8000"

func main() {
    // TODO: Use an incorrect endpoint like '/todoos' to trigger a 404 error
    resp, err := http.Get(baseURL + "/todoos")
    if err != nil {
        log.Fatalf("Failed to get todos: %v", err)
    }

    if resp.StatusCode == http.StatusOK {
        fmt.Println("Todos retrieved successfully:")
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Failed to read response body: %v", err)
        }
        fmt.Println(string(body))
    }

    // TODO: Handle the 404 status code
    if resp.StatusCode == http.StatusNotFound {
    // - Output an appropriate message
        fmt.Println("\nNot Found. The requested resource could not be found on the server.")

    // - Display the error details
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Failed to read error response: %v", err)
        }
        fmt.Printf("Error detailes: %s\n", body)
    }
}
