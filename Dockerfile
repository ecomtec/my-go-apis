# Stage 1: Build the Go app
FROM golang:1.23.0-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set workdir
WORKDIR /app

# Copy source
COPY . .

# Build with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp

# Stage 2: Run with scratch (empty base image)
FROM scratch

# Copy binary from builder
COPY --from=builder /app/myapp /myapp

# Command to run the app
ENTRYPOINT ["/myapp"]
