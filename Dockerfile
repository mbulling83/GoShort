# Stage 1: Build
FROM golang:1.22.4 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o goshort ./cmd/server

# Stage 2: Production
FROM gcr.io/distroless/static:nonroot

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/goshort /app/goshort

# Expose the application port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/app/goshort"]
