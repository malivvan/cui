export CGO_ENABLED := 0W
.DEFAULT_GOAL := help

.PHONY: validate
validate: check-fmt check-static test ## Validates the go code format, executes staticcheck and runs tests

.PHONY: test
test: ## Run tests
	@CGO_ENABLED=1 gotestsum --format-hide-empty-pkg -- -race -v ./...

.PHONY: check-static
check-static: ## Run static analysis
	@staticcheck -checks all ./...

.PHONY: check-fmt
check-fmt: ## Check go format
	@gofmt_out=$$(gofmt -d -e . 2>&1) && [ -z "$${gofmt_out}" ] || (echo "$${gofmt_out}" 1>&2; exit 1)

.PHONY: fmt
fmt: ## Formats the go code
	@go fmt ./...

.PHONY: install
install: ## Installs required tools
	@echo "installing required tools..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install gotest.tools/gotestsum@latest

.PHONY: help
help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
