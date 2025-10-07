# Tangent

**Design terminal agents with Unicode block characters**

Two ways to use: **CLI Builder** (visual design) or **Go Package** (code directly).

---

## Choose Your Path

### 🎨 Path 1: Visual Design (Recommended for Beginners)

**Use the Tangent CLI builder** - No coding required, visual character design.

**Install:**
```bash
# Download binary from releases
# OR build from source:
git clone https://github.com/wildreason/tangent.git
cd tangent/cmd/tangent
go build -o tangent .
```

**Use:**
```bash
./tangent  # Start interactive builder
```

**What you get:**
- Visual character designer
- Live animation preview
- Export ready-to-use Go code
- Save characters for later

---

### 💻 Path 2: Go Package (For Developers)

**Use as a Go library** - Write code directly, full programmatic control.

**Install:**
```bash
go get github.com/wildreason/tangent/pkg/characters
```

**Use:**
```go
package main

import (
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    // Option A: Use library character
    alien, _ := characters.Library("alien")
    characters.Animate(os.Stdout, alien, 5, 3)
    
    // Option B: Create your own
    spec := characters.NewCharacterSpec("robot", 9, 3).
        AddFrame("idle", []string{
            "__R6FFF6L__",
            "_T6FFFFF5T_",
            "___11_22___",
        })
    
    robot, _ := spec.Build()
    characters.Animate(os.Stdout, robot, 5, 3)
}
```

---

## Which Should I Use?

| Scenario | Use This |
|----------|----------|
| 🎨 Want to design visually | **CLI Builder** (tangent) |
| 💻 Building a Go app | **Go Package** |
| 🚀 Quick prototyping | **CLI Builder** → export code |
| 🔧 Need programmatic control | **Go Package** |
| 📚 Want pre-built characters | **Both work!** |

**Best workflow**: Design in CLI builder → Export code → Use in your Go app

---

## Features

### Tangent CLI Builder
- ✓ Interactive visual design
- ✓ Live animation preview
- ✓ Multi-frame session management
- ✓ Auto-save & resume
- ✓ Export code or save to `.go` file
- ✓ Duplicate frames for easy animation

### Go Package
- ✓ Zero external dependencies
- ✓ Pattern-based character definition
- ✓ Built-in animation engine
- ✓ Pre-built character library
- ✓ Fluent builder API

---

## Pattern Codes

Single-character codes for block elements:

| Code | Block | Name | Code | Block | Name |
|------|-------|------|------|-------|------|
| `F` | █ | Full | `1` | ▘ | Upper Left |
| `T` | ▀ | Top Half | `2` | ▝ | Upper Right |
| `B` | ▄ | Bottom Half | `3` | ▖ | Lower Left |
| `L` | ▌ | Left Half | `4` | ▗ | Lower Right |
| `R` | ▐ | Right Half | `5` | ▛ | Three Quads (5) |
| `.` | ░ | Light Shade | `6` | ▜ | Three Quads (6) |
| `:` | ▒ | Medium Shade | `7` | ▙ | Three Quads (7) |
| `#` | ▓ | Dark Shade | `8` | ▟ | Three Quads (8) |
| `_` | (space) | Space | `X` | | Mirror marker |

**Full guide**: See [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md)

---

## Examples

### Run the Demo

```bash
go run examples/demo/main.go
```

Shows:
- Library characters in action
- Custom character creation
- Animation examples

### Example Character (created with Tangent)

See [`examples/tokyo/`](examples/tokyo/) - A character designed in Tangent, exported, and animated in Go.

---

## API Reference

### Create Characters

```go
// Pattern-based (recommended)
spec := characters.NewCharacterSpec("name", width, height).
    AddFrame("idle", []string{"pattern1", "pattern2", ...}).
    AddFrame("wave", []string{...})

char, err := spec.Build()
```

### Animate

```go
// Animate at 5 FPS for 3 loops
characters.Animate(os.Stdout, char, 5, 3)

// Show single frame
characters.ShowIdle(os.Stdout, char)
```

### Library Characters

```go
// Load pre-built character
alien, _ := characters.Library("alien")

// List available
names := characters.ListLibrary()  // ["alien"]
```

**Library reference**: See [`docs/LIBRARY.md`](docs/LIBRARY.md)

---

## Installation Guide

### 🎨 CLI Builder (Tangent)

**Option A: Download Binary** (Easiest)
```bash
# Go to Releases page and download for your platform
# https://github.com/wildreason/tangent/releases

# macOS/Linux
chmod +x tangent
./tangent
```

**Option B: Build from Source**
```bash
git clone https://github.com/wildreason/tangent.git
cd tangent/cmd/tangent
go build -o tangent .
./tangent

# Optional: Install to PATH
cp tangent ~/.local/bin/
```

---

### 💻 Go Package

**Add to your project:**
```bash
go get github.com/wildreason/tangent/pkg/characters
```

**Use in your code:**
```go
import "github.com/wildreason/tangent/pkg/characters"
```

That's it! Zero external dependencies.

---

## Dependencies

**Zero external dependencies** - Uses only Go standard library:
- `fmt`, `strings`, `time`, `io`, `sync`, `math`

Works with Go 1.21+

---

## Documentation

| Doc | Description |
|-----|-------------|
| [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md) | Pattern code reference |
| [`docs/LIBRARY.md`](docs/LIBRARY.md) | Pre-built characters |
| [`examples/`](examples/) | Usage examples |
| [`CHANGELOG.md`](CHANGELOG.md) | Version history |

---

## Project Structure

```
tangent/
├── cmd/tangent/        # CLI builder
├── pkg/characters/     # Go package
├── examples/           # Usage examples
├── docs/              # User documentation
└── tools/             # Build scripts
```

---

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run examples to verify
5. Submit a pull request

---

## Contributing

We welcome contributions! See [`CONTRIBUTING.md`](CONTRIBUTING.md) for guidelines.

## License

MIT License - © 2025 Wildreason, Inc  
See [`LICENSE`](LICENSE) for details.  
https://wildreason.com

---

## Links

- **Repository**: https://github.com/wildreason/tangent
- **Issues**: https://github.com/wildreason/tangent/issues
- **Website**: https://wildreason.com

---

**Built with ◆ by for AI agent buildersr**
© 2025 Wildreason, Inc - https://wildreason.com