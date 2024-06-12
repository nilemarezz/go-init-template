.PHONY: dev sit test

dev:
	go run main.go -env=dev -date=${date}

sit:
	go run main.go -env=sit -date=${date}

test:
	go test ./...

test-coverage:
	go test -cover ./... 

swag_init:
	swag init -g main.go -o ./docs

build: clean
	mkdir -p build/bin
	go build -o build/bin/app

clean:
	rm -rf build/bin