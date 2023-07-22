NETWORK_NAME := internal-net

.PHONY: create_network

up-all:
	@if ! docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 ; then \
		@docker network create internal-net \
		echo "Network $(NETWORK_NAME) created." ; \
	else \
		echo "Network $(NETWORK_NAME) already exists." ; \
	fi
	@docker compose up -d --scale kafkaui=0
	@printf "Up and running"

run-db-tests:
	@docker compose stop db-test
	@docker compose rm -f db-test
	@docker compose up -d db-test
	@printf "Up and running the test database"
	@for i in 1 2 3 4 5 6; do \
		sleep 0.5; \
		echo -n '.'; \
	done
	@echo
	@printf "Tests executiing: \n"
	@cd services/account; go test ./internal/configuration/database/sqlc; cd ../..
	@cd services/transaction; go test ./internal/configuration/database/sqlc; cd ../..
	@printf "Down and removing the test database"
	@docker compose stop db-test
	@docker compose rm -f db-test

run-unit-tests:
	@printf "Running unit tests\n"
	@cd services/account; go test -cover ./internal/application/services; cd ../..
	@cd services/transaction; go test -cover ./internal/application/services; cd ../..
