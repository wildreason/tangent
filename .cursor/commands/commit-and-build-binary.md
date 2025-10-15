## Release workflow

```bash
# 1. Ensure clean state
git status

# 2. Update CHANGELOG.md

# 3. Create and push release tag
make release
# Enter version: v0.1.2

# 4. GitHub Actions automatically:
#    - Builds for all platforms
#    - Creates GitHub Release
#    - Uploads binaries

# 5. Verify release
#    Visit: https://github.com/wildreason/tangent/releases
```

## Release Process

### Development Builds
```bash
make build
./tangent version
# Shows: v0.1.1-11-gc80a59a-dirty (dev build)
```

### Production Releases
```bash
# 1. Ensure clean state
git status

# 2. Update CHANGELOG.md

# 3. Create and push release tag
make release
# Enter version: v0.1.2

# 4. GitHub Actions automatically:
#    - Builds for all platforms
#    - Creates GitHub Release
#    - Uploads binaries

# 5. Verify release
#    Visit: https://github.com/wildreason/tangent/releases
```

### Version Identification

- **Dev:** `v0.1.1-11-gc80a59a-dirty` (has suffix)
- **Prod:** `v0.1.2` (clean tag only)