# Tangent Makefile
# Easy development and release management

.PHONY: build build-cli test clean install release help generate verify-generate

# Default target
all: build-cli

# Generate character code from constants
generate:
	@echo "Generating character files from constants..."
	@cd pkg/characters/library && go generate
	@go fmt ./pkg/characters/library > /dev/null 2>&1
	@echo "✓ Generated 7 character files + themes"

# Verify generated code is up-to-date (for CI)
verify-generate:
	@echo "Verifying generated code..."
	@cd pkg/characters/library && go run generator_codegen.go -verify
	@echo "✓ Generated code is up-to-date"

# Build internal CLI tool with version injection
build-cli:
	@echo "Building tangent-cli (internal tool)..."
	@./scripts/build.sh

# Legacy target (redirects to build-cli)
build: build-cli

# Run tests (with verification that generated code is up-to-date)
test: verify-generate
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f tangent tangent-cli
	@go clean

# Install locally (internal tool)
install: build-cli
	@echo "Installing tangent-cli to ~/.local/bin..."
	@mkdir -p ~/.local/bin
	@cp tangent-cli ~/.local/bin/
	@echo "✓ Installed: ~/.local/bin/tangent-cli (internal tool)"

# Create a new release
# Usage: make release
# Force re-release: FORCE_RELEASE=1 make release
release:
	@echo "Creating release..."
	@read -p "Version (e.g., v0.1.0-beta.5): " version; \
	FORCE_RELEASE=$(FORCE_RELEASE) ./scripts/validate-release.sh "$$version" && \
	git tag $$version && \
	git push origin $$version && \
	echo "✓ Released: $$version"

# Show help
help:
	@echo "Tangent Makefile"
	@echo ""
	@echo "Commands:"
	@echo "  make generate        - Generate character files from constants.go"
	@echo "  make verify-generate - Verify generated code is up-to-date"
	@echo "  make build-cli       - Build tangent-cli (internal development tool)"
	@echo "  make test            - Run tests (with generate verification)"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make install         - Install tangent-cli to ~/.local/bin"
	@echo "  make release         - Create new release"
	@echo "  make help            - Show this help"
