# Dockerfile for Go backend

# Build the Go app
FROM golang:latest AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Run the Go app
FROM golang:latest

WORKDIR /app
COPY --from=build /app/main .
EXPOSE 5001
CMD ["./main"]
