install:
	brew install yq
	brew install pre-commit
	pre-commit --version
	pre-commit install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/cosmtrek/air@v1.27.3

# Set up database
setup_db:
	./bin/init_db.sh

# Migrate scheme in ent to database
migrate_schema:
	go run ./cmd/migration/main.go

# Start dev server
start:
	air

# Testing
setup_test_db:
	./bin/init_db_test.sh

test_repository:
	go test ./pkg/adapter/repository/...

# E2E
setup_e2e_db:
	./bin/init_db_e2e.sh

e2e:
	go test ./test/e2e/...

.PHONY: install start setup_db migrate_schema setup_test_db setup_e2e_db e2e
