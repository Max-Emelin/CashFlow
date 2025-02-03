FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/main .

COPY --from=builder /app/migrations ./migrations

COPY .env .env

CMD ["./main"]
