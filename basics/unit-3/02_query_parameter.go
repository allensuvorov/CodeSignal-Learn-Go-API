package main

import (
    "fmt"
    "log"
    "net/http"
    "io"
)

const baseURL = "http://localhost:8000"

func main() {
    // TODO: Include a query parameter to filter for not done items
    
    url := fmt.Sprintf("%s/todos?done=false", baseURL)
    
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching todos: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        data, _ := io.ReadAll(resp.Body)
        fmt.Printf("Todos: %s\n", data)
    } else {
        log.Printf("Failed to fetch todos. Status Code: %d", resp.StatusCode)
    }
}
