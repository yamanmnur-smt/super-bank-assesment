FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN cp .env.example .env

RUN go build -o backend .

EXPOSE 3001

CMD ["./backend"]