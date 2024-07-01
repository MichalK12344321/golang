IMAGE_NAME ?= scheduler

docker-build: swag
	docker build \
	--pull \
	--tag="$(IMAGE_NAME):latest" \
	.

swag:
	~/go/bin/swag init --parseDependency -g scheduler/server.go

run:
	go run .
