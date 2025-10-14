# Tangent Makefile
# Easy development and release management

.PHONY: build test clean install release help

# Default target
all: build

# Build with version injection
build:
	@echo "Building Tangent..."
	@./scripts/build.sh

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f tangent
	@go clean

# Install locally
install: build
	@echo "Installing to ~/.local/bin..."
	@mkdir -p ~/.local/bin
	@cp tangent ~/.local/bin/
	@echo "✓ Installed: ~/.local/bin/tangent"

# Create a new release
release:
	@echo "Creating release..."
	@read -p "Version (e.g., v0.1.0-beta.5): " version; \
	./scripts/validate-release.sh "$$version" && \
	git tag $$version && \
	git push origin $$version && \
	echo "✓ Released: $$version"

# Show help
help:
	@echo "Tangent Makefile"
	@echo ""
	@echo "Commands:"
	@echo "  make build     - Build with version injection"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make install   - Install to ~/.local/bin"
	@echo "  make release   - Create new release"
	@echo "  make help      - Show this help"
