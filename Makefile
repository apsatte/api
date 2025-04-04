include .env
export

run:
	go run ./cmd/main.go

test:
	go clean -testcache
	go test -p=1 ./internal/...

swag:
	swag init -g ./internal/transport/http/http.go

database_up:
	docker run --name test-db --rm \
	-e POSTGRES_USER=${DB_USER} \
	-e POSTGRES_PASSWORD=${DB_PASS} \
	-e POSTGRES_DB=${DB_NAME} \
	-e PGDATA=/var/lib/postgresql/data \
	-p 5432:5432 \
	-v $(CURDIR)/migrations:/docker-entrypoint-initdb.d \
	-d postgres:15.3-bullseye

database_down:
	docker stop test-db