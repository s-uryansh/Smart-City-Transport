# --------- Build Stage ---------
    FROM golang:1.24.2 AS builder

    # Create app directory inside container
    WORKDIR /app
    
    # Copy go module files and download dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the source code
    COPY . .
    
    # Build the Go app from cmd/app
    RUN go build -o server ./cmd/app
    
    # --------- Run Stage ---------
    FROM debian:bullseye-slim
    
    # Copy the binary from builder stage
    COPY --from=builder /app/server .
    
    # Cloud Run listens on 8080 by default
    EXPOSE 8080
    
    CMD ["./server"]
    