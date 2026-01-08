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

    // TODO: Define a json string with the updated todo data, including "title", "description", and "done"
    updatedTodo := `{
        "title": "Master method Put of HTTP requests",
        "done": false,
        "description": "Put method HTTP request update the a resource completly"
    }`

    // TODO: Send a PUT request to update the todo with the specified ID using the updated data
    req, err := http.NewRequest(http.MethodPut, baseUrl + "/todos/" + fmt.Sprint(todoID), bytes.NewBufferString(updatedTodo))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")
    
    client := new(http.Client)
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    // TODO: Check if the request was successful (status code 200) and print a success message along with the updated todo
    if resp.StatusCode == http.StatusOK {
        fmt.Println("Todo updated successfully with PUT!")
        fmt.Println(string(body))
    } else {
        fmt.Println("Error updating the todo with PUT")
        fmt.Printf("Status Code: %d\n", resp.StatusCode)
        fmt.Println("Error Details:", string(body))
    }
    // TODO: Handle potential errors by printing an error message, status code, and error details if the update fails
    
}
