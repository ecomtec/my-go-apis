# Use the official Golang image as base
FROM golang:1.23.0

# Set working directory
WORKDIR /app

# Copy all files to container
COPY . .

# Build the Go app
RUN go build -o myapp

# Command to run the app
CMD ["./myapp"]