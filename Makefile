lint: get-linter
	golangci-lint run --timeout=5m

get-linter:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)

run:
	go run cmd/server/main.go