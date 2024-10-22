# ToDo Service

This is a simple ToDo service built using Go, gRPC, Gin, and MongoDB. It allows you to manage your ToDo items through a gRPC
API and provides a HTTP gateway for easy access.

## Features

- Create a new ToDo item
- Retrieve a ToDo item by ID
- List all ToDo items
- Update a ToDo item
- Delete a ToDo item

## Technologies Used

- Go
- gRPC
- Gin
- MongoDB
- Swagger for API documentation

## Getting Started

### Prerequisites

- Go installed on your machine
- MongoDB installed and running

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/goaziz/todo.git
   cd todo
   ```
2. Install the required Go packages:

   ```bash
   go mod tidy
   ```
3. Start the MongoDB server:

   ```bash
    mongod
    ```
4. Run the server:

   ```bash
   go run cmd/server/main.go
   ```