# Use the official Golang image to create a build artifact.
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /server .

# Use a minimal base image for the final image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /server /server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/server"]
