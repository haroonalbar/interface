build:
	@go build -o ./bin/inter

run: build
	@./bin/inter

test:
	@go test ./... -v
