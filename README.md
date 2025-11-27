# Getting Started with Go - A Beginner's Guide

## Why Go?
I chose Go because it's designed for simplicity, has built-in concurrency, and is used by companies like Google, Uber, and Dropbox for backend services. I wanted to learn a compiled language that's fast and has a clean syntax.

## What is Go?
Go (Golang) is an open-source programming language created at Google in 2009. It's statically typed, compiled, and designed for building scalable, concurrent applications. It's particularly popular for web servers, cloud services, and CLI tools.

**Real-world use:** Docker and Kubernetes are written in Go.

## System Requirements
- **OS:** Linux (Ubuntu/Debian) - can also work on Windows, macOS
- **Go version:** 1.21+ (I used 1.21.5)
- **Editor:** VS Code with Go extension (recommended) or any text editor
- **Shell:** Zsh or Bash

---

## Installation

### For Linux (Ubuntu/Debian)

**Step 1: Remove Old Version (if exists)**
```bash
sudo rm -rf /usr/local/go
```

**Step 2: Download and Install Go**
```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
# Extract to /usr/local
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
```

**Step 3: Add Go to PATH**

For Zsh users:
```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
# Reload your shell configuration
source ~/.zshrc
```

For Bash users:
```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
# Reload your shell configuration
source ~/.bashrc
```

**Step 4: Verify Installation**
```bash
go version
```
Expected Output:
```
go version go1.21.5 linux/amd64
```

### For macOS
```bash
# Using Homebrew (easiest method)
brew install go

# Verify installation
go version
```

### For Windows
1. Visit https://go.dev/dl/
2. Download the Windows installer (.msi file)
3. Run the installer and follow the prompts
4. Open Command Prompt and verify:
```cmd
go version
```

---

## Common Installation Errors

### Error 1: `zsh: command not found: go`

**Problem:** After installing Go, the `go` command wasn't recognized in the terminal.

**Root Cause:** Go was added to `.bashrc` but I was using zsh shell (not bash), so the PATH wasn't set correctly.

**Solution:**
```bash
# Add Go to zsh PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc

# Reload zsh configuration
source ~/.zshrc

# Verify it works
go version
```

**How to know which shell you're using:**
```bash
echo $SHELL
```
Output: `/bin/zsh` or `/bin/bash`

### Error 2: shopt and bash_completion warnings

**Problem:** Warnings appeared when sourcing `.bashrc`:
```
/home/patricia/.bashrc:16: command not found: shopt
/usr/share/bash-completion/bash_completion:45: command not found: shopt
```

**Root Cause:** Using `source ~/.bashrc` in a zsh shell (bash-specific commands don't work in zsh).

**Solution:** Use `~/.zshrc` instead of `~/.bashrc` for zsh users. These warnings don't affect Go's functionality, but to avoid them, always use the correct config file for your shell.

---

## Creating and Running Your First Go Project

### Step 1: Create Your Project Directory
```bash
# Create a directory for learning Go
mkdir ~/Go-learning  
cd ~/Go-learning
```

### Step 2: Initialize Go Module
```bash
# Initialize a Go module (like package.json for Node.js)
go mod init myproject
```
Output:
```
go: creating new go.mod: module myproject
```
This creates a `go.mod` file that tracks your project's dependencies.

### Step 3: Create Your First Program
```bash
# Create main.go file
touch main.go
```

Open `main.go` in your editor and add this simple Hello World program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World! This is my first Go program!")
    fmt.Println("I'm learning Go programming! 🚀")
}
```

### Step 4: Run Your Program
```bash
go run main.go
```

Output:
```
Hello, World! This is my first Go program!
I'm learning Go programming! 🚀
```

---

## Building an HTTP Server 

This shows how Go programs can evolve from simple console apps to full web servers:

```go
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
```

### Running the HTTP Server

```bash
go run main.go
```

You should see:
```
Hello, World! This is my first Go program!
I'm learning Go programming! 🚀

🚀 Server starting on http://localhost:8080
📍 Try these endpoints:
   - http://localhost:8080
   - http://localhost:8080/greet/YourName
   - http://localhost:8080/about

Press Ctrl+C to stop the server
```

### Testing Your Endpoints

Open your browser or use `curl` to test:

```bash
# Test home endpoint
curl http://localhost:8080

# Test greet endpoint with your name
curl http://localhost:8080/greet/Patricia

# Test about endpoint
curl http://localhost:8080/about
```

---

## Understanding the Code

### Package and Imports
```go
package main
```
Every Go program starts with a package declaration. `main` is special - it's the entry point for executable programs.

```go
import (
    "fmt"      // For formatted I/O (printing to console/web)
    "net/http" // For HTTP server functionality
    "time"     // For working with time and dates
)
```

### Handler Functions
Handler functions process HTTP requests and write responses:

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // w is where we write our response
    // r contains information about the incoming request
    fmt.Fprintf(w, "Welcome to my Go server! 🚀\n")
}
```

- `w http.ResponseWriter` - writes data back to the client
- `r *http.Request` - contains request details (URL, headers, method, etc.)

### Registering Routes
```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/greet/", greetHandler)
http.HandleFunc("/about", aboutHandler)
```
This connects URL paths to handler functions. When someone visits `/greet/`, Go calls `greetHandler`.

### Starting the Server
```go
err := http.ListenAndServe(":8080", nil)
```
This starts the server on port 8080. The function returns an error if something goes wrong (like port already in use).

### Error Handling
```go
if err != nil {
    fmt.Println("❌ Error starting server:", err)
}
```
Go uses explicit error checking instead of try-catch blocks. This makes errors visible and forces you to handle them.

---

## Key Go Concepts Learned

### 1. Strong Static Typing
Go is statically typed, meaning variable types are checked at compile time:
```go
var name string = "Patricia"  // Explicit type
age := 25                     // Type inferred (int)
```

### 2. Error Handling Philosophy
Go returns errors as values, not exceptions:
```go
result, err := someFunction()
if err != nil {
    // Handle the error
    fmt.Println("Error:", err)
    return
}
// Use result
```

### 3. Simplicity and Readability
Go intentionally has minimal syntax. There's usually one obvious way to do things, making code easier to read and maintain.

### 4. Built-in Concurrency (Goroutines)
Go makes concurrent programming easy with goroutines - lightweight threads managed by the Go runtime:
```go
go functionName()  // Runs function concurrently
```

### 5. Fast Compilation
Go compiles directly to machine code, creating single binary executables that run fast and are easy to deploy.

### 6. Powerful Standard Library
Go's standard library includes HTTP servers, JSON handling, file I/O, cryptography, and much more - no external frameworks needed for basic tasks.

---

## Helpful AI Prompts for Learning Go

As you continue learning, here are some useful prompts to deepen your understanding:

### Installation & Setup
- "Give me step-by-step installation instructions for Go on Linux Ubuntu"
- "Why is go command not found after installation in zsh?"
- "How do I set up VS Code for Go development?"

### Basic Concepts
- "Write a complete Hello World program in Go with detailed comments explaining each part"
- "Explain Go's basic data types and how variable declaration works"
- "What's the difference between var declaration and short declaration (:=) in Go?"
- "Explain what go mod init does and why I need it for Go projects"

### Building Programs
- "How do I extend my Hello World program to also run an HTTP server with multiple routes?"
- "Show me how to read JSON from an HTTP request body in Go"
- "How do I add query parameters to my HTTP endpoints?"
- "What's the best way to organize a Go web server project with multiple files?"

### Error Handling
- "How does error handling work in Go? Why doesn't it have try-catch?"
- "Show me examples of proper error handling in Go HTTP handlers"
- "When should I return errors in Go?"

### Debugging & Troubleshooting
- "My Go server returns 404 for all routes, what could be wrong?"
- "How do I debug Go programs? What tools are available?"
- "Why is my goroutine not executing?"

### Best Practices
- "What are Go naming conventions for variables and functions?"
- "How should I structure error messages in Go?"

---

## Useful Go Commands

```bash
# Run a Go program (compiles and runs)
go run main.go

# Create something like package.json
go mod init myproject  

```

---

## Next Steps in Your Go Journey

1. **Learn about structs and methods** - Go's way of creating custom types
2. **Explore JSON handling** - Most web APIs use JSON for data exchange
3. **Understand interfaces** - Go's approach to polymorphism
4. **Practice with goroutines and channels** - Go's concurrency primitives
5. **Build a REST API** - Create, read, update, delete operations
6. **Connect to a database** - Learn to work with PostgreSQL or MongoDB
7. **Write tests** - Go has built-in testing support
8. **Deploy your application** - Learn about containerization with Docker

---

## Resources

- **Official Go Website:** https://go.dev/
- **Go Tour:** https://go.dev/tour/ (Interactive learning)
- **Go by Example:** https://gobyexample.com/ (Code examples)
- **Effective Go:** https://go.dev/doc/effective_go (Best practices)
- **Go Playground:** https://go.dev/play/ (Try Go in your browser)

---

## My Learning Journal Summary

### What I've Accomplished
✅ Successfully installed Go on Linux  
✅ Fixed PATH configuration issues with zsh  
✅ Created and ran my first Hello World program  
✅ Built a working HTTP server with multiple routes  
✅ Understood Go's basic syntax and project structure  
✅ Learned about error handling in Go  
✅ Gained exposure to Go's concurrency concepts  

### Key Takeaways
- Go's simplicity makes it approachable for beginners
- The standard library is incredibly powerful
- Error handling feels explicit but predictable
- Building web servers is surprisingly straightforward
- The tooling (go fmt, go mod) makes development smooth

### What's Next
- Dive deeper into structs and custom types
- Learn about Go interfaces
- Practice with goroutines and channels
- Build a complete REST API with database integration

---

## AI Prompt Journal - My Learning Sessions

### Prompt 1: Installation Guide
**Prompt Used:** "Give me step-by-step installation instructions for Go on Linux Ubuntu"

**AI Response Summary:** The AI provided comprehensive installation steps including downloading the Go tarball, extracting it to /usr/local, and adding Go to the PATH. It also included verification steps using go version.

**Result:** Successfully installed Go 1.21.5 on my Linux system.

**Evaluation:** ⭐⭐⭐⭐⭐ Extremely helpful. The step-by-step approach made installation straightforward even for a complete beginner.

---

### Prompt 2: Troubleshooting PATH Issue
**Prompt Used:** "Why is go command not found after installation in zsh?"

**AI Response Summary:** The AI explained that the PATH was added to .bashrc but I was using zsh shell, which reads from .zshrc. It provided the exact commands to add Go to the zsh PATH and reload the configuration.

**Result:** Fixed the command not found: go error by adding the PATH to ~/.zshrc instead of ~/.bashrc.

**Evaluation:** ⭐⭐⭐⭐⭐ Critical help! Without this, I would have been stuck. The AI quickly identified the shell mismatch issue and provided the exact solution.

---

### Prompt 3: Hello World Program
**Prompt Used:** "Write a complete Hello World program in Go with detailed comments explaining each part"

**AI Response Summary:** The AI provided a simple Hello World program with inline comments explaining package main, import "fmt", and the main() function. It also explained why each component is necessary.

**Result:** Created my first runnable Go program and understood the basic structure of Go programs.

**Evaluation:** ⭐⭐⭐⭐ Very helpful. The comments made it easy to understand what each line does, though I needed follow-up prompts for deeper understanding.

---

### Prompt 4: Understanding Data Types
**Prompt Used:** "Explain Go's basic data types and how variable declaration works"

**AI Response Summary:** The AI explained Go's type system including int, float64, string, bool, and the difference between var declaration and short declaration (:=). It provided examples of both explicit and implicit typing.

**Result:** Learned about type safety and how Go handles variables differently from dynamically typed languages.

**Evaluation:** ⭐⭐⭐⭐ Good foundational knowledge. Would have been better with more practical examples.

---

### Prompt 5: Building HTTP Server
**Prompt Used:** "How do I extend my Hello World program to also run an HTTP server with multiple routes? I want it to first print Hello World, then start a web server"

**AI Response Summary:** The AI provided code that combines console output with an HTTP server in a single main() function. It showed how to create handler functions, register routes, and start the server after printing to console. The code demonstrated progression from simple to complex in one file.

**Result:** Successfully extended my Hello World program to include a working REST API with three endpoints, all in one file showing clear progression.

**Evaluation:** ⭐⭐⭐⭐⭐ Excellent! This approach shows how Go programs can evolve and grow. The single-file approach made it easy to see the progression from basic to advanced.

---

### Prompt 6: Error Handling Concept
**Prompt Used:** "How does error handling work in Go? Why doesn't it have try-catch?"

**AI Response Summary:** The AI explained Go's philosophy of explicit error handling, showing how functions return error values that must be checked. It contrasted this with exception-based systems and showed the pattern of if err != nil checks.

**Result:** Understood Go's explicit error handling approach and why it's considered more predictable than exceptions.

**Evaluation:** ⭐⭐⭐⭐ Good conceptual explanation. I'll need more practice to see the benefits in real projects.

---

### Prompt 7: Concurrency Introduction
**Prompt Used:** "What are goroutines and how do they make Go good for concurrent programming?"

**AI Response Summary:** The AI introduced goroutines as lightweight threads managed by the Go runtime. It explained the basic syntax (go functionName()) and mentioned channels for communication between goroutines.

**Result:** Got a high-level understanding of Go's concurrency model. This is something I'll explore deeper later.

**Evaluation:** ⭐⭐⭐ Useful introduction but too advanced for my current level. I need to master the basics first.

---

### Prompt 8: Project Structure
**Prompt Used:** "Explain what go mod init does and why I need it for Go projects"

**AI Response Summary:** The AI explained that go mod init creates a go.mod file that tracks the module name and dependencies, similar to package.json in Node.js. It's required for managing packages in modern Go projects.

**Result:** Understood why every project needs module initialization and how Go manages dependencies.

**Evaluation:** ⭐⭐⭐⭐ Clear and concise. Helped me understand the Go module system basics.

---

## Key Learnings from AI Sessions

- **Simple Syntax:** Go's syntax is minimal and easy to read
- **Fast Compilation:** Programs compile quickly into single executables
- **Strong Typing:** Go catches many errors at compile time
- **Explicit Error Handling:** No exceptions - errors are returned values
- **Powerful Standard Library:** Built-in HTTP server, JSON handling, and more
- **Great Tooling:** go fmt, go build, go mod make development smooth

---
