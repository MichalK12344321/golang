FROM golang:1.22-alpine3.20 AS build
WORKDIR /src
COPY . .
RUN go mod download
RUN go test ./test/unit/...
RUN go build -o /app /src/cmd/web/main.go

FROM golang:1.22-alpine3.20 AS runtime
COPY --from=build /app /app
CMD [ "/app" ]