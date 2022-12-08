all: build check test

build:
	go mod tidy && go build ./...

check:
	cd ./client && staticcheck .

test:
	cd ./client && go test
