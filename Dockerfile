# ---- Build stage ----
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o server ./cmd/server

# ---- Run stage ----
FROM alpine:latest

WORKDIR /app

# Copy only the binary
COPY --from=builder /app/server .

# Set environment variable
ENV PORT=8080

EXPOSE 8080

CMD ["./server"]
