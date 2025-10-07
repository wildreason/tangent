#!/bin/bash
# Quick install script for Wildreason Character Builder
# Usage: curl -sSL https://wildreason.com/install.sh | bash

set -e

echo "◢ Installing Wildreason Character Builder..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "✗ Error: Go is not installed. Please install Go 1.20+ first."
    echo "  Visit: https://go.dev/dl/"
    exit 1
fi

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64";;
    aarch64|arm64) ARCH="arm64";;
    *) echo "✗ Unsupported architecture: $ARCH" && exit 1;;
esac

BINARY_NAME="character-builder-${OS}-${ARCH}"
if [ "$OS" = "windows" ]; then
    BINARY_NAME="${BINARY_NAME}.exe"
fi

# Download URL (replace with actual release URL)
DOWNLOAD_URL="https://releases.wildreason.com/characters/v0.0.1/${BINARY_NAME}"

# Install location
INSTALL_DIR="$HOME/.wildreason/bin"
mkdir -p "$INSTALL_DIR"

echo "◢ Downloading $BINARY_NAME..."
# curl -sSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/character-builder"
# For now, build from source
echo "◢ Building from source..."
go install github.com/wildreason/characters/tools/builder@latest

chmod +x "$INSTALL_DIR/character-builder"

# Add to PATH suggestion
echo ""
echo "✓ Installation complete!"
echo ""
echo "Add to your PATH:"
echo "  export PATH=\"\$HOME/.wildreason/bin:\$PATH\""
echo ""
echo "Then run:"
echo "  character-builder"

