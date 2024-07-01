FROM golang:1.18 AS builder
ARG SERVICE

WORKDIR /app
COPY $SERVICE/go.mod $SERVICE/go.sum ./
COPY $SERVICE .
COPY common /common
RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
