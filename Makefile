include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Clean up all environment files? Data loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Environment files were deleted"; \
	else \
		echo "Environment cleanup cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder
	
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Missing required parameter seq. Example: make migrate-create seq=init"; \
		exit 1; \
	fi;
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Missing required parameter action. Example: migrate-action action=up 1"; \
		exit 1; \
	fi; 
	@docker compose run  --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)"

todoapp-run:
	@go run cmd/todoapp/main.go