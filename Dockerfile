# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy Go module file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy application source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/server .

# Application port
EXPOSE 3000

# Start application
CMD ["./server"]