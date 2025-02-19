FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./src/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY src/.env .env
COPY src/db.yaml db.yaml

EXPOSE 8080

CMD ["./main"]
