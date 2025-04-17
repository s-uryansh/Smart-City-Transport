FROM ubuntu:22.04 AS builder

# Install dependencies: CA certs, wget, and unzip tools
RUN apt-get update && apt-get install -y wget ca-certificates tzdata

# Set timezone to UTC to avoid SSL validation issues (optional)
ENV TZ=UTC

# Install Go manually
ENV GO_VERSION=1.22.2
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/bin/go

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app

COPY go.mod go.sum ./
COPY .env .
RUN go mod download

COPY . .

RUN go build -o server ./cmd/app

# Runtime image
FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/server .

CMD ["./server"]
