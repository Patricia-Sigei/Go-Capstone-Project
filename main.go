package main

import (
    "fmt"
    "net/http"
    "time"
)

// Handler for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my Go server! 🚀\n")
    fmt.Fprintf(w, "Current time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// Handler for the greet route
func greetHandler(w http.ResponseWriter, r *http.Request) {
    // Extract name from URL path
    name := r.URL.Path[len("/greet/"):]
    if name == "" {
        fmt.Fprintf(w, "Hello, stranger! 👋\n")
    } else {
        fmt.Fprintf(w, "Hello, %s! 👋\n", name)
    }
}

// Handler for the about route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a simple Go web server\n")
    fmt.Fprintf(w, "Built by a beginner learning Go!\n")
}

func main() {
    // First: Print Hello World to console
    fmt.Println("Hello, World! This is my first Go program!")
    fmt.Println("I'm learning Go programming! 🚀")
    fmt.Println()
    
    // Second: Start the HTTP server
    // Register route handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/greet/", greetHandler)
    http.HandleFunc("/about", aboutHandler)
    
    // Display startup information
    fmt.Println("🚀 Server starting on http://localhost:8080")
    fmt.Println("📍 Try these endpoints:")
    fmt.Println("   - http://localhost:8080")
    fmt.Println("   - http://localhost:8080/greet/YourName")
    fmt.Println("   - http://localhost:8080/about")
    fmt.Println("\nPress Ctrl+C to stop the server")
    
    // Start the server
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("❌ Error starting server:", err)
    }
}