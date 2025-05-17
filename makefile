.PHONY: build run docker

build:
	go build -o bin/server ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

docker:
	docker build -t german-concordance .
	docker run -p 8080:8080 german-concordance