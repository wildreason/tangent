#!/bin/bash
# Build script for Tangent - injects version at build time

set -e

# Get version from git tag or use dev
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

echo "Building Tangent..."
echo "  Version: $VERSION"
echo "  Commit:  $COMMIT"
echo "  Date:    $DATE"

# Build with version injection
go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE" \
  -o tangent ./cmd/tangent

echo "✓ Built: tangent"
echo "  Run: ./tangent version"
