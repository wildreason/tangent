#!/bin/bash
# Build script for Tangent - Terminal Agent Designer
# Supports -m flag for git commit message

set -e

VERSION="v0.0.1"
BUILD_DIR="dist"
COMMIT_MSG=""

# Parse flags
while getopts "m:" opt; do
  case $opt in
    m) COMMIT_MSG="$OPTARG";;
    *) echo "Usage: $0 [-m 'commit message']" && exit 1;;
  esac
done

echo "╔══════════════════════════════════════════╗"
echo "║  Building Tangent $VERSION                 ║"
echo "╚══════════════════════════════════════════╝"
echo ""

# Create dist directory
mkdir -p "$BUILD_DIR"

# Build for multiple platforms
echo "◢ Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o "$BUILD_DIR/tangent-macos-arm64" .

echo "◢ Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o "$BUILD_DIR/tangent-macos-amd64" .

echo "◢ Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "$BUILD_DIR/tangent-linux-amd64" .

echo "◢ Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "$BUILD_DIR/tangent-windows-amd64.exe" .

echo ""
echo "✓ Build complete!"
echo ""
echo "Binaries created in $BUILD_DIR/:"
ls -lh "$BUILD_DIR/"
echo ""

# Calculate file sizes
echo "◢ Binary sizes:"
du -h "$BUILD_DIR"/* | awk '{print "  " $2 ": " $1}'
echo ""

# Optional: Git commit
if [ -n "$COMMIT_MSG" ]; then
  echo "◢ Creating git commit..."
  git add "$BUILD_DIR/"
  git commit -m "$COMMIT_MSG"
  echo "✓ Committed: $COMMIT_MSG"
  echo ""
fi

echo "╔══════════════════════════════════════════╗"
echo "║  Ready to distribute!                    ║"
echo "╚══════════════════════════════════════════╝"
echo ""
echo "Test locally:"
echo "  ./$BUILD_DIR/tangent-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')"
echo ""
echo "Install locally:"
echo "  sudo cp $BUILD_DIR/tangent-* /usr/local/bin/tangent"

