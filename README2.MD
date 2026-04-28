Go Toolkit – Code Explanation & Documentation

Overview

This Go program creates a simple web server with three endpoints: /, /api, and /time. It demonstrates
HTTP routing, JSON responses, and dynamic time handling.


Main Function
The main function is the entry point. It registers routes and starts the server.
http.HandleFunc("/", homeHandler)
http.HandleFunc("/api", apiHandler)
http.HandleFunc("/time", timeHandler)
http.ListenAndServe(":8080", nil)
Handlers


Home Handler: Returns plain text response.
fmt.Fprintf(w, "Welcome to Go Web Server n")


API Handler: Returns JSON response.
{
"message": "Hello from Go API!",
"status": "success"
}
Time Handler: Returns current server time.
{
"current_time": "Mon, 28 Apr 2026 10:00:00 UTC"
}


JSON Encoding
The json encoder converts Go maps into JSON format and sends them as HTTP responses.
json.NewEncoder(w).Encode(response)

Key Concepts
Handler Function: Responds to HTTP requests
Routing: Mapping URLs to functions
JSON: Data format for APIs


HTTP Server: Serves web requests
How to Run
go run main.go
Visit http://localhost:8080