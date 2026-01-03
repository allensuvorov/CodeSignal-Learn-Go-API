package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

const baseURL = "http://localhost:8000"

func main() {
    resp, err := http.Get(baseURL + "/todos")
    if err != nil {
        log.Fatalf("Failed to get todos: %v", err)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

    // TODO: Check if the response status code is 200
    if resp.StatusCode == http.StatusOK {
        // TODO: Print the response data
        fmt.Println(string(body))
    }
}
