# Tangent

**Design terminal agents with Unicode block characters**

A Go package and CLI tool for creating animated terminal characters using Unicode Block Elements (U+2580–U+259F).

---

## Quick Start

### 1. Install the CLI Tool

```bash
# Build from source
cd cmd/tangent
go build -o tangent .
cp tangent ~/.local/bin/  # Optional: add to PATH
```

### 2. Create Your First Character

```bash
tangent
```

Follow the interactive prompts to:
- Design characters visually
- Preview animations live
- Export Go code
- Save reusable `.go` files

### 3. Use in Your Go Project

```go
package main

import (
    "os"
    "local/characters/pkg/characters"
)

func main() {
    // Use a library character
    alien, _ := characters.Library("alien")
    characters.Animate(os.Stdout, alien, 5, 3)
    
    // Or create your own
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

## Installation

### For Users (CLI Tool)

```bash
# Clone and build
git clone https://github.com/wildreason/tangent.git
cd tangent/cmd/tangent
go build -o tangent .
cp tangent ~/.local/bin/
```

### For Developers (Package)

```bash
go get github.com/wildreason/tangent/pkg/characters
```

Or use as a local module:

```bash
# In your go.mod
replace github.com/wildreason/tangent => /path/to/tangent
```

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

## License

Proprietary - © 2025 Wildreason, Inc  
https://wildreason.com

---

## Links

- **Repository**: https://github.com/wildreason/tangent
- **Issues**: https://github.com/wildreason/tangent/issues
- **Website**: https://wildreason.com

---

**Built with ❤️ for terminal agent designers**
