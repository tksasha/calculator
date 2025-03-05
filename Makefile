MODULE=github.com/tksasha/calculator

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go tool -modfile go.tool.mod gofumpt -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@go tool -modfile go.tool.mod golangci-lint run

.PHONY: test
test:
	@echo "go test"
	@go test ./...

.PHONY: prepare
prepare:
	@if [ ! -f go.mod ]; then go mod init $(MODULE); go mod tidy; fi
	@if [ ! -f go.tool.mod ]; then go mod init -modfile go.tool.mod $(MODULE); go mod tidy -modfile go.tool.mod; fi
	go get -tool -modfile go.tool.mod github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go get -tool -modfile go.tool.mod mvdan.cc/gofumpt@latest
