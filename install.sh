#!/bin/bash
# Tangent installer - installs CLI and makes package available
# Usage: curl -sSL https://wildreason.com/install-tangent.sh | bash

set -e

VERSION="${VERSION:-latest}"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
REPO="wildreason/tangent"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "╔══════════════════════════════════════════╗"
echo "║  Tangent Installer                       ║"
echo "║  Terminal Agent Designer                 ║"
echo "╚══════════════════════════════════════════╝"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}✗ Go is not installed${NC}"
    echo "  Install Go from: https://go.dev/dl/"
    exit 1
fi

echo -e "${GREEN}✓${NC} Go detected: $(go version | awk '{print $3}')"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}✗ Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

echo -e "${GREEN}✓${NC} Platform: $OS-$ARCH"

# Create install directory
mkdir -p "$INSTALL_DIR"
echo -e "${GREEN}✓${NC} Install directory: $INSTALL_DIR"

# Determine download URL
if [ "$VERSION" = "latest" ]; then
    DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/tangent_${OS}_${ARCH}"
else
    DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/tangent_${OS}_${ARCH}"
fi

echo ""
echo "Downloading tangent..."
echo "  From: $DOWNLOAD_URL"

# Download binary
TEMP_FILE=$(mktemp)
if curl -fsSL -o "$TEMP_FILE" "$DOWNLOAD_URL" 2>/dev/null; then
    echo -e "${GREEN}✓${NC} Downloaded tangent binary"
else
    echo -e "${RED}✗ Failed to download${NC}"
    echo "  URL: $DOWNLOAD_URL"
    echo ""
    echo "Alternative: Install via Go"
    echo "  go install github.com/$REPO/cmd/tangent@latest"
    rm -f "$TEMP_FILE"
    exit 1
fi

# Install binary
mv "$TEMP_FILE" "$INSTALL_DIR/tangent"
chmod +x "$INSTALL_DIR/tangent"
echo -e "${GREEN}✓${NC} Installed to $INSTALL_DIR/tangent"

# Install Go package
echo ""
echo "Installing Go package..."
if go install "github.com/$REPO/cmd/tangent@latest" 2>/dev/null; then
    echo -e "${GREEN}✓${NC} Go package installed"
else
    echo -e "${YELLOW}⚠${NC}  Package install skipped (optional)"
fi

# Check if in PATH
echo ""
if echo "$PATH" | grep -q "$INSTALL_DIR"; then
    echo -e "${GREEN}✓${NC} $INSTALL_DIR is in your PATH"
else
    echo -e "${YELLOW}⚠${NC}  $INSTALL_DIR is not in your PATH"
    echo ""
    echo "Quick fix - run this command:"
    echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
    echo ""
    echo "To make it permanent, add to your shell config:"
    if [ -n "$ZSH_VERSION" ]; then
        echo "  echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.zshrc"
        echo "  source ~/.zshrc"
    else
        echo "  echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.bashrc"
        echo "  source ~/.bashrc"
    fi
fi

# Success message
echo ""
echo "╔══════════════════════════════════════════╗"
echo "║  ✓ Installation Complete!                ║"
echo "╚══════════════════════════════════════════╝"
echo ""
echo "Quick Start:"
echo "  1. Run the builder:  tangent"
echo "  2. Use in Go code:   import \"github.com/$REPO/pkg/characters\""
echo ""
echo "Examples:"
echo "  tangent                    # Start visual builder"
echo "  tangent --version          # Check version"
echo ""
echo "Documentation:"
echo "  https://github.com/$REPO"
echo ""
echo "Need help? https://github.com/$REPO/issues"

