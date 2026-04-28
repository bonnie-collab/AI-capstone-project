# Go Beginner Toolkit – Simple Web Server

# Project Overview
This project is a beginner-friendly toolkit for learning Go by building a simple web server.

## ▶️ Run
go run main.go

## 🌐 Output
Visit http://localhost:8080


# 🚀 Go Beginner Toolkit – Simple Web Server

## 📌 Project Overview
This project is a beginner-friendly toolkit for learning the Go programming language by building a simple web server.

The server responds with:
"Hello, World from Go!"

---

##  Objective
- Learn Go basics
- Build and run a simple web server
- Understand how HTTP handling works in Go

---

##  Requirements
- Go (Golang) installed
- VS Code or any code editor
- Terminal / Command Prompt

---

##  Installation & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/go-toolkit.git
cd go-toolkit



## 🔌 API Endpoints

#GET /
Returns a welcome message

#GET /api
Returns JSON response:
{
  "message": "Hello from Go API!",
  "status": "success"
}

### GET /time
Returns current server time in JSON


the code documentation towards to what it does and carries out explaination and documentation

Great—you’ve got a solid piece of code here. I’ll break it down **line by line**, then give you a **clean beginner-friendly documentation section** you can include in your project.

---

# 🧠 1. Line-by-Line Explanation

```go
package main
```

* Defines the package name.
* `main` is special—it tells Go this is an **executable program** (not a library).

---

```go
import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)
```

These are built-in Go packages:

* `encoding/json` → used to convert data into JSON format
* `fmt` → used for printing text (like console output or responses)
* `net/http` → used to create web servers and handle HTTP requests
* `time` → used to work with date and time

---

## Root Handler

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
```

* Defines a function to handle requests to `/`
* `w` → where you send the response (output to browser)
* `r` → contains request data (URL, headers, etc.)

---

```go
    fmt.Fprintf(w, "Welcome to Go Web Server 🚀")
```

* Sends text back to the browser
* `Fprintf` writes formatted output to `w` (the response)

---

```go
}
```

* Ends the function

---

## 🔌 API Handler (JSON Response)

```go
func apiHandler(w http.ResponseWriter, r *http.Request) {
```

* Handles requests to `/api`

---

```go
    w.Header().Set("Content-Type", "application/json")
```

* Tells the browser: “This response is JSON”

---

```go
    response := map[string]string{
```

* Creates a **map (dictionary)** to hold data

---

```go
        "message": "Hello from Go API!",
        "status":  "success",
```

* Key-value pairs inside the map

---

```go
    }
```

* Ends the map

---

```go
    json.NewEncoder(w).Encode(response)
```

* Converts the map into JSON
* Sends it to the browser

---

```go
}
```

* Ends the function

---

## ⏰ Time Handler

```go
func timeHandler(w http.ResponseWriter, r *http.Request) {
```

* Handles requests to `/time`

---

```go
    w.Header().Set("Content-Type", "application/json")
```

* Again, response will be JSON

---

```go
    response := map[string]string{
```

* Create a map

---

```go
        "current_time": time.Now().Format(time.RFC1123),
```

* Gets current time
* Formats it into readable format like:
  `Mon, 28 Apr 2026 10:00:00 UTC`

---

```go
    }
```

* Ends the map

---

```go
    json.NewEncoder(w).Encode(response)
```

* Converts to JSON and sends to browser

---

```go
}
```

* Ends the function

---

## 🚀 Main Function (Program Entry Point)

```go
func main() {
```

* The starting point of the program

---

```go
    http.HandleFunc("/", homeHandler)
```

* Maps URL `/` → `homeHandler`

---

```go
    http.HandleFunc("/api", apiHandler)
```

* Maps `/api` → API handler

---

```go
    http.HandleFunc("/time", timeHandler)
```

* Maps `/time` → time handler

---

```go
    fmt.Println("Server running at http://localhost:8080")
```

* Prints message in terminal

---

```go
    http.ListenAndServe(":8080", nil)
```

* Starts the web server
* Runs on port **8080**
* `nil` → uses default settings

---

```go
}
```

* Ends program

---

# 📘 2. Beginner-Friendly Documentation (Add to Your Project)

You can include this as a section in your toolkit:

---

## 📖 Code Explanation Guide

### 🔹 Overview

This Go program creates a simple web server with three endpoints:

* `/` → returns plain text
* `/api` → returns JSON data
* `/time` → returns current server time in JSON

---

## 🔹 How It Works

### 1. The `main` Function

This is the entry point of the application. It:

* Registers URL routes
* Starts the server

---

### 2. Routing System

```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/api", apiHandler)
http.HandleFunc("/time", timeHandler)
```

Each route is connected to a function:

* `/` → `homeHandler`
* `/api` → `apiHandler`
* `/time` → `timeHandler`

---

### 3. Handlers Explained

#### 🏠 Home Handler

Returns plain text:

```go
fmt.Fprintf(w, "Welcome to Go Web Server 🚀")
```

---

#### 🔌 API Handler

Returns JSON:

```json
{
  "message": "Hello from Go API!",
  "status": "success"
}
```

---

#### ⏰ Time Handler

Returns current time:

```json
{
  "current_time": "Mon, 28 Apr 2026 10:00:00 UTC"
}
```

---

### 4. JSON Encoding

```go
json.NewEncoder(w).Encode(response)
```

* Converts Go data (map) → JSON
* Sends it as HTTP response

---

### 5. Starting the Server

```go
http.ListenAndServe(":8080", nil)
```

* Starts server on port 8080
* Waits for incoming requests

---

## 🔹 Key Concepts for Beginners

| Concept          | Explanation                             |
| ---------------- | --------------------------------------- |
| Handler Function | Function that responds to a web request |
| Route            | URL path (e.g., `/api`)                 |
| JSON             | Format used to send data                |
| Map              | Key-value data structure                |
| HTTP Server      | Program that serves web requests        |

---

## 🔹 How to Test

Run:

```bash
go run main.go
```

Then open:

* [http://localhost:8080](http://localhost:8080)
* [http://localhost:8080/api](http://localhost:8080/api)
* [http://localhost:8080/time](http://localhost:8080/time)

---
This project demonstrates fundamental backend concepts including HTTP routing, JSON serialization, and dynamic data handling using Go.”

---# AI-capstone-project
