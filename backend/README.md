# Simple Dashboard Backend

This project is the backend service for the Simple Dashboard application, built using Go (Golang). It provides APIs to manage and serve data for the dashboard.

## Features

- RESTful API endpoints
- Customer CRUD
- Authentication and authorization

## Prerequisites

- Go 1.20 or later
- Docker
- A running database (e.g., PostgreSQL, MySQL)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/simple-dashboard-backend.git
    cd simple-dashboard-backend
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Configure environment variables in `.env` file.

4. Run the application:
    ```bash
    go run main.go
    ```

## API Documentation

Refer to the [API Documentation](docs/api.md) for details on available endpoints.
