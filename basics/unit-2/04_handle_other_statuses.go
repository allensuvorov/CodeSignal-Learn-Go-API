package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

const baseURL = "http://localhost:8000"

func main() {
    // TODO: Perform a GET request to the /todos endpoint using http.Get()
    resp, err := http.Get(baseURL + "/todos")
    if err != nil {
        log.Fatalf("Failed to get todos: %v", err)
    }
    
    // TODO: Handle status code 200 to parse and print the to-do items if the request was successful
    if resp.StatusCode == http.StatusOK {
        fmt.Println("Todos retrieved successfully:")
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Failed to read response body: %v", err)
        }
        fmt.Println(string(body))
    // TODO: Generally handle unexpected status codes, printing the status code and error details from the response
    } else {
        fmt.Printf("Unexpected Status Code: %d \n", resp.StatusCode)
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Failed to read error reponse: %v", err)
        }
        fmt.Printf("Error Details: %s\n", body)
    }
}
