# CRUD App Built with Go, Gin, Gorm, and PostgreSQL

## Description
This is a simple CRUD (Create, Read, Update, Delete) application built with Go, Gin, Gorm, and PostgreSQL. It allows you to manage users by performing basic CRUD operations. The app is designed for simplicity and serves as a great starting point for learning Go web development

## Features
- RESTful API design.
- PostgreSQL as the database.
- GORM for ORM (Object Relational Mapping).
- Docker support for seamless setup.


## Installation
1. Clone the repository
```bash
git clone https://github.com/pvfarooq/go-gin-crud
cd go-gin-crud
```
2. Set Up Environment Variables
```bash
cp env.example .env
```
Update the `.env` file with your PostgreSQL database credentials.

3. Run the App with Docker
```bash
docker-compose up
```
This will:
- Start a PostgreSQL database container.
- Launch the Go server on http://localhost:8080.

## Usage
You can interact with the application using the following API endpoints:

- `GET /users`: Get all users
- `GET /users/:id`: Get a user by ID
- `POST /users`: Create a new user
- `PUT /users/:id`: Update a user by ID
- `DELETE /users/:id`: Delete a user by ID

## Project Structure
```
├ api/
├── controllers/            # Contains the request handlers
│   └── user_controller.go
├── routes/                 # Contains the route definitions
│   └── user_routes.go
│   └── routes.go
├ config/                   # Contains the configuration settings
│   └── config.go
├ database/ 
│   └── migrations/         # Contains the database migrations
│   └── db.go               # Contains the database connection
├ models/                   # Contains the data models
│   └── user.go 
├ repositories/             # Contains the database operations
│   └── user_repository.go
├ services/                 # Contains the business logic
│   └── user_service.go
.env
docker-compose.yml
Dockerfile
env.example
go.mod
go.sum
Makefile                    # Contains the build and run commands
server.go                   # Contains the main function (entry point)
```

## Contributing
Contributions are welcome! If you have suggestions or find issues, feel free to open a pull request or create an issue in the repository.