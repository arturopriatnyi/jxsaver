.PHONY: build
build:
	go build -o ./build/jxsaver ./cmd/jxsaver/main.go

run:
	go run ./cmd/jxsaver/main.go

start:
	./build/jxsaver
