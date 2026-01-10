package main

import (
    "bytes"
    "fmt"
    "io"
    "net/http"
)

func main() {
    // Base URL for the API
    const baseUrl = "http://localhost:8000"

    // Todo ID to update
    todoID := 1

    // TODO: Modify the data to send only the updated 'done' status, remove 'title' and 'description'
    updatedTodo := `{
        "done": true
    }`

    // TODO: Change the request method from PUT to PATCH
    request, err := http.NewRequest(http.MethodPatch, baseUrl+"/todos/"+fmt.Sprint(todoID), bytes.NewBufferString(updatedTodo))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    request.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer response.Body.Close()

    // Check if the request was successful
    if response.StatusCode == http.StatusOK {
        body, _ := io.ReadAll(response.Body)
        fmt.Println("Todo updated successfully!")
        fmt.Println(string(body))
    } else {
        // Handle potential errors
        fmt.Println("Error updating the todo")
        fmt.Printf("Status Code: %d\n", response.StatusCode)
        body, _ := io.ReadAll(response.Body)
        fmt.Println("Error Details:", string(body))
    }
}
