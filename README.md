# auth-service

**Acceptance Criteria:**
- The Go module is initialized with the name `auth-service`.
- The basic project structure is in place.
- The `cmd/main.go` file runs and prints "Hello World" to the console.
- `.gitignore` and `README.md` files are created and contain appropriate content.

**Additional Information:**
- Ensure that the project adheres to Go's best practices for project structure.
- Document any setup or configuration steps clearly in the README.md file.

# auth-service

## Overview

This repository contains a Go application called `auth-service`. This guide will walk you through containerizing the application using Docker, setting up Air for live reloading, and configuring everything with Docker Compose.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (for local development and testing)

### 1. Containerizing the Go Application

To containerize the `auth-service` application, follow these steps:

#### 1.1. Create a Dockerfile for Production

In the root of your project directory, create a file named `Dockerfile` with the following content:

```Dockerfile
# Use the official Golang image as a parent image
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM debian:bullseye-slim

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]
```

1.2. Create a Dockerfile for Development
For development with live reloading, create a Dockerfile.dev in the root of your project:

# Use the official Golang image as a parent image
FROM golang:1.20 AS dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Start Air
CMD ["air"]

2. Set Up Air for Live Reloading
  
2.1. Install Air Locally To use Air for live reloading, you need to install it:

go install github.com/cosmtrek/air@latest

or by sh

# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s

air -v

2.2. Create Air Configuration File

On the command line type in the command air init to initialize air.toml file

3. Configure Docker Compose To manage your development environment with Docker Compose, 
create a docker-compose.yml file in the root of your project:

"services:
auth-service:
build:
context: .
dockerfile: Dockerfile.dev
volumes:
- .:/app
ports:
- "8080:8080"
command: air"

4. Running the Application
   4.1. For Development
   To run the application with live reloading, use Docker Compose:

"docker-compose up --build"

4.2. For Production
To build and run the application for production, follow these steps:

Build the production image:

"docker build -t auth-service ."

Run the container:

"docker run -p 8080:8080 auth-service"


Contribution

Feel free to open issues or submit pull requests if you have any suggestions or improvements!



