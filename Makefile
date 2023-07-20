run-database-tests:
	docker-compose stop account-db-test
	docker-compose rm -f account-db-test
	docker-compose up -d account-db-test
	sleep 3
	cd services/account; go test ./internal/configuration/database/sqlc; cd ../..
