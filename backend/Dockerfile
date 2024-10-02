# Use a base Go image
FROM golang:1.23.0-alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o out ./cmd

# Expose port 3000
EXPOSE 3000

# Default command (this can be overridden by docker-compose)
CMD ["./out"]



# FROM golang:1.23.0

# WORKDIR /usr/src/app

# # RUN go install github.com/air-verse/air@latest

# COPY . .
# RUN go mod tidy