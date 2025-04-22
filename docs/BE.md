# Backend Setup Guide

This guide explains how to set up the backend locally and using Docker Compose.

## Prerequisites
- **Go** (version 1.24.1 or higher)
- **Go Fiber** 
- **PostgreSQL** (version 15 or latest)
- **Gorm** (version 1.25 or latest)
- **Minio** (latest)
- **Docker** and **Docker Compose**

---

## Local Setup

1. **Clone the Repository**:
    ```bash
    git clone git@github.com:yamanmnur/super-bank-assesment.git
    cd super-bank-assesment
    ```

2. **Install Dependencies**:
    ```bash
    go mod tidy
    # or
    go mod download
    ```

3. **Set Up Environment Variables**:
    - Copy `.env.example` to `.env`:
      ```bash
      cp .env.example .env
      ```
    - Update the `.env` file with your local configuration.
    - for url database use docker container name database if you want to run backend using docker
    ```
    example : postgres://postgres:password@simplecustomerdb:5432/customer_db

    ```
    - if you want to run server without docker use this
    ```
    example : postgres://postgres:password@localhost:5432/<database_name>

    ```

4. **Database Migrations & Seeder** (if applicable):
    
    Migration and seeder will run automatically when the backend server starts.

    ```bash
    go run main.go
    ```

5. **Start the Development Server**:
    ```bash
    go run main.go
    ```

6. **Access the Application**:
    - Open Postman, use the existing collection in this repo, and change the **server_url** variable to http://localhost:3001.

---

## Docker Compose Setup

1. **Clone the Repository**:
    ```bash
    git clone git@github.com:yamanmnur/super-bank-assesment.git
    cd super-bank-assesment
    ```

2. **Set Up Environment Variables**:
    - Copy `.env.example` to `.env`:
      ```bash
      cd backend
      cp .env.example .env
      ```
    - Update the `.env` file with your Docker configuration.

3. **Start Services**:
    ```bash
    docker-compose up --build -d
    ```

4. **Access the Application**:
    - Open Postman, use the existing collection in this repo, and change the server_url variable to http://localhost:3001.

---
# Backend Structure
## Backend Structure

The backend is organized using a layered architecture with the **Service Repository Pattern** and **Manual Dependency Injection**. Below is an explanation of the structure:

### Folder Structure
```
backend/
├── internal/
│   ├── dto/            # DTO
│   ├── repositories/   # Data access layer (e.g., database interactions)
│   ├── service/        # Business logic layer
│   ├── injectors/      # Manual Depedency Injector
│   ├── handler/        # HTTP handlers (controllers)
│   ├── models/         # Models for Tables
│   ├── middlewares/    # Custom middlewares
│   ├── utils/          # Utility functions
├── seeder/             # Database Seeder
├── routes/             # API Route Definition
├── tests/              # Unit Test
├── pkg/                # Init DB Migration, Gorm Client, Base  Response and others 
├── main.go             # Main application file
```

### Explanation of Layers

1. **DTO Layer**:
    - Contains the interfaces for response, requests and data.
    - Example: `UserData`, `CustomerData`, `CustomerDetailData`.

2. **Repository Layer**:
    - Responsible for interacting with the database.
    - Implements the interfaces defined in the domain layer.
    - Example: `user_repository`, `auth_service`.

3. **Service Layer**:
    - Contains the business logic of the application.
    - Calls the repository layer to fetch or persist data.
    - Example: `customer_service`, `auth_service`.

4. **Handler Layer**:
    - Handles HTTP requests and responses.
    - Calls the service layer to process the request.
    - Example: `auth_controller`, `customer_controller`.

5. **Middleware Layer**:
    - Contains custom middleware for request processing.
    - Example: Authentication, Error Handler.

    - Using Error Handler if each service or controller return error it will automatically mapping to determined response structure
    ```json
    {
        "meta_data" : {
            "status" : "error",
            "message" : "server error",
            "code" : "500"
        }
    }
    ```

6. **Utils Layer**:
    - Contains helper functions or utilities used across the application.
    - Example: Generate Random Number, Format Currency.

### Dependency Injection
- The application uses **Manual Dependency Injection** to wire up the dependencies between layers.
- Dependencies are passed explicitly through constructors.
- Example:
     ```go
     // In main_injector.go
     type Container struct {
        CustomerRepository      *repositories.CustomerRepository
        CustomerService         *services.CustomerService
        CustomerController      *services.CustomerController
     }
     ...
     InjectorCustomerDI(dbHandler, container)
     ```

This approach ensures loose coupling and makes the application easier to test and maintain.

## Notes
- Replace `<repository-url>` and `<PORT>` with the actual values for your project.
- Ensure Docker is running before using Docker Compose.
