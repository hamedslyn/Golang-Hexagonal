.PHONY: run test

run:
	docker-compose up --build

test:
	go test ./...

