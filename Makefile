build:
	@go build -o bin/spark -ldflags "-s -w" cmd/spark/main.go