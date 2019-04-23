all: index.html

index.html: format
	go run cmd/generate/main.go

format:
	go run cmd/prettyprint/main.go
