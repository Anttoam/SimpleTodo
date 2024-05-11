db-shall:
	turso db shell todo

lint:
	golangci-lint run

.PHONY: db-shall lint
