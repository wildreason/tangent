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

### Interactive Mode

```bash
tangent  # Opens visual builder
```

Create your character visually:
- Design frame by frame
- Preview animation live
- Export Go code

### CLI Mode (for AI agents)

```bash
# Create character
tangent create --name robot --width 11 --height 3 \
  --frame idle '__R6FFF6L__,_T5FFFFF6T_,___11_22___' \
  --output robot.go --package agent

# Animate library character
tangent animate --name alien --fps 10 --loops 5

# Export saved session
tangent export --session mychar --output mychar.go
```

### Use in Your Code

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

## CLI Commands

### `tangent create`
Create character from command line:
```bash
tangent create --name NAME --width W --height H \
  --frame FRAME_NAME "pattern,lines,here" \
  [--output file.go] [--package pkg]
```

### `tangent animate`
Show animation in terminal:
```bash
# Library character
tangent animate --name alien --fps 10 --loops 5

# From saved session
tangent animate --session mychar --fps 5 --loops 3
```

### `tangent export`
Export session to Go code:
```bash
tangent export --session mychar --output mychar.go --package agent
```

**Get help**: `tangent help`

---

## Features

- **Zero dependencies** - Pure Go stdlib
- **Interactive + CLI modes** - Visual builder or command line
- **AI-friendly** - Non-interactive CLI for agents
- **Pattern-based** - Simple, intuitive codes
- **Library characters** - Pre-built animations
- **Export ready** - Generate Go code
- **Cross-platform** - macOS, Linux, Windows

---

## Use Cases

- **AI Agents** - Characters for terminal agents
- **CLI Apps** - Interactive terminal UIs
- **Terminal Games** - Animated sprites
- **Loading States** - Progress indicators
- **Status Display** - Visual feedback
- **DevOps Tools** - Build/deploy animations

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
