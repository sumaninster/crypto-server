# First stage: build the Go application
FROM golang:alpine AS builder

# Set the working directory and copy the source code
WORKDIR /app
COPY . .

# Build the application with optimizations and no debug information
RUN go build -ldflags="-s -w" -o app

# Second stage: create a minimal Docker image
FROM alpine:latest

# Set the working directory and copy the compiled binary from the previous stage
WORKDIR /app
COPY --from=builder /app/app .

# Run the application
ENTRYPOINT ["./app"]
EXPOSE 8080