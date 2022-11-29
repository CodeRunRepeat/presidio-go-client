all: build check test

build:
	cd ./generated && go mod tidy && go build
	cd ./client && go mod tidy && go build

check:
	cd ./client && staticcheck .

test:
	cd ./client && go test
