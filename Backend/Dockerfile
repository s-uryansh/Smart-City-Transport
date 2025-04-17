# Use a more recent base image with the required GLIBC version
FROM ubuntu:22.04 as builder

# Install Go and other necessary dependencies
RUN apt-get update && apt-get install -y golang-go

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/app

# Use a smaller runtime image
FROM ubuntu:22.04

COPY --from=builder /app/server .

CMD ["./server"]
