all: tidy format build test vet

tidy:
	go mod tidy

format:
	gofumpt -l -w .

build:
	go build ./...

test:
	go test ./... -cover

vet:
	go vet ./...