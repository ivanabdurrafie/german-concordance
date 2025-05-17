FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /server .
COPY configs/config.yaml ./configs/

EXPOSE 8080
CMD ["./server"]