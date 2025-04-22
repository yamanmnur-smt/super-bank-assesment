
# Super Bank Assesment

Simple Dashboard using Go & Next Js for Super Bank Assesment

## Specification
### Backend
- **Go** (version 1.24.1 or higher)
- **Go Fiber** 
- **PostgreSQL** (version 15 or latest)
- **Gorm** (version 1.25 or latest)
- **Minio** (latest)
- **Docker** and **Docker Compose**

### Frontend
- **Node** (version 22 or higher)
- **Next Js** (version 15)
- **Tailwind CSS** 
- **Zustand** 
- **D3** (For Chart) 
- **Docker** and **Docker Compose**

## System Design

![Architecture](/docs/images/architecture_use_docker.png)

By default, this project uses Docker Compose to manage the containers (backend, frontend, MinIO, and PostgreSQL). Each container shares the same network.

![Basic Communication](/docs/images/communication.png)

In this project, the frontend (Next.js client) will first fetch data from the Next.js API server, which will then fetch data from the backend server.

if there are any routes from backend server that need jwt token, next server will use token from cookies

![How JWT Token Stored in Frontend](/docs/images/how_token_stored.png)

After the client fetches the `api/login` endpoint from Next.js, the Next.js server will call the login API from the backend to retrieve the JWT token and user data. If the HTTP request for login is successful, the token will be stored in the cookies on the Next.js server.


## Install & Running - Using Docker Compose

```bash
  docker-compose up --build -d
```
### Local address Frontend
```bash
  http://localhost:3000
```

### Local address Backend
```bash
  http://localhost:3001
```

### Credential Admin CMS
```bash
  username : yaman
  password : password
```
### Credential Postgre SQL
use this credential if use Postgre from docker-compose
```bash
  host     : localhost
  port     : 5432
  username : postgres
  password : password
  database : customer_db
```
### Credential Minio
use this credential if use minio from docker-compose
    
minio running on port 9001 for ui and 9000 for api

http://localhost:9001

```bash
  username : admin
  password : password123
```

## Installation - frontend

1. Update the .env
2. Install for Frontend

```bash
  cd frontend
  npm install
```

## Installation - backend
1. Update the .env
2. Install for Backend

## Generate Test Coverage Backend
### Without Docker
1. Up container minio from docker compose
    ```
    docker-compose up --build minio -d
    ```
2. Create .env from .env.example
    ```
    cd backend
    cp .env.example .env
    ```
3. Update the .env value for minio host
    ```
    APP_MINIO_HOST=localhost:9000
    ```
4. Test & Generate Cover.out
    ```
    cd backend
    go test ./... -coverprofile ./cover.out -covermode atomic -coverpkg ./...
    ```
5. Get cover.html
    ```
    cd backend
    go tool cover -html cover.out -o cover.html
    ```
6. Get Total Coverage
    ```
    # Show All Percentages
    cd backend
    go tool cover -func cover.out

    # On Powershell
    cd backend
    go tool cover -func cover.out |  Select-String 'total:'
    
    # On Unix
    cd backend
    go tool cover -func cover.out | grep total:

    # On 
    cd backend
    go tool cover -func cover.out | findstr total:
    ```

### Use Docker

1. Up container minio from docker compose
    ```
    docker-compose up --build minio -d
    ```
1. Generate Coverage with docker-compose
    ```
    docker-compose up --build customerbackend-test -d
    ```
2. Get cover.out and cover.html file from container
    ```
    docker cp customerbackend-test:/app/cover.out ./backend/cover.out
    docker cp customerbackend-test:/app/cover.html ./backend/cover.html
    ```
3. Get Total Coverage
    ```
    docker cp customerbackend-test:/app/coverage_percentage.txt ./backend/coverage_percentage.txt

    or just show result after copy cover.out
    
    # On Powershell
    cd backend
    go tool cover -func cover.out |  Select-String 'total:'
    
    # On Unix
    cd backend
    go tool cover -func cover.out | grep total:

    # On CMD
    cd backend
    go tool cover -func cover.out | findstr total:
    ```
## Test Threshold Backend Using Github Action
1. Generate cover.out First
2. Modify if needed the workflows file
    ```
    .github\workflows\coverage-threshold.yml
    ```
3. Currently action works on branch ***test/threshold***

## Postman Collection
```
Simple Dashboard Backend.postman_collection.json
```

## More Setup FE & BE Locally

### Backend Setup Documentation

For detailed backend documentation & setup instructions, refer to [Backend Setup Guide](docs/BE.md).

### Frontend Setup Documentation

For detailed frontend setup documentation & instructions, refer to [Frontend Setup Guide](docs/FE.md).

### Features Documentation

For detailed System features, refer to [Feature Guide](docs/FEATURES.md).