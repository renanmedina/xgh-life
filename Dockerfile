## multi stage build
# builder stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o -ldflags '-s -w' server server.go

# runtime stage
FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE ${PORT}
CMD ["./server"]