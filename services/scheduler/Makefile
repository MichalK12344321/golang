IMAGE_NAME ?= scheduler

docker-build: swag
	docker build \
	--pull \
	--tag="$(IMAGE_NAME):latest" \
	.

swag:
	~/go/bin/swag init -g app/app.go

run:
	go run .
