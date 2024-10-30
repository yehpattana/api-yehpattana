# Use golang:1.20-alpine for a lightweight base
FROM golang:1.20-alpine

# Install necessary dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory within the container
WORKDIR /go/src/app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the application's port
EXPOSE 8080

# Set the default command
CMD ["go", "run", "main.go"]
