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

mockur:
	mockgen -source internal/usecase/user_usecase.go -destination internal/usecase/mock/user_repository.go

mocktr:
	mockgen -source internal/usecase/todo_usecase.go -destination internal/usecase/mock/todo_repository.go

mockuu:
	mockgen -source internal/controller/user_controller.go -destination internal/controller/mock/user_usecase.go

swag:
	swag init -g cmd/main.go
	swag fmt

.PHONY: db-shall check-code redis atlas-inspect atlas-apply migrate mockur mocktr mockuu swag
