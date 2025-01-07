include .env

DOCKER_EXEC = docker compose exec web 
DB_URL = postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: api-sh
api-sh: ## Run sh in the api container
	$(DOCKER_EXEC) sh

.PHONY: migrations
migrations: ## Create migrations using golang-migrate
	@read -p "Enter migration name: " NAME; \
	$(DOCKER_EXEC) migrate create -ext sql -dir database/migrations -seq $$NAME
	

.PHONY: migrate
migrate: ## Run migrations using golang-migrate
	$(DOCKER_EXEC) migrate -database $(DB_URL) -path database/migrations up

.PHONY: rollback
rollback: ## Rollback migrations using golang-migrate
	$(DOCKER_EXEC) migrate -database $(DB_URL) -path database/migrations down 1

.PHONY: flush-migrations
flush-migrations: ## Rollback all migrations using golang-migrate
	$(DOCKER_EXEC) migrate -database $(DB_URL) -path database/migrations down

.PHONY: clean-migrations
clean-migrations: ## Rollback all migrations using golang-migrate and remove all migrations files
	$(DOCKER_EXEC) migrate -database $(DB_URL) -path database/migrations down 0
	rm -rf database/migrations/*.sql