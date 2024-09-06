## Development Setup

### Prerequisites

- Docker
- Docker Compose

### Dockerizing the Application

To containerize the `auth-service` application, follow these steps:

1. **Create a Dockerfile**

   In the root of your project, create a file named `Dockerfile` with the following content:

   ```dockerfile
   # Use the official Golang image to create a build stage
   FROM golang:1.20 AS builder

   # Set the Current Working Directory inside the container
   WORKDIR /app


   # Copy the source code into the container
   COPY . .

   # Build the Go app
   RUN go build -o auth-service .

   # Use a smaller base image for the final stage
   FROM alpine:latest

   # Set the Current Working Directory inside the container
   WORKDIR /root/

   # Copy the Pre-built binary file from the previous stage
   COPY --from=builder /app/auth-service .

   # Expose port 8080 to the outside world
   EXPOSE 8080

   # Command to run the executable
   CMD ["./auth-service"]
