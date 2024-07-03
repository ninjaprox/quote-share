# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:3.18

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy fonts from the builder stage
COPY fonts fonts

# Run the binary
CMD ["./main"]
