package main

import (
    "fmt"
    "net/http"
)

// Base URL for the API
const baseUrl = "http://localhost:8000"

func main() {
    // Todo ID to update
    todoID := 3

    // TODO: Modify the PATCH request to a DELETE request and remove patchData
    request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/todos/%d", baseUrl, todoID), nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer response.Body.Close()

    // TODO: Update response handling to support DELETE method and 204 status code
    if response.StatusCode == http.StatusNoContent {
        fmt.Println("Todo deleted successfully!")
    } else {
        // Handle potential errors
        fmt.Println("Error updating the todo")
        fmt.Printf("Status Code: %d\n", response.StatusCode)
    }
}
