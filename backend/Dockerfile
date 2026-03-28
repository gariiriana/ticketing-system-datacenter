# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/api/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
# Copy firebase key if it exists (user needs to provide this)
# COPY --from=builder /app/serviceAccountKey.json .

EXPOSE 8080

CMD ["./main"]
