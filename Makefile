build:
	@go build -o /bin/api

run: 
	@go run main.go

.PHONY: seed

seed:
	go run ./seed


test:
	@go test ./...
