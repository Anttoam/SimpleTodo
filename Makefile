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
	mockery --dir=internal/usecase --name=UserRepository --filename=user_repository.go --output internal/usecase/mocks --outpkg=mocks

mocktr:
	mockery --dir=internal/usecase --name=TodoRepository --filename=todo_repository.go --output internal/usecase/mocks --outpkg=mocks

mockuu:
	mockery --dir=internal/controller --name=UserUsecase --filename=user_usecase.go --output internal/controller/mocks --outpkg=mocks

mocktu:
	mockery --dir=internal/controller --name=TodoUsecase --filename=todo_usecase.go --output internal/controller/mocks --outpkg=mocks

swag:
	swag init -g cmd/main.go
	swag fmt

test:
	go test -v -cover ./...

.PHONY: db-shall check-code redis atlas-inspect atlas-apply migrate mockur mocktr mockuu mocktu swag test
