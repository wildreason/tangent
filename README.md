# Block Characters

A Go library for creating terminal characters using Unicode Block Elements (U+2580–U+259F).

## Installation

```bash
go get local/characters
```

## Quick Start

Create characters using hex-style patterns - compact and intuitive like hex colors:

```go
package main

import (
    "os"
    "local/characters/pkg/characters"
)

func main() {
    // Create a character using hex-style patterns
    alien := characters.NewCharacterSpec("alien", 11, 3).
        AddFrame("idle", []string{
            "00R9FFF9L00",
            "0T9FFFFF7T0",
            "00011000220",
        }).
        AddFrame("wave", []string{
            "00R9FFF9L00",
            "7T9FFFFF7T0",
            "00011000220",
        })

    char, _ := alien.Build()
    characters.Register(char)
    characters.Animate(os.Stdout, char, 4, 2)
}
```

## Pattern Codes

Single-character codes for easy editing:

| Code | Block | Code | Block | Code | Block |
|------|-------|------|-------|------|-------|
| `F` | █ | `7` | ▛ | `1` | ▘ |
| `T` | ▀ | `9` | ▜ | `2` | ▝ |
| `B` | ▄ | `6` | ▙ | `3` | ▖ |
| `L` | ▌ | `8` | ▟ | `4` | ▗ |
| `R` | ▐ | `.` | ░ | `0` | (space) |
|     |   | `:` | ▒ | `_` | (space) |
|     |   | `#` | ▓ |     |       |

## API

### Pattern-Based (Recommended)

```go
// Create character specification
spec := characters.NewCharacterSpec(name, width, height).
    AddFrame(frameName, []string{"pattern1", "pattern2", ...})

char, err := spec.Build()
```

### Builder API

```go
// Traditional builder approach
char, err := characters.NewBuilder(name, width, height).
    Pattern("L9FFF9R").
    NewFrame().
    Pattern("R9FFF9L").
    Build()
```

### Animation

```go
// Animate character
characters.Animate(os.Stdout, char, fps, loops)

// Show static frame
characters.ShowIdle(os.Stdout, char)
```

### Registry

```go
// Register and retrieve characters
characters.Register(char)
char, err := characters.Get("name")
names := characters.List()
```

## Examples

```bash
# Hex-style patterns
go run examples/hex_style/main.go
go run examples/compact/main.go

# Builder API
go run examples/basic/main.go
go run examples/simple/main.go

# Single line character
go run examples/one_line/main.go
```

## Documentation

See `PATTERN_GUIDE.md` for detailed pattern syntax and examples.

## License

Proprietary - Compass AI Testing Intelligence
