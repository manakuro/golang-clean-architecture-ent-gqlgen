install:
	brew install yq
	brew install pre-commit
	pre-commit --version
	pre-commit install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/cosmtrek/air@v1.27.3

# Start dev server
start:
	air

.PHONY: install start
