# Build stage
FROM golang:1.17 AS build-env

WORKDIR /app

# Copy the Go source files
COPY . .

# Compile the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o goapp

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build-env /app/goapp .

# Expose the application on port 8000
EXPOSE 8000

# Command to run the application
CMD ["./goapp"]
