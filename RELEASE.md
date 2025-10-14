# Release Guide

This document describes the release process for Tangent, covering both development and production releases.

---

## Version Identification

### Development Build
```bash
tangent v0.1.1-11-gc80a59a-dirty (commit: c80a59a, built: 2025-10-14T18:11:26Z)
```

**Format:** `v{tag}-{commits}-g{hash}[-dirty]`
- `v0.1.1` - Last release tag
- `11` - Commits ahead of tag
- `gc80a59a` - Current commit hash
- `dirty` - Uncommitted changes (if present)

**Indicates:** Work in progress, not for production

### Production Release
```bash
tangent v0.1.2 (commit: abc1234, built: 2025-10-14T18:30:00Z)
```

**Format:** `v{tag}` (clean, no suffix)

**Indicates:** Official release, production-ready

---

## Development Workflow

### Building Dev Versions

```bash
# Clean build with version tracking
make clean
make build

# Verify version
./tangent version
# Output: tangent v0.1.1-11-gc80a59a-dirty (...)
```

**Version Source:** `git describe --tags --always --dirty`

### Testing Dev Builds

```bash
# Run tests
make test

# Test CLI
./tangent

# Test library
go run examples/agent_states.go
```

### When to Build Dev Versions

- During active development
- Before committing changes
- Testing new features
- Debugging issues

---

## Production Release Workflow

### Prerequisites

1. **Clean working directory**
   ```bash
   git status  # Should show no uncommitted changes
   ```

2. **All tests passing**
   ```bash
   make test
   ```

3. **CHANGELOG.md updated**
   - Document all changes since last release
   - Follow existing format

4. **Version number decided**
   - Follow semantic versioning: `vMAJOR.MINOR.PATCH`
   - Use suffixes for pre-releases: `-alpha.N`, `-beta.N`, `-rc.N`

### Release Steps

#### 1. Prepare Release

```bash
# Ensure clean state
git status

# Update CHANGELOG.md
# Add new version section with changes

# Commit final changes
git add CHANGELOG.md
git commit -m "chore: prepare for v0.1.2 release"
git push origin main
```

#### 2. Create Release Tag

```bash
# Option A: Using Makefile (interactive)
make release
# Enter version when prompted: v0.1.2

# Option B: Manual
git tag v0.1.2
git push origin v0.1.2
```

#### 3. Automated Build & Release

**GitHub Actions automatically:**
1. Detects tag push
2. Runs GoReleaser
3. Builds for all platforms (Linux, macOS, Windows)
4. Creates GitHub Release
5. Uploads binaries and checksums
6. Generates changelog

#### 4. Verify Release

```bash
# Check GitHub Releases page
# https://github.com/wildreason/tangent/releases

# Download and test binary
curl -L https://github.com/wildreason/tangent/releases/download/v0.1.2/tangent_v0.1.2_Darwin_arm64.tar.gz -o tangent.tar.gz
tar -xzf tangent.tar.gz
./tangent version
# Should show: tangent v0.1.2 (commit: ..., built: ...)
```

---

## Version Numbering

### Semantic Versioning

**Format:** `vMAJOR.MINOR.PATCH[-PRERELEASE]`

- **MAJOR** - Breaking changes
- **MINOR** - New features (backward compatible)
- **PATCH** - Bug fixes (backward compatible)
- **PRERELEASE** - Pre-release identifier

### Pre-release Identifiers

- **alpha** - Early development, unstable
  - Example: `v0.1.0-alpha.1`, `v0.1.0-alpha.2`
  - Use for: Initial features, experimental changes

- **beta** - Feature complete, testing phase
  - Example: `v0.1.0-beta.1`, `v0.1.0-beta.2`
  - Use for: Feature-complete versions needing testing

- **rc** - Release candidate, final testing
  - Example: `v0.1.0-rc.1`, `v0.1.0-rc.2`
  - Use for: Production-ready candidates

### Examples

```
v0.1.0-alpha.1  → First alpha release
v0.1.0-alpha.2  → Second alpha release
v0.1.0-beta.1   → First beta release
v0.1.0-rc.1     → First release candidate
v0.1.0          → Stable release
v0.1.1          → Patch release
v0.2.0          → Minor version bump
v1.0.0          → Major version bump
```

---

## Release Automation

### GoReleaser Configuration

**File:** `.goreleaser.yaml`

**Key Features:**
- Multi-platform builds (Linux, macOS, Windows)
- Version injection via ldflags
- Checksum generation
- GitHub Release creation
- Archive generation (tar.gz, zip)

### GitHub Actions Workflow

**File:** `.github/workflows/release.yml`

**Trigger:** Push to tag matching `v*`

**Process:**
1. Checkout code
2. Setup Go environment
3. Run GoReleaser
4. Upload artifacts to GitHub Release

---

## Quick Reference

### Development

```bash
# Build dev version
make build

# Check version
./tangent version
# Output: v0.1.1-N-g<hash>[-dirty]
```

### Production

```bash
# Create release
make release
# Enter: v0.1.2

# Automated:
# - Tag created
# - Tag pushed
# - GitHub Actions triggered
# - Binaries built
# - Release published
```

### Version Check

| Version String | Type | Status |
|---------------|------|--------|
| `v0.1.2` | Production | ✓ Ready |
| `v0.1.2-dirty` | Dev | ✗ Uncommitted changes |
| `v0.1.2-5-gabc1234` | Dev | ✗ 5 commits ahead |
| `v0.1.2-5-gabc1234-dirty` | Dev | ✗ Ahead + uncommitted |

---

## Troubleshooting

### "Version shows 'dev'"

**Problem:** Binary built without version injection

**Solution:**
```bash
# Use make build instead of go build
make clean
make build
```

### "Tag already exists"

**Problem:** Trying to create duplicate tag

**Solution:**
```bash
# Delete local tag
git tag -d v0.1.2

# Delete remote tag (if pushed)
git push origin :refs/tags/v0.1.2

# Create new tag
git tag v0.1.2
git push origin v0.1.2
```

### "GitHub Actions failed"

**Problem:** Build or test failures

**Solution:**
1. Check GitHub Actions logs
2. Fix issues locally
3. Delete failed tag
4. Create new tag after fixes

---

## Best Practices

### Before Release

- [ ] All tests passing
- [ ] CHANGELOG.md updated
- [ ] Version number decided
- [ ] Clean working directory
- [ ] Code reviewed
- [ ] Documentation updated

### After Release

- [ ] Verify GitHub Release created
- [ ] Test downloaded binaries
- [ ] Announce release (if applicable)
- [ ] Update documentation links
- [ ] Monitor for issues

### Version Strategy

- **Alpha** - Weekly or feature-based releases
- **Beta** - Bi-weekly releases for testing
- **RC** - Final testing before stable
- **Stable** - Monthly or milestone-based

---

## Additional Resources

- **GoReleaser Docs:** https://goreleaser.com
- **Semantic Versioning:** https://semver.org
- **GitHub Releases:** https://docs.github.com/en/repositories/releasing-projects-on-github

---

**Last Updated:** 2025-10-14

