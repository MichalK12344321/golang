IMAGE_NAME ?= ssh-host
SSHD_PORT ?= 22

.DEFAULT_GOAL := help
.PHONY: *

docker-build: ## Builds Docker image
	docker build \
	--pull \
	--build-arg PORT="$(SSHD_PORT)" \
	--tag="$(IMAGE_NAME):latest" \
	.

help: ## Displays this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ \
	{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
