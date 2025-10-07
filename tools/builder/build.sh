#!/bin/bash
# Build script for character builder
# Supports -m flag for commit message

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

echo "◢ Building Character Builder $VERSION..."

# Create dist directory
mkdir -p "$BUILD_DIR"

# Build for multiple platforms
echo "◢ Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o "$BUILD_DIR/character-builder-macos-arm64" .

echo "◢ Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o "$BUILD_DIR/character-builder-macos-amd64" .

echo "◢ Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o "$BUILD_DIR/character-builder-linux-amd64" .

echo "◢ Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o "$BUILD_DIR/character-builder-windows-amd64.exe" .

echo ""
echo "✓ Build complete!"
echo ""
echo "Binaries created in $BUILD_DIR/:"
ls -lh "$BUILD_DIR/"

# Optional: Git commit
if [ -n "$COMMIT_MSG" ]; then
  echo ""
  echo "◢ Creating git commit..."
  git add "$BUILD_DIR/"
  git commit -m "$COMMIT_MSG"
  echo "✓ Committed: $COMMIT_MSG"
fi

