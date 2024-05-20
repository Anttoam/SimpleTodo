include .env
export $(shell sed 's/=.*//' .env)

db-shall:
	turso db shell todo

check-code:
	gofmt -w .
	goimports -w .
	golangci-lint run

redis:
	docker-compose exec redis redis-cli -h localhost -p 6379

atlas-inspect:
	atlas schema inspect --url "sqlite://migration/todo.db" > migration/schema.hcl

atlas-apply:
	atlas schema apply --url "sqlite://migration/todo.db" --to "file://migration/schema.hcl"

migrate:
	atlas schema apply -u "${TURSO_DB_URL}?authToken=${TURSO_DB_TOKEN}" --to sqlite://migration/todo.db

.PHONY: db-shall check-code redis atlas-inspect atlas-apply migrate
