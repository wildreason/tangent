# Distribution Plan: Tangent v0.0.1

**Repository**: https://github.com/wildreason/tangent  
**Date**: October 7, 2025

---

## Pre-Distribution Checklist

### ✓ Code Ready
- [x] Zero external dependencies
- [x] Examples tested and working
- [x] Documentation streamlined (2 essential docs)
- [x] README comprehensive
- [x] Clean repository structure
- [x] All changes committed

### ✓ Documentation Audit
- [x] Removed 9 distracting internal docs
- [x] Kept 2 essential user docs (PATTERN_GUIDE, LIBRARY)
- [x] Enhanced README as main entry point
- [x] Examples provide hands-on learning

### Repository Structure (Public-Facing)
```
tangent/
├── cmd/tangent/        ← CLI builder
├── pkg/characters/     ← Go package
├── examples/           ← Usage demos (2)
├── docs/              ← Essential docs (2)
├── tools/             ← Build scripts
├── README.md          ← Comprehensive guide
├── CHANGELOG.md
├── VERSION.md
└── go.mod
```

---

## Distribution Steps

### 1. Merge to Main Branch

```bash
# From builder branch
git checkout main
git merge builder --no-ff -m "Release v0.1.0: Add Tangent CLI builder with animation and export features"
git push origin main
```

### 2. Tag Release

```bash
git tag -a v0.1.0 -m "v0.1.0: Tangent CLI Builder

Features:
- Interactive visual character designer
- Live animation preview
- Multi-frame session management
- Export to code or save to file
- Duplicate frame for easy animation
- Zero external dependencies

Package Features:
- Pattern-based character creation
- Built-in animation engine
- Pre-built character library
- Fluent builder API"

git push origin v0.1.0
```

### 3. Push to GitHub

```bash
# Add remote (already created via gh cli)
git remote add origin https://github.com/wildreason/tangent.git

# Push main branch and tags
git push -u origin main
git push origin --tags
```

### 4. Create GitHub Release

Via GitHub web interface:
- Go to https://github.com/wildreason/tangent/releases
- Click "Create a new release"
- Choose tag: v0.1.0
- Title: "Tangent v0.1.0 - Terminal Agent Designer"
- Description: (see below)
- Attach binaries:
  - `tangent-darwin-amd64`
  - `tangent-darwin-arm64`
  - `tangent-linux-amd64`
  - `tangent-linux-arm64`

#### Release Description Template

```markdown
# Tangent v0.1.0 - Terminal Agent Designer

Design animated terminal characters using Unicode block elements.

## 🚀 Quick Start

### Install via Binary
Download for your platform:
- macOS (Intel): `tangent-darwin-amd64`
- macOS (Apple Silicon): `tangent-darwin-arm64`
- Linux (x86_64): `tangent-linux-amd64`
- Linux (ARM): `tangent-linux-arm64`

Make it executable and run:
```bash
chmod +x tangent-*
./tangent-*
```

### Build from Source
```bash
git clone https://github.com/wildreason/tangent.git
cd tangent/cmd/tangent
go build -o tangent .
./tangent
```

## ✨ Features

### Tangent CLI Builder
- Interactive visual design
- Live animation preview
- Multi-frame session management
- Auto-save & resume projects
- Export Go code or save to file
- Duplicate frames for animation

### Go Package
- Zero external dependencies
- Pattern-based API
- Built-in animation engine
- Pre-built character library

## 📚 Documentation
- [README](https://github.com/wildreason/tangent/blob/main/README.md)
- [Pattern Guide](https://github.com/wildreason/tangent/blob/main/docs/PATTERN_GUIDE.md)
- [Library Characters](https://github.com/wildreason/tangent/blob/main/docs/LIBRARY.md)
- [Examples](https://github.com/wildreason/tangent/tree/main/examples)

## 🐛 Known Issues
None reported

## 📝 Changelog
See [CHANGELOG.md](https://github.com/wildreason/tangent/blob/main/CHANGELOG.md)
```

---

## Build Binaries for Release

### macOS (Intel)
```bash
cd cmd/tangent
GOOS=darwin GOARCH=amd64 go build -o tangent-darwin-amd64 .
```

### macOS (Apple Silicon)
```bash
GOOS=darwin GOARCH=arm64 go build -o tangent-darwin-arm64 .
```

### Linux (x86_64)
```bash
GOOS=linux GOARCH=amd64 go build -o tangent-linux-amd64 .
```

### Linux (ARM)
```bash
GOOS=linux GOARCH=arm64 go build -o tangent-linux-arm64 .
```

---

## Post-Distribution

### 1. Update Website
- Add project page at wildreason.com/tangent
- Link to GitHub repository
- Embed demo GIF/video

### 2. Announce
- Twitter/X
- Reddit (r/golang, r/commandline)
- Hacker News
- Dev.to blog post

### 3. Monitor
- Watch GitHub issues
- Respond to questions
- Collect feedback

---

## Package Usage (for Go developers)

### Direct import (after push)
```bash
go get github.com/wildreason/tangent/pkg/characters
```

### In go.mod
```go
require github.com/wildreason/tangent v0.1.0
```

---

## Marketing Points

### For Users
- "Design terminal agents visually, no code required"
- "Zero external dependencies, works anywhere"
- "Export ready-to-use Go code"

### For Developers
- "Lightweight terminal animation library"
- "Pattern-based character definition"
- "Built-in animation engine"
- "Zero external dependencies"

---

## Success Metrics

### Week 1
- [ ] 50+ stars on GitHub
- [ ] 5+ watchers
- [ ] 3+ forks

### Month 1
- [ ] 100+ stars
- [ ] 10+ issues/PRs
- [ ] Featured in Go Weekly

---

## Repository Settings

### GitHub Settings
- [x] Public repository
- [ ] Enable Issues
- [ ] Enable Discussions
- [ ] Add topics: `golang`, `terminal`, `cli`, `animation`, `unicode`, `block-elements`
- [ ] Add description: "Design terminal agents with Unicode block characters"
- [ ] Add website: https://wildreason.com

### Branch Protection (main)
- Require pull request reviews
- Require status checks to pass
- No force push

---

## Next Steps (Execute in Order)

1. ✅ **DONE**: Clean up documentation
2. ✅ **DONE**: Commit changes
3. ⏭️ **NEXT**: Merge builder → main
4. ⏭️ Build release binaries
5. ⏭️ Create GitHub release
6. ⏭️ Push to GitHub
7. ⏭️ Configure repository settings
8. ⏭️ Announce

---

**Ready to distribute!** 🚀

© 2025 Wildreason, Inc

