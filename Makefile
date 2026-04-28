.PHONY: help build install validate list test fmt sidecar-init clean

help:
	@echo "ann-research vault — Make targets"
	@echo ""
	@echo "  build          Build the Go vault CLI to ./bin/vault"
	@echo "  install        Install the vault CLI to \$$GOBIN (or \$$GOPATH/bin)"
	@echo "  validate       Validate all artifacts in the vault"
	@echo "  list           List all artifacts in the vault"
	@echo "  test           Run Go tests"
	@echo "  fmt            Run gofmt + go vet"
	@echo "  sidecar-init   Initialize the Python sidecar with uv"
	@echo "  clean          Remove build outputs"

build:
	go build -o bin/vault ./cmd/vault

install:
	go install ./cmd/vault

validate: build
	./bin/vault validate

list: build
	./bin/vault list

test:
	go test ./...

fmt:
	gofmt -w .
	go vet ./...

sidecar-init:
	cd sidecar && uv sync

clean:
	rm -rf bin/
