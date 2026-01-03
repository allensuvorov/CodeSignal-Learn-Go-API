package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
)

const baseURL = "http://localhost:8000"

func main() {
    // TODO: Change the parameters to filter todos that are done and start with 'w'
    doneStatus := "true"
    titlePrefix := "w"
    url := fmt.Sprintf("%s/todos?done=%s&title=%s", baseURL, doneStatus, titlePrefix)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching todos with query parameters: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        data, _ := io.ReadAll(resp.Body)
        fmt.Printf("Filtered Todos: %s\n", data)
    } else {
        log.Printf("Failed to fetch todos. Status Code: %d", resp.StatusCode)
    }
}
