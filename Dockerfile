# Stage 1: Build the Go binary
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with CGO disabled (allows it to run smoothly on Alpine scratch envs)
RUN CGO_ENABLED=0 GOOS=linux go build -o taskflow main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install bash and curl (required for goose and entrypoint scripting)
RUN apk --no-cache add bash curl

# Install Goose for DB Migrations
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

# Set working directory for the runtime container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/taskflow .

# Copy migrations folder and entrypoint script
COPY migrations/ ./migrations/
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

# Expose port required for APIs
EXPOSE 8080

# Run entrypoint script
ENTRYPOINT ["./entrypoint.sh"]
