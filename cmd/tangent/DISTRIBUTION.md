# Tangent Distribution Quick Guide

## ✓ What You Have

A standalone CLI binary called **`tangent`** (Terminal Agent Designer) that:
- Uses your characters package internally
- Provides interactive character building
- Saves sessions to `~/.tangent/`
- Exports Go code for users
- Works cross-platform (macOS, Linux, Windows)
- Size: ~3 MB per binary

## ▢ Three Ways to Distribute

### 1. Direct Binary Download (Easiest)

**Who**: Beta testers, early users, non-technical folks

**How**:
```bash
# Build all platforms
cd cmd/tangent
./build.sh

# Upload to:
# - AWS S3 / CloudFront
# - GitHub Releases
# - Your website
```

**User runs**:
```bash
curl -L https://your-domain.com/tangent-macos-arm64 -o tangent
chmod +x tangent
./tangent
```

**Pros**: No dependencies, works immediately, easiest for users
**Cons**: Manual download, PATH setup required

---

### 2. One-Line Installer (Developer Friendly)

**Who**: Developers, power users

**Create** `install-tangent.sh`:
```bash
#!/bin/bash
curl -sSL https://releases.wildreason.com/tangent-$(uname -s)-$(uname -m) \
  -o ~/.local/bin/tangent
chmod +x ~/.local/bin/tangent
export PATH="$HOME/.local/bin:$PATH"
tangent
```

**User runs**:
```bash
curl -sSL https://wildreason.com/install-tangent.sh | bash
```

**Pros**: One command, auto-detects platform, clean install
**Cons**: Requires curl, PATH config

---

### 3. Homebrew (macOS Premium)

**Who**: macOS developers

**Setup**:
1. Create `homebrew-tools` GitHub repo
2. Add `Formula/tangent.rb`:
```ruby
class Tangent < Formula
  desc "Design terminal agents with character"
  homepage "https://wildreason.com/tangent"
  url "https://github.com/wildreason/characters/releases/download/v0.0.1/tangent-macos-arm64"
  sha256 "..."
  
  def install
    bin.install "tangent-macos-arm64" => "tangent"
  end
end
```

**User runs**:
```bash
brew tap wildreason/tools
brew install tangent
```

**Pros**: Standard macOS workflow, auto-updates
**Cons**: macOS only, requires maintenance

---

## ◆ Recommended Strategy

### Phase 1: Private Beta (Now)
✓ Use **Direct Binary Download**
- Send download links via email
- 50-100 users
- Collect feedback

### Phase 2: Public Launch
✓ Use **Direct Binary + Installer**
- Post on Product Hunt, HackerNews
- GitHub Releases page
- 1,000-5,000 users

### Phase 3: Scale
✓ Add **Homebrew + Docker**
- Full package manager support
- Enterprise distribution
- Unlimited users

---

## 🚀 Launch Checklist

- [x] Build `tangent` binary
- [x] Cross-platform build script (`build.sh`)
- [x] User README (`cmd/tangent/README.md`)
- [ ] Host binaries (S3/GitHub/CDN)
- [ ] Create install script
- [ ] Landing page (wildreason.com/tangent)
- [ ] Demo GIF/video
- [ ] Social media posts
- [ ] Analytics tracking

---

## 📦 Quick Test

```bash
# Build it
cd cmd/tangent
go build -o tangent .

# Run it
./tangent

# You should see:
# ╔══════════════════════════════════════════╗
# ║  TANGENT - Terminal Agent Designer      ║
# ║  Design characters for your CLI agents  ║
# ║  v0.0.1                                  ║
# ╚══════════════════════════════════════════╝
```

---

## 💡 Key Selling Points

1. **"No coding required"** - Visual builder, menu-driven
2. **"Works immediately"** - Single binary, no install
3. **"Resume anytime"** - Auto-saves to ~/.tangent/
4. **"Export Go code"** - Copy-paste into your project
5. **"Cross-platform"** - macOS, Linux, Windows

---

## 📊 Success Metrics

- Downloads per week
- Characters created (track sessions)
- Code exports (engagement)
- 7-day retention
- Social shares

---

Made with ◆ by Wildreason, Inc

