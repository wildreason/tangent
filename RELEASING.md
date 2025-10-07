# Release Guide

This document describes how to create new releases of Tangent using GoReleaser.

## Prerequisites

- GoReleaser installed (optional for manual releases)
- GitHub token with repo permissions
- Write access to the repository

## Automated Release (Recommended)

Releases are automated via GitHub Actions when you push a tag.

### 1. Update Version Files

```bash
# Update CHANGELOG.md with new version changes
vim CHANGELOG.md

# Commit changes
git add CHANGELOG.md
git commit -m "docs: prepare for v0.2.0 release"
git push origin main
```

### 2. Create and Push Tag

```bash
# Create annotated tag
git tag -a v0.2.0 -m "Release v0.2.0

Features:
- Feature 1
- Feature 2

Fixes:
- Fix 1
- Fix 2
"

# Push tag to trigger release
git push origin v0.2.0
```

### 3. Monitor Release

GitHub Actions will automatically:
- Build binaries for all platforms (Linux, macOS, Windows)
- Create checksums
- Generate changelog
- Create GitHub release with artifacts
- Upload binaries

Watch progress at:
https://github.com/wildreason/tangent/actions

## Manual Release (If Needed)

### Install GoReleaser

```bash
# macOS
brew install goreleaser

# Linux
go install github.com/goreleaser/goreleaser@latest
```

### Test Release

```bash
# Dry run (no publishing)
goreleaser release --snapshot --clean

# Check dist/ folder for binaries
ls -la dist/
```

### Create Release

```bash
# Set GitHub token
export GITHUB_TOKEN="your_github_token"

# Create tag
git tag -a v0.2.0 -m "Release v0.2.0"

# Run GoReleaser
goreleaser release --clean

# Push tag
git push origin v0.2.0
```

## Release Checklist

### Before Release

- [ ] All tests pass: `go test ./...`
- [ ] Examples work: `go run examples/demo/main.go`
- [ ] Tangent builds: `cd cmd/tangent && go build`
- [ ] CHANGELOG.md updated
- [ ] VERSION.md updated (optional)
- [ ] No uncommitted changes
- [ ] On main branch

### Version Naming

Follow [Semantic Versioning](https://semver.org/):

- **v0.1.0-alpha** - Alpha release (current)
- **v0.1.0-beta** - Beta release
- **v0.1.0-rc.1** - Release candidate
- **v0.1.0** - Stable release
- **v0.2.0** - Minor version (new features)
- **v0.2.1** - Patch version (bug fixes)
- **v1.0.0** - Major stable release

### After Release

- [ ] Verify release on GitHub
- [ ] Download and test binaries
- [ ] Update documentation if needed
- [ ] Announce release (optional)
- [ ] Close milestone (if using)

## What Gets Built

GoReleaser builds binaries for:

### Platforms
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Archives
- `.tar.gz` for Linux/macOS
- `.zip` for Windows
- Includes: LICENSE, README.md, CHANGELOG.md, docs/

### Artifacts
- Binaries: `tangent-{os}-{arch}`
- Archives: `tangent_{version}_{OS}_{arch}.{ext}`
- Checksums: `checksums.txt`

## GoReleaser Configuration

Configuration is in `.goreleaser.yaml`:

```yaml
builds:
  - binary: tangent
    main: ./cmd/tangent
    goos: [linux, darwin, windows]
    goarch: [amd64, arm64]
    ldflags:
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}
```

Version info is injected at build time and shown in the banner.

## Troubleshooting

### Release Failed

Check GitHub Actions logs:
```
https://github.com/wildreason/tangent/actions
```

Common issues:
- Missing GITHUB_TOKEN permissions
- Invalid tag format (must start with 'v')
- Uncommitted changes
- Build failures

### Re-release Same Version

```bash
# Delete local tag
git tag -d v0.2.0

# Delete remote tag
git push origin :refs/tags/v0.2.0

# Delete GitHub release (via web UI)

# Recreate tag and push
git tag -a v0.2.0 -m "..."
git push origin v0.2.0
```

### Test GoReleaser Config

```bash
# Validate configuration
goreleaser check

# Build snapshot without release
goreleaser build --snapshot --clean
```

## Version Variables

The built binary includes version info:

```go
// cmd/tangent/main.go
var (
    version = "dev"      // Set by GoReleaser
    commit  = "none"     // Git commit hash
    date    = "unknown"  // Build date
)
```

Shown in the banner:
```
╔══════════════════════════════════════════╗
║  TANGENT - Terminal Agent Designer      ║
║  Design characters for your CLI agents  ║
║  v0.1.0-alpha                           ║
╚══════════════════════════════════════════╝
```

## GitHub Release Template

Automatically generated from `.goreleaser.yaml`:

```markdown
## Tangent v0.2.0 - Terminal Agent Designer

Design animated terminal characters using Unicode block elements.

### Installation
Download the binary for your platform below...

### Quick Start
`./tangent`

### Changelog
- Feature 1
- Feature 2
- Bug fix 1

### Documentation
- README
- Pattern Guide
- Library Characters
```

## Next Release

To prepare for the next release:

1. Create a new branch: `git checkout -b release/v0.2.0`
2. Update CHANGELOG.md
3. Test thoroughly
4. Merge to main
5. Tag and release

---

© 2025 Wildreason, Inc.

