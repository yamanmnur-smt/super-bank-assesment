
# Super Bank Assesment

A brief description of what this project does and who it's for

## Install & Running - Using Docker Compose

```bash
  docker-compose up --build -d
```
#### Local address Frontend
```bash
  http://localhost:3000
```

#### Local address Backend
```bash
  http://localhost:3001
```

#### Credential Admin CMS
```bash
  username : yaman
  password : password
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
1. Generate Coverage
    ```
    cd backend 
    go test ./... -coverprofile ./cover.out -covermode atomic -coverpkg ./...
    ```
2. Generate visualize Coverage
    ```
    cd backend
    go tool cover -html cover.out -o cover.html
