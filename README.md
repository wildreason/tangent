# Tangent

**Terminal character animation library for Go**

Build animated CLI agents with zero dependencies.

---

## Install

```bash
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash
```

This installs:
- ✅ `tangent` command (visual builder)
- ✅ Go package (for your code)

**Alternative**: Via Go
```bash
go install github.com/wildreason/tangent/cmd/tangent@latest
```

---

## Quick Start

### 1. Design a Character

```bash
tangent  # Opens interactive builder
```

Create your character visually:
- Design frame by frame
- Preview animation live
- Export Go code

### 2. Use in Your Code

```go
import "github.com/wildreason/tangent/pkg/characters"

// Option A: Use library character
alien, _ := characters.Library("alien")
characters.Animate(os.Stdout, alien, 5, 3)

// Option B: Use your own (from tangent export)
spec := characters.NewCharacterSpec("robot", 11, 3).
    AddFrame("idle", []string{
        "__R6FFF6L__",
        "_T6FFFFF5T_",
        "___11_22___",
    })

robot, _ := spec.Build()
characters.Animate(os.Stdout, robot, 5, 3)
```

---

## Pattern Codes

Single-character codes for block elements:

| Code | Block | Code | Block | Code | Block |
|------|-------|------|-------|------|-------|
| `F` | █ | `1` | ▘ | `.` | ░ |
| `T` | ▀ | `2` | ▝ | `:` | ▒ |
| `B` | ▄ | `3` | ▖ | `#` | ▓ |
| `L` | ▌ | `4` | ▗ | `_` | space |
| `R` | ▐ | `5` | ▛ | `X` | mirror |
|     |   | `6` | ▜ |     |       |
|     |   | `7` | ▙ |     |       |
|     |   | `8` | ▟ |     |       |

**Full guide**: [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md)

---

## Features

- **Zero dependencies** - Pure Go stdlib
- **Visual builder** - Design without coding
- **Pattern-based** - Simple, intuitive codes
- **Library characters** - Pre-built animations
- **Export ready** - Copy-paste Go code
- **Cross-platform** - macOS, Linux, Windows

---

## Use Cases

- CLI applications
- Terminal games  
- Loading animations
- Status indicators
- Agent UX
- Developer tools

---

## Examples

See working examples in [`examples/`](examples/):
- [`examples/demo/`](examples/demo/) - Library + custom characters
- [`examples/tokyo/`](examples/tokyo/) - Character created with tangent

Run demo:
```bash
cd examples/demo
go run main.go
```

---

## Library Characters

Pre-built characters you can use immediately:

```go
alien, _ := characters.Library("alien")
characters.Animate(os.Stdout, alien, 5, 3)
```

**Available**: `alien` (3 frames, waving animation)

**Full library**: [`docs/LIBRARY.md`](docs/LIBRARY.md)

---

## API Reference

### Create Character
```go
spec := characters.NewCharacterSpec(name, width, height).
    AddFrame("idle", []string{"pattern..."}).
    AddFrame("move", []string{"pattern..."})

char, err := spec.Build()
```

### Animate
```go
// Animate at 5 FPS for 3 loops
characters.Animate(os.Stdout, char, 5, 3)

// Show single frame
characters.ShowIdle(os.Stdout, char)
```

### Registry
```go
// Register character
characters.Register(char)

// Retrieve later
char, _ := characters.Get("name")
```

---

## Documentation

- [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md) - Pattern codes reference
- [`docs/LIBRARY.md`](docs/LIBRARY.md) - Pre-built characters
- [`CONTRIBUTING.md`](CONTRIBUTING.md) - How to contribute
- [`CHANGELOG.md`](CHANGELOG.md) - Version history
- [`ROADMAP.md`](ROADMAP.md) - Future plans

---

## Requirements

- Go 1.21 or higher
- Terminal with Unicode support

---

## Contributing

Contributions welcome! See [`CONTRIBUTING.md`](CONTRIBUTING.md).

---

## License

MIT License - © 2025 Wildreason, Inc

See [`LICENSE`](LICENSE) for details.

---

## Links

- **Repository**: https://github.com/wildreason/tangent
- **Issues**: https://github.com/wildreason/tangent/issues
- **Releases**: https://github.com/wildreason/tangent/releases

---

**Built for terminal developers**
