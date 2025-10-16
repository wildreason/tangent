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
echo "║  Terminal Avatars for AI Agents          ║"
echo "╚══════════════════════════════════════════╝"
echo ""

# Check for existing installation
EXISTING_VERSION=""
FORCE_INSTALL=false

if command -v tangent &> /dev/null; then
    EXISTING_VERSION=$(tangent --version 2>/dev/null | awk '{print $2}' || echo "unknown")
    echo -e "${YELLOW}⚠${NC}  Tangent is already installed: v$EXISTING_VERSION"
fi

# Get latest available version early
echo "Checking latest version..."
LATEST_TAG=$(curl -fsSL https://api.github.com/repos/$REPO/releases 2>/dev/null | grep -o '"tag_name": *"[^"]*"' | head -1 | sed 's/"tag_name": *"\(.*\)"/\1/')
if [ -z "$LATEST_TAG" ]; then
    echo -e "${RED}✗ Failed to fetch latest version from GitHub${NC}"
    exit 1
fi
LATEST_VERSION=$(echo "$LATEST_TAG" | sed 's/^v//')

# Decide what to do
if [ -n "$EXISTING_VERSION" ] && [ "$EXISTING_VERSION" != "unknown" ]; then
    if [ "$EXISTING_VERSION" = "$LATEST_VERSION" ]; then
        echo -e "${GREEN}✓${NC} You have the latest version ($LATEST_VERSION)"
        echo ""
        echo "To reinstall anyway, run:"
        echo "  curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | FORCE=1 bash"
        exit 0
    else
        echo -e "${YELLOW}→${NC} Upgrading from v$EXISTING_VERSION to v$LATEST_VERSION"
        FORCE_INSTALL=true
    fi
else
    echo -e "${GREEN}→${NC} Installing Tangent v$LATEST_VERSION"
fi

# Check if FORCE environment variable is set
if [ "$FORCE" = "1" ]; then
    echo -e "${YELLOW}→${NC} Force install requested"
    FORCE_INSTALL=true
fi

echo ""

# Check if Go is installed (optional for CLI users)
GO_AVAILABLE=false
if command -v go &> /dev/null; then
    echo -e "${GREEN}✓${NC} Go detected: $(go version | awk '{print $3}')"
    GO_AVAILABLE=true
else
    echo -e "${YELLOW}⚠${NC}  Go not detected (optional - CLI will work without it)"
fi

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

# Use version info we already fetched
echo ""
if [ "$VERSION" = "latest" ]; then
    RELEASE_TAG="$LATEST_TAG"
    VERSION_NUM="$LATEST_VERSION"
else
    # User specified a specific version
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

# Note: Go package developers can import the package in their code:
# import "github.com/wildreason/tangent/pkg/characters"
# No separate installation needed - go mod will handle it automatically

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

# Concise success output (5 lines)
echo ""
VERSION_TAG="$RELEASE_TAG"
echo "Tangent Installer — Terminal Avatars for AI Agents"
echo "Installing ${VERSION_TAG} to ${INSTALL_DIR} ..."
echo "Downloading and extracting… ✓"
echo "Installed: ${INSTALL_DIR}/tangent"
echo "Next: tangent browse (discover avatars for your AI agent)"

