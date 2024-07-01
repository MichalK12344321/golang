FROM golang:1.18

WORKDIR /app/scheduler
COPY go.mod ./
RUN go mod download
COPY ../common ../common
COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
