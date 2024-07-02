NETWORK_NAME ?= test-network

.DEFAULT_GOAL := help
.PHONY: *

create-env:
	echo "NETWORK_NAME=${NETWORK_NAME}" > .env

create-network:
	docker network ls | grep -E "\s${NETWORK_NAME}\s" || docker network create --internal ${NETWORK_NAME}

run: create-env create-network ## Run docker compose
	docker-compose -f docker-compose.yaml up -d

stop: ## Stop docker compose
	docker-compose -f docker-compose.yaml stop

restart: stop run ## Restart docker compose

help: ## Displays this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ \
	{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
