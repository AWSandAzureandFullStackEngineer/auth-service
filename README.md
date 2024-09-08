# auth-service

Overview

This repository contains a Go application named auth-service using the Fiber web framework. 
This guide will walk you through:

Containerizing the Go application with Docker

Setting up Air for live reloading during development

Configuring Docker Compose to manage both the application and a PostgreSQL database
Connecting the application to the PostgreSQL database

Getting Started

Prerequisites

Before you begin, ensure you have the following installed on your machine:

Docker

Docker Compose

Go (for local development and testing)
1. Containerizing the Go Application
   
1.1. Create a Dockerfile
In the root of your project directory, 
create a Dockerfile with the following content:

```
FROM golang:1.23-alpine

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest


COPY . .

RUN go mod tidy
```

1.2. Create a Docker Compose File
To manage your development environment with Docker Compose, 
create a docker-compose.yml file in the root of your project:

```
version: '3.8'

services:
auth-service:
build: .
env_file:
- .env
ports:
- "8080:8080"
volumes:
- .:/usr/src/app
command: air cmd/auth-service/main.go -b 0.0.0.0

db:
image: postgres:16.3-alpine
environment:
- POSTGRES_USER=${DB_USER}
- POSTGRES_PASSWORD=${DB_PASSWORD}
- POSTGRES_DB=${DB_NAME}
ports:
- "5432:5432"
volumes:
- postgres-db:/var/lib/postgresql/data

volumes:
postgres-db:
```

2. Setting Up Air for Live Reloading
   
2.1. Install Air Locally
To use Air for live reloading, install it with:

go install github.com/air-verse/air@latest

Alternatively, install it via script:

# Install Air to the default GOPATH bin directory
```
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
# Or install it into ./bin/
```
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
```

Verify the installation:

2.2. Create Air Configuration File
Ensure you have an air.toml file in the root of your project with the following configuration:
```
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
args_bin = []
bin = "./tmp/main"
cmd = "go build -o ./tmp/main ./cmd/auth-service"
delay = 1000
exclude_dir = ["assets", "tmp", "vendor", "testdata"]
exclude_file = []
exclude_regex = ["_test.go"]
exclude_unchanged = false
follow_symlink = false
full_bin = ""
include_dir = []
include_ext = ["go", "tpl", "tmpl", "html"]
include_file = []
kill_delay = "0s"
log = "build-errors.log"
poll = false
poll_interval = 0
post_cmd = []
pre_cmd = []
rerun = false
rerun_delay = 500
send_interrupt = false
stop_on_error = false

[color]
app = ""
build = "yellow"
main = "magenta"
runner = "green"
watcher = "cyan"

[log]
main_only = false
time = false

[misc]
clean_on_exit = false

[proxy]
app_port = 0
enabled = false
proxy_port = 0

[screen]
clear_on_rebuild = false
keep_scroll = true
```

3. Running the Application
   
3.1. For Development
To run the application with live reloading, use Docker Compose:

```
docker-compose up --build
```

3.2. For Production
To build and run the application for production:

Build the production image:

```
docker build -t auth-service .
```

Run the container:

```
docker run -p 8080:8080 auth-service
```

4. Connecting to PostgreSQL Database
   This section covers how to connect your Go application to a PostgreSQL database using GORM, 
   with configuration managed through a .env file.

4.1. Create .env File
Create a .env file in the root of your project with the following content:

```
DATABASE_URL=postgres://engineer25:thewordistheword@db:5432/authdb?sslmode=disable
DB_USER=engineer25
DB_PASSWORD=thewordistheword
DB_NAME=authdb
```
4.2. Update Go Application to Use GORM
Ensure that your Go application uses GORM to connect to the PostgreSQL database. 
Here's an example of how to initialize GORM in your Go application:

```
database/database.go

package database

import (
"auth-service/models"
"fmt"
"gorm.io/driver/postgres"
"gorm.io/gorm"
"gorm.io/gorm/logger"
"log"
"os"
)

type Dbinstance struct {
Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
os.Getenv("DB_USER"),
os.Getenv("DB_PASSWORD"),
os.Getenv("DB_NAME"),
)
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
Logger: logger.Default.LogMode(logger.Info),
})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	DB = Dbinstance{Db: db}
}
```

Update your main.go to initialize the database:

```
cmd/auth-service/main.go
```
```
package main

import (
"auth-service/database"
"github.com/gofiber/fiber/v2"
"log"
)

func main() {
database.ConnectDb()
// Initialize a new Fiber app
app := fiber.New()

    setupRoutes(app)

    log.Fatal(app.Listen(":8080"))
}

```

Contribution
Feel free to open issues or submit pull requests if you have any suggestions or improvements!
