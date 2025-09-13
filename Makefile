# GoFunc Makefile
# Common development tasks for the GoFunc utility library

.PHONY: help test coverage bench lint fmt clean install build tidy deps check-deps mod-update

# Default target
.DEFAULT_GOAL := help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Test parameters
TEST_FLAGS=-v -race
COVERAGE_FLAGS=-v -race -coverprofile=coverage.out -covermode=atomic
BENCH_FLAGS=-bench=. -benchmem

help: ## Display this help message
	@echo "GoFunc Development Commands:"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

# Development commands
test: ## Run all tests
	@echo "Running tests..."
	$(GOTEST) $(TEST_FLAGS) ./...

test-short: ## Run tests without race detection (faster)
	@echo "Running short tests..."
	$(GOTEST) -v ./...

coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	$(GOTEST) $(COVERAGE_FLAGS) ./...
	@echo "Coverage report generated: coverage.out"
	@echo "View coverage in browser:"
	@echo "  go tool cover -html=coverage.out"

coverage-html: coverage ## Generate and open HTML coverage report
	@echo "Opening coverage report in browser..."
	$(GOCMD) tool cover -html=coverage.out

bench: ## Run benchmark tests
	@echo "Running benchmarks..."
	$(GOTEST) $(BENCH_FLAGS) ./...

bench-cpu: ## Run CPU profiling benchmarks
	@echo "Running CPU profiling benchmarks..."
	$(GOTEST) -bench=. -benchmem -cpuprofile=cpu.prof ./...
	@echo "CPU profile generated: cpu.prof"
	@echo "View profile: go tool pprof cpu.prof"

bench-mem: ## Run memory profiling benchmarks
	@echo "Running memory profiling benchmarks..."
	$(GOTEST) -bench=. -benchmem -memprofile=mem.prof ./...
	@echo "Memory profile generated: mem.prof"
	@echo "View profile: go tool pprof mem.prof"

# Code quality
lint: check-deps ## Run linters
	@echo "Running linters..."
	$(GOLINT) run ./...

fmt: ## Format Go code
	@echo "Formatting code..."
	$(GOFMT) -s -w .

fmt-check: ## Check if code is formatted
	@echo "Checking code formatting..."
	@if [ -n "$$($(GOFMT) -l .)" ]; then \
		echo "Code is not formatted. Run 'make fmt' to fix."; \
		$(GOFMT) -l .; \
		exit 1; \
	fi

vet: ## Run go vet
	@echo "Running go vet..."
	$(GOCMD) vet ./...

# Dependencies
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	$(GOMOD) download

tidy: ## Clean up dependencies
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

mod-update: ## Update all dependencies
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

check-deps: ## Check if required tools are installed
	@echo "Checking dependencies..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)

install-deps: ## Install development dependencies
	@echo "Installing development dependencies..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Build and install
build: ## Build the package
	@echo "Building package..."
	$(GOBUILD) ./...

install: ## Install the package
	@echo "Installing package..."
	$(GOCMD) install ./...

# Cleanup
clean: ## Clean build artifacts and test files
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -f coverage.out
	rm -f cpu.prof
	rm -f mem.prof

# CI/CD helpers
ci-test: deps fmt-check vet lint test ## Run all CI tests
	@echo "All CI tests passed!"

ci-coverage: deps test coverage ## Run CI with coverage
	@echo "CI with coverage completed!"

# Release helpers
check-clean: ## Check if working directory is clean
	@echo "Checking if working directory is clean..."
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "Working directory is not clean. Commit changes first."; \
		git status --short; \
		exit 1; \
	fi

# Documentation
docs: ## Generate and serve documentation locally
	@echo "Starting documentation server..."
	@echo "Visit: http://localhost:6060/pkg/github.com/kingrain94/gofunc/"
	godoc -http=:6060

# Examples
run-examples: ## Run all example programs
	@echo "Running basic usage example..."
	@cd examples/basic && go run main.go
	@echo ""
	@echo "Running advanced usage example..."
	@cd examples/advanced && go run main.go

test-examples: ## Test all example functions
	@echo "Testing example functions..."
	$(GOTEST) -run Example ./...

# Performance testing
perf: bench ## Alias for bench
	@echo "Performance testing completed!"

# All-in-one commands
check: fmt-check vet lint test ## Run all checks (format, vet, lint, test)
	@echo "All checks passed!"

full-test: clean deps check coverage ## Run comprehensive testing
	@echo "Full testing suite completed!"

# Development workflow
dev-setup: install-deps deps ## Set up development environment
	@echo "Development environment ready!"

pre-commit: fmt vet lint test ## Run pre-commit checks
	@echo "Pre-commit checks passed!"

# Version info
version: ## Show Go version and module info
	@echo "Go version:"
	@$(GOCMD) version
	@echo ""
	@echo "Module info:"
	@$(GOCMD) list -m

# Help for specific commands
test-help: ## Show test command options
	@echo "Test Commands:"
	@echo "  make test        - Run all tests with race detection"
	@echo "  make test-short  - Run tests without race detection (faster)"
	@echo "  make coverage    - Run tests with coverage report"
	@echo "  make bench       - Run benchmark tests"
	@echo "  make bench-cpu   - Run benchmarks with CPU profiling"
	@echo "  make bench-mem   - Run benchmarks with memory profiling"

lint-help: ## Show linting command options
	@echo "Linting Commands:"
	@echo "  make lint        - Run golangci-lint"
	@echo "  make fmt         - Format code with gofmt"
	@echo "  make fmt-check   - Check if code is formatted"
	@echo "  make vet         - Run go vet"
