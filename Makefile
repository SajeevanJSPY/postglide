run: build
	@go run ./cmd/postglide

build:
	@go build -o postglide ./cmd/postglide

install:
	@go mod tidy

test:
	@go test ./... -v

vet:
	@go vet ./...

fmt:
	@go fmt ./...

clean:
	@rm -rf postglide

