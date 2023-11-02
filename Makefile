all: build

build:
	go build -o bin/trueconf cmd/app/main.go
