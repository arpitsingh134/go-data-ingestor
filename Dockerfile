FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod and go.sum first to cache modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . ./
RUN go build -o main ./cmd/main.go

CMD ["./main"]
