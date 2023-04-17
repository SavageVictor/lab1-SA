package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// ServerTimeResponse is a struct to represent the JSON response containing the server time
type ServerTimeResponse struct {
    Time string json:"time"
}

func main() {
    // Set up a route handler for incoming requests
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Check if the request is a GET request for the "/gettime" path
        if r.Method == http.MethodGet && r.URL.Path == "/gettime" {
            // Get the current server time in RFC3339 format
            serverTime := time.Now().Format(time.RFC3339)
            // Create a response object with the server time
            response := ServerTimeResponse{Time: serverTime}
            // Marshal the response object into a JSON string
            responseJSON, err := json.Marshal(response)
            if err != nil {
                // If there's an error generating the JSON string, return an error response
                http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
                return
            }
            // Set the response content type to JSON
            w.Header().Set("Content-Type", "application/json")
            // Write the JSON response
            w.Write(responseJSON)
            return
        }
        // If the request isn't for "/gettime", serve the index.html file
        http.ServeFile(w, r, "index.html")
    })

    // Start the server on port 8795
    fmt.Println("Starting server on :8795...")
    if err := http.ListenAndServe(":8795", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}