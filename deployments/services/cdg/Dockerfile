FROM golang:1.22.4-bullseye as build
WORKDIR /src
COPY . .
RUN go mod download
RUN go test ./test/unit/...
RUN go build -o /app /src/cmd/cdg/main.go

FROM golang:1.22.4-bullseye as runtime
COPY --from=build /app /app
CMD [ "/app" ]