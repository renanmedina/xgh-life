## multi stage build
ARG PORT=8080
# builder stage
FROM golang:1.25.7-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags "-s -w" -o server server.go

# runtime stage
FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE ${PORT}
CMD ["./server"]