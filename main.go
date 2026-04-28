package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// Root handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Go Web Server 🚀")
}

// API handler (returns JSON)
func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    response := map[string]string{
        "message": "Hello from Go API!",
        "status":  "success",
    }

    json.NewEncoder(w).Encode(response)
}

// Time handler (dynamic JSON)
func timeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    response := map[string]string{
        "current_time": time.Now().Format(time.RFC1123),
    }

    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/api", apiHandler)
    http.HandleFunc("/time", timeHandler)

    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}