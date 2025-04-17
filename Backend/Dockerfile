# # --------- Build Stage ---------
#     FROM golang:1.24 AS builder

#     # Create the app directory inside the container
#     WORKDIR /app
    
#     # Copy Go module files and download dependencies
#     COPY go.mod go.sum ./
#     RUN go mod download
    
#     # Copy the source code into the container
#     COPY . .
    
#     # Build the Go app from cmd/app
#     RUN go build -o server ./cmd/app
    
#     # --------- Run Stage ---------
#     FROM debian:bullseye-slim
    
#     # Copy the binary from the builder stage
#     COPY --from=builder /app/server .
    
#     # Cloud Run listens on port 8080 by default
#     EXPOSE 8080
    
#     # Set the default command to run your server
#     CMD ["./server"]
# --------- Build Stage ---------
    FROM golang:1.24 AS builder

    # Set environment variables for static build
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    
    # Create the app directory inside the container
    WORKDIR /app
    
    # Copy Go module files and download dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the source code into the container
    COPY . .
    
    # Build the Go app from cmd/app
    RUN go build -o server ./cmd/app
    
    # --------- Run Stage ---------
    FROM debian:bullseye-slim
    
    # Copy the binary from the builder stage
    COPY --from=builder /app/server .
    
    # Cloud Run / Railway listens on port 8080 by default
    EXPOSE 8080
    
    # Set the default command to run your server
    CMD ["./server"]
    