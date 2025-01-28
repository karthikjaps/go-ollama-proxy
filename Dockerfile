# Use the official Go image as the base image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o ollama-proxy main.go

# Use a minimal base image for the runtime
FROM debian:bullseye-slim

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ollama-proxy .

# Expose the proxy port
EXPOSE 8080

# Run the proxy server
CMD ["./ollama-proxy"]
