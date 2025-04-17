# Use a recent base image
FROM ubuntu:22.04 as builder

# Install Go and CA certificates
RUN apt-get update && apt-get install -y golang-go ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/app

# Final stage with runtime only
FROM ubuntu:22.04

# Install CA certificates in runtime too (in case your server uses HTTPS)
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/server .

CMD ["./server"]
