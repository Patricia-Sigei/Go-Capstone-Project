# Getting Started with Go - A Beginner's Guide
## Why Go?
I chose Go because it's designed for simplicity, has built-in concurrency, and is used by companies like Google, Uber, and Dropbox for backend services. I wanted to learn a compiled language that's fast and has a clean syntax.
## What is Go?
Go (Golang) is an open-source programming language created at Google in 2009. It's statically typed, compiled, and designed for building scalable, concurrent applications. It's particularly popular for web servers, cloud services, and CLI tools.
Real-world use: Docker and Kubernetes are written in Go.

## System Requirements
OS: Linux (Ubuntu/Debian) - can also work on Windows, macOS
Go version: 1.21+ (I used 1.21.5)
Editor: VS Code with Go extension (recommended) or any text editor
Shell: Zsh or Bash


## Installation
For Linux (Ubuntu/Debian)
Step 1: Remove Old Version (if exists)

    bashsudo rm -rf /usr/local/go
Step 2: Download and Install Go
   
    wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz

 Extract to /usr/local
 
    sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

Step 3: Add Go to PATH

    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc

 Reload your shell configuration
  
    source ~/.zshrc
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

Reload your shell configuration
      
      source ~/.bashrc

Step 4: Verify Installation

    bashgo version
Expected Output:

    go version go1.21.5 linux/amd64
    
For macOS
bash# Using Homebrew (easiest method)

    brew install go

Verify installation

    go version
    
For Windows

Visit https://go.dev/dl/
Download the Windows installer (.msi file)
Run the installer and follow the prompts
Open Command Prompt and verify:

    cmdgo version

### Errors Encountered During Installation
Error 1: zsh: command not found: go
Problem: After installing Go, the go command wasn't recognized in the terminal.
Root Cause: Go was added to .bashrc but I was using zsh shell (not bash), so the PATH wasn't set correctly.
Solution:
bash# Add Go to zsh PATH
      
      echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc

Reload zsh configuration

    source ~/.zshrc

Verify it works

    go version
How to know which shell you're using:

    bashecho $SHELL
Output: /bin/zsh or /bin/bash

Error 2: shopt and bash_completion warnings
Problem: Warnings appeared when sourcing .bashrc:
/home/patricia/.bashrc:16: command not found: shopt
/usr/share/bash-completion/bash_completion:45: command not found: shopt
Root Cause: Using source ~/.bashrc in a zsh shell (bash-specific commands don't work in zsh).
Solution: Use ~/.zshrc instead of ~/.bashrc for zsh users. These warnings don't affect Go's functionality, but to avoid them, always use the correct config file for your shell.

## Running A Project in Go

Step 1: Create Your Project Directory
bash# Create a directory for learning Go

    mkdir ~/Go-learning  
    cd ~/Go-learning
    
Step 2: Initialize Go Module
Initialize a Go module (like package.json for Node.js)

    go mod init myproject
Output:
go: creating new go.mod: module myproject
This creates a go.mod file that tracks your project's dependencies.

Step 3: Create Your First Program
bash# Create main.go file

    touch main.go

Step 4: Run Your Program

    go run main.go
Output:
Hello, World! This is my first Go program!
I'm learning Go programming! 🚀
