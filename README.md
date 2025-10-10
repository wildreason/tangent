# Tangent

**Terminal Character Library for Go**

Simple character library with one-line access to animated Unicode block characters. Perfect for CLI tools, TUIs, and terminal applications.

---

## Quick Start

**Get a character and use it immediately:**

```go
import "github.com/wildreason/tangent/pkg/characters"

// Get pre-built character
alien, _ := characters.Library("alien")

// Use it immediately
characters.Animate(os.Stdout, alien, 5, 3)  // Done!
```

**That's it!** No configuration, no setup, no complexity.

---

## Core API (3 Functions Only)

### 1. Get Library Character
```go
alien, _ := characters.Library("alien")
robot, _ := characters.Library("robot")
pulse, _ := characters.Library("pulse")
```

### 2. Create Custom Character
```go
robot := characters.NewCharacterSpec("my-robot", 8, 4).
    AddFrame("idle", []string{"FRF", "LRL", "FRF", "LRL"}).
    Build()
```

### 3. Use Character
```go
characters.Animate(os.Stdout, alien, 5, 3)  // Animated
characters.ShowIdle(os.Stdout, robot)       // Static
```

---

## Library Characters

Pre-built, ready-to-use characters:

| Character | Type | Use Case | Example |
|-----------|------|----------|---------|
| **alien** | Animated (3 frames) | General purpose | `characters.Library("alien")` |
| **robot** | Static (1 frame) | Static display | `characters.Library("robot")` |
| **pulse** | Animated (3 frames) | Loading indicator | `characters.Library("pulse")` |
| **wave** | Animated (5 frames) | Progress indicator | `characters.Library("wave")` |
| **rocket** | Animated (4 frames) | Launch sequence | `characters.Library("rocket")` |

**Browse all characters:**
```bash
tangent gallery  # CLI playground
```

---

## Install

**One command:**
```bash
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash
```

**For Go developers:**
```go
import "github.com/wildreason/tangent/pkg/characters"
go mod tidy  # That's it!
```

---

## Use Cases

### Simple CLI Tools
```go
// Loading indicator
pulse, _ := characters.Library("pulse")
characters.Animate(os.Stdout, pulse, 8, 5)

// Success message
robot, _ := characters.Library("robot")
characters.ShowIdle(os.Stdout, robot)
```

### TUI Applications
```go
// Extract frames for your TUI framework
alien, _ := characters.Library("alien")
for _, frame := range alien.Frames {
    // Use frame.Lines in your TUI
    fmt.Println(frame.Lines)
}
```

### Custom Characters
```go
// Create your own
myChar := characters.NewCharacterSpec("custom", 6, 3).
    AddFrame("idle", []string{"FRF", "LRL", "FRF"}).
    Build()

characters.Animate(os.Stdout, myChar, 5, 3)
```

---

## Pattern System

Design characters with simple codes:

| Code | Block | Code | Block | Code | Block |
|------|-------|------|-------|------|-------|
| `F` | █ Full | `1` | ▘ Quad UL | `.` | ░ Light |
| `T` | ▀ Top | `2` | ▝ Quad UR | `:` | ▒ Medium |
| `B` | ▄ Bottom | `3` | ▖ Quad LL | `#` | ▓ Dark |
| `L` | ▌ Left | `4` | ▗ Quad LR | `_` | Space |
| `R` | ▐ Right | `5` | ▛ 3-Quad | `X` | Mirror |

**Example:** `"FRF"` becomes `"█▐█"`

**Full guide:** [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md)

---

## CLI Playground

The CLI is a playground for visual character creation:

```bash
tangent gallery    # Browse library characters
tangent create     # Create custom character visually
tangent patterns   # Pattern reference guide
```

**The CLI is NOT the primary tool** - it's just a playground for experimentation and visual creation.

---

## Architecture

```
┌─────────────────────────────────────────────────┐
│ Your Application                                │
│ • characters.Library("alien")                   │
│ • characters.Animate(os.Stdout, alien, 5, 3)   │
├─────────────────────────────────────────────────┤
│ Tangent Core                                    │
│ • Character library                             │
│ • Pattern compiler                              │
│ • Animation engine                              │
└─────────────────────────────────────────────────┘
```

**Tangent = Character Library**  
**Your App = Uses Characters**  
**Simple, focused, effective**

---

## Requirements

- **Go 1.21+** only
- **Terminal** with Unicode block element support
- **Zero external dependencies** for core functionality

---

## Examples

### Basic Usage
```go
package main

import (
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    // Get character
    alien, _ := characters.Library("alien")
    
    // Use it
    characters.Animate(os.Stdout, alien, 5, 3)
}
```

### Custom Character
```go
package main

import (
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    // Create custom character
    robot := characters.NewCharacterSpec("my-robot", 8, 4).
        AddFrame("idle", []string{"FRF", "LRL", "FRF", "LRL"}).
        Build()
    
    // Use it
    characters.ShowIdle(os.Stdout, robot)
}
```

---

## Documentation

- [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md) - Pattern reference
- [`docs/LIBRARY.md`](docs/LIBRARY.md) - Character library details
- [`CONTRIBUTING.md`](CONTRIBUTING.md) - How to contribute
- [`CHANGELOG.md`](CHANGELOG.md) - Version history

---

## Contributing

Contributions welcome! See [`CONTRIBUTING.md`](CONTRIBUTING.md).

**Want to add a character to the library?**
1. Design with `tangent create`
2. Test thoroughly
3. Submit a PR

---

## License

MIT License © 2025 Wildreason, Inc

---

## Links

- **Repository:** https://github.com/wildreason/tangent
- **Issues:** https://github.com/wildreason/tangent/issues
- **Releases:** https://github.com/wildreason/tangent/releases

---

**Built for terminal developers**  
**Designed for simplicity**  
**Works anywhere**
