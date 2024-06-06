# # Dockerfile
# FROM golang:latest

# WORKDIR /app

# COPY . .

# RUN go build -o server .

# CMD ["./server"]



# following apline is creating compatibility issues with the go version
# 
# # Stage 1: Build the Go application
# FROM golang:1.16 AS builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o server .

# # Stage 2: Create the final image
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=builder /app/server .
# CMD ["./server"]



# this one is working well enough
# Dockerfile
# Stage 1: Build the Go application
FROM golang:1.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 2: Create the final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /root/
COPY --from=builder /app/server .
CMD ["./server"]
