GLOLANGCI_LINT_VERSION=1.30.0

test: .FORCE
	go test -v ./...

.PHONY: .FORCE

install-test:
	@(cd; GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GLOLANGCI_LINT_VERSION))

lint:
	golangci-lint run -v
