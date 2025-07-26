# ---- Build Stage ----
FROM golang:1.23.0 as builder

WORKDIR /app

COPY . .

RUN go build -o myapp

# ---- Final Stage ----
FROM debian:bullseye-slim

WORKDIR /app

# Copy only the binary from the builder stage
COPY --from=builder /app/myapp .

# Only run the binary (no shell or dev tools)
CMD ["./myapp"]
