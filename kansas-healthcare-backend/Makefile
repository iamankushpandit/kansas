.PHONY: test test-verbose test-coverage run build clean

# Run tests
test:
	go test ./...

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...

# Run tests with detailed coverage report
test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o kansas-healthcare-api main.go

# Clean build artifacts
clean:
	rm -f kansas-healthcare-api coverage.out coverage.html

# Install dependencies
deps:
	go mod download
	go mod tidy