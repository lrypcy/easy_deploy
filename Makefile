.PHONY: build
build:
	@go build .

.PHONY: fmt
fmt: 
	@gofmt -s -w .

.PHONY: clean
clean:
	@go clean