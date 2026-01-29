# justfile

# Default task
default:
    just --list

# Run application
run:
    go run ./cmd/shrimp

# Build application
build:
    go build -o bin/shrimp.exe ./cmd/shrimp

# Run tests
test:
    go test ./...

# Format code
fmt:
    go fmt ./...

# Tidy modules
tidy:
    go mod tidy

# Clean build artifacts
clean:
    rm -rf bin