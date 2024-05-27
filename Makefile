.PHONY: dev sit test

dev:
	go run cmd/main.go -env=dev

sit:
	go run cmd/main.go -env=sit

test:
	go test ./...

test-coverage:
	go test -cover ./... 

swag_init:
	swag init -g ./cmd/main.go -o ./docs