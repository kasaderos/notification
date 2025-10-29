# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build all applications with static linking and proper binary names
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/tgbotapi ./cmd/tgbotapi
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/crawler ./cmd/crawler
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/percolator ./cmd/percolator

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binaries from the builder stage
COPY --from=builder /out/tgbotapi .
COPY --from=builder /out/crawler .
COPY --from=builder /out/percolator .

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Expose port (if needed for health checks)
EXPOSE 8080

# Default command runs the telegram bot
# Use 'docker run <image> ./crawler' to run the crawler
# Use 'docker run <image> ./percolator' to run the percolator
CMD ["./tgbotapi"]