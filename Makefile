.PHONY: test vet

TEST_FORMAT ?= pkgname

install:
	@go install gotest.tools/gotestsum@latest

test-race:
	@CGO_ENABLED=1 gotestsum --format $(TEST_FORMAT) --format-hide-empty-pkg --hide-summary skipped --raw-command -- go test -json ./...

test:
	@CGO_ENABLED=0 gotestsum --format $(TEST_FORMAT) --format-hide-empty-pkg --hide-summary skipped --raw-command -- go test -json ./...

vet:
	@go vet -composites=false ./...

