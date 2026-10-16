run:
	go run ./cmd/http-server.go

build:
	go build -o ./bin/ ./...

build-run: build
	./bin/http-go