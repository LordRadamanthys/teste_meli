FROM golang:1.23.4

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/process_ordes src/main.go

EXPOSE 8080

ENTRYPOINT [ "/app/process_ordes" ]