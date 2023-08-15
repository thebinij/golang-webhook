# Start with the official Go image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-webhook

# Start a new, smaller image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=builder /app/go-webhook .

# Expose the port the Go server will listen on
EXPOSE $PORT

# Install necessary certificates for Alpine
RUN apk --no-cache add ca-certificates

# Start the Go server
CMD ["./go-webhook"]
