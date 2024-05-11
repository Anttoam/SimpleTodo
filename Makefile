db-shall:
	turso db shell todo

lint:
	golangci-lint run

redis:
	docker-compose exec redis redis-cli -h localhost -p 6379

.PHONY: db-shall lint redis
