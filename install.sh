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

# Get release info and determine download URL
echo ""
echo "Fetching release information..."

if [ "$VERSION" = "latest" ]; then
    # Try to get latest release (including pre-releases)
    RELEASE_TAG=$(curl -fsSL https://api.github.com/repos/$REPO/releases 2>/dev/null | grep -o '"tag_name": *"[^"]*"' | head -1 | sed 's/"tag_name": *"\(.*\)"/\1/')
    if [ -z "$RELEASE_TAG" ]; then
        echo -e "${RED}✗ Failed to fetch latest release${NC}"
        echo ""
        echo "Alternative: Install via Go"
        echo "  go install github.com/$REPO/cmd/tangent@latest"
        exit 1
    fi
    VERSION_NUM=$(echo "$RELEASE_TAG" | sed 's/^v//')
else
    RELEASE_TAG="$VERSION"
    VERSION_NUM=$(echo "$VERSION" | sed 's/^v//')
fi

# Determine file extension based on OS
case $OS in
    windows)
        FILE_EXT="zip"
        ;;
    *)
        FILE_EXT="tar.gz"
        ;;
esac

# Construct platform string (capitalize first letter for release assets)
OS_CAPITALIZED="$(tr '[:lower:]' '[:upper:]' <<< ${OS:0:1})${OS:1}"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$RELEASE_TAG/tangent_${VERSION_NUM}_${OS_CAPITALIZED}_${ARCH}.${FILE_EXT}"

echo "Downloading tangent $RELEASE_TAG..."
echo "  From: $DOWNLOAD_URL"

# Download archive
TEMP_DIR=$(mktemp -d)
TEMP_FILE="$TEMP_DIR/tangent.${FILE_EXT}"
if ! curl -fsSL -o "$TEMP_FILE" "$DOWNLOAD_URL" 2>/dev/null; then
    echo -e "${RED}✗ Failed to download${NC}"
    echo "  URL: $DOWNLOAD_URL"
    echo ""
    echo "Alternative: Install via Go"
    echo "  go install github.com/$REPO/cmd/tangent@latest"
    rm -rf "$TEMP_DIR"
    exit 1
fi
echo -e "${GREEN}✓${NC} Downloaded tangent archive"

# Extract binary
echo "Extracting binary..."
cd "$TEMP_DIR"
if [ "$FILE_EXT" = "zip" ]; then
    if ! unzip -q "$TEMP_FILE"; then
        echo -e "${RED}✗ Failed to extract archive${NC}"
        rm -rf "$TEMP_DIR"
        exit 1
    fi
else
    if ! tar -xzf "$TEMP_FILE"; then
        echo -e "${RED}✗ Failed to extract archive${NC}"
        rm -rf "$TEMP_DIR"
        exit 1
    fi
fi

# Find and install binary
BINARY_NAME="tangent"
if [ "$OS" = "windows" ]; then
    BINARY_NAME="tangent.exe"
fi

if [ ! -f "$BINARY_NAME" ]; then
    echo -e "${RED}✗ Binary not found in archive${NC}"
    rm -rf "$TEMP_DIR"
    exit 1
fi

# Install binary
mv "$BINARY_NAME" "$INSTALL_DIR/tangent"
chmod +x "$INSTALL_DIR/tangent"
rm -rf "$TEMP_DIR"
echo -e "${GREEN}✓${NC} Installed to $INSTALL_DIR/tangent"

# Install Go package
echo ""
echo "Installing Go package..."
if go install "github.com/$REPO/cmd/tangent@$VERSION" 2>/dev/null; then
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

