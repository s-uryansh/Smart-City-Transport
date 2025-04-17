# Use a more recent base image with the required GLIBC version
FROM ubuntu:22.04 as builder

# Install Go and other necessary dependencies
RUN apt-get update && apt-get install -y golang-go

WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o server ./cmd/app

# Use a smaller runtime image, also based on Ubuntu 22.04
FROM ubuntu:22.04

# Copy the Go server binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the app will run on
EXPOSE 8080

# Define the default command to run your app
CMD ["./server"]
