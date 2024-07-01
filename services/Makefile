SERVICE_NAME ?= scheduler
IMAGE_NAME = $(SERVICE_NAME)

swag:
	cd $(SERVICE_NAME) && make swag

docker-build: swag
	docker build \
	--pull \
	--build-arg SERVICE=$(IMAGE_NAME) \
	--tag="$(IMAGE_NAME):latest" \
	.
