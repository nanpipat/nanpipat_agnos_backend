# Build stage
FROM golang:1.23rc1-alpine3.20 AS builder

# Install necessary tools and libraries
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh tzdata build-base

# Set the working directory
WORKDIR /app

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN go build -o main .

# Final stage
FROM alpine:3.17

# Copy the binary and timezone data from the build stage
COPY --from=builder /app/main /main
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Set timezone environment variable
ENV TZ=Asia/Bangkok

# Run the application
CMD ["./main"]
