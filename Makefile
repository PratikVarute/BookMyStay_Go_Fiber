build:
	@go build -o /bin/api

run: 
	@go run main.go

test:
	@go test ./...
