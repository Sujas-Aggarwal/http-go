run:
	go run ./cmd/http-server/main.go

build:
	go build -o ./bin/ ./...

build-run: build
	./bin/http-go

run-api:
	go run ./cmd/api-server/main.go