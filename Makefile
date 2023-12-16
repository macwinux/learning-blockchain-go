build:
	go build -o ./bin/blockchain

run: build
	./bin/blockchain

r:
	go run main.go

test:
	go test -v ./...