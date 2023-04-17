package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ServerTimeResponse struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/gettime" {
			serverTime := time.Now().Format(time.RFC3339)
			response := ServerTimeResponse{Time: serverTime}
			responseJSON, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Starting server on :8795...")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
