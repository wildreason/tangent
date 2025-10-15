#!/bin/bash
# Pre-tag hook to ensure CHANGELOG is updated

set -e

TAG_NAME="$1"

if [ -z "$TAG_NAME" ]; then
    echo "❌ Error: Tag name required"
    exit 1
fi

echo "🔍 Pre-tag validation for: $TAG_NAME"

# Check if CHANGELOG.md exists
if [ ! -f "CHANGELOG.md" ]; then
    echo "❌ Error: CHANGELOG.md not found"
    exit 1
fi

# Check if CHANGELOG has been updated for this version
if ! grep -q "## \[$TAG_NAME\]" CHANGELOG.md; then
    echo "❌ Error: CHANGELOG.md has not been updated for version $TAG_NAME"
    echo ""
    echo "Please add an entry like:"
    echo "## [$TAG_NAME] - $(date +%Y-%m-%d)"
    echo ""
    echo "To CHANGELOG.md before creating the tag."
    exit 1
fi

echo "✅ CHANGELOG.md updated for $TAG_NAME"
echo "✅ Proceeding with tag creation"
