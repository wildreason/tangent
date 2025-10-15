#!/bin/bash
# Release validation script - ensures CHANGELOG is updated before release

set -e

VERSION="$1"

if [ -z "$VERSION" ]; then
    echo "❌ Error: Version required"
    echo "Usage: $0 <version>"
    exit 1
fi

echo "🔍 Validating release: $VERSION"

# Check if CHANGELOG.md exists
if [ ! -f "CHANGELOG.md" ]; then
    echo "❌ Error: CHANGELOG.md not found"
    exit 1
fi

# Check if CHANGELOG has been updated for this version
# Handle both v0.1.0 and 0.1.0 formats
VERSION_NO_V="${VERSION#v}"  # Remove 'v' prefix if present
VERSION_WITH_V="v$VERSION_NO_V"  # Add 'v' prefix

if ! grep -q "## \[$VERSION\]" CHANGELOG.md && ! grep -q "## \[$VERSION_NO_V\]" CHANGELOG.md; then
    echo "❌ Error: CHANGELOG.md has not been updated for version $VERSION"
    echo ""
    echo "Please add an entry like:"
    echo "## [$VERSION_NO_V] - $(date +%Y-%m-%d)"
    echo ""
    echo "To CHANGELOG.md before creating the release."
    exit 1
fi

# Check if there are uncommitted changes to CHANGELOG.md
if ! git diff --quiet CHANGELOG.md; then
    echo "❌ Error: CHANGELOG.md has uncommitted changes"
    echo ""
    echo "Please commit your CHANGELOG.md changes before creating the release:"
    echo "  git add CHANGELOG.md"
    echo "  git commit -m \"Update CHANGELOG for $VERSION\""
    exit 1
fi

# Check if working directory is clean (except for untracked files)
if ! git diff --quiet HEAD; then
    echo "❌ Error: Working directory has uncommitted changes"
    echo ""
    echo "Please commit all changes before creating the release:"
    echo "  git add ."
    echo "  git commit -m \"Prepare for $VERSION release\""
    exit 1
fi

# Check if tag already exists locally
if git tag -l | grep -q "^$VERSION$"; then
    echo "❌ Error: Tag $VERSION already exists locally"
    echo ""
    echo "Existing tags:"
    git tag -l | sort -V
    exit 1
fi

# Check if tag exists on remote (GitHub)
if git ls-remote --tags origin 2>/dev/null | grep -q "refs/tags/$VERSION$"; then
    if [ "$FORCE_RELEASE" != "1" ]; then
        echo "❌ Error: Tag $VERSION already exists on GitHub"
        echo ""
        echo "This usually means a release was already created."
        echo "If you need to re-release, you must:"
        echo "  1. Delete the remote tag: git push origin --delete $VERSION"
        echo "  2. Delete the GitHub release manually at:"
        echo "     https://github.com/wildreason/tangent/releases"
        echo "  3. Run make release again"
        echo ""
        echo "Or use: FORCE_RELEASE=1 make release (not recommended)"
        exit 1
    else
        echo "⚠️  Warning: Tag $VERSION exists on GitHub but FORCE_RELEASE=1"
        echo "⚠️  This may cause conflicts. Proceed with caution."
    fi
fi

echo "✅ CHANGELOG.md updated for $VERSION"
echo "✅ Working directory is clean"
echo "✅ Tag $VERSION does not exist locally"
echo "✅ Tag $VERSION does not exist on GitHub"
echo "✅ Ready to create release!"
