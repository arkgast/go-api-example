build:
	go build -o ./bin/main ./cmd/api/main.go

dev:
	air

test:
	go test ./...

clean:
	rm -f ./bin/main
