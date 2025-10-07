# Tangent - Terminal Agent Designer

**Design characters for your CLI agents** - An interactive visual builder powered by Wildreason Characters.

## What is Tangent?

Tangent is a standalone CLI tool that makes it easy to design Unicode block characters for terminal applications. Perfect for creating:

- **Terminal agents** with personality
- **CLI loading indicators** with character
- **ASCII art logos** for your tools
- **Animated mascots** for your projects

## Installation

### Option 1: Direct Download (Easiest)

Download the pre-built binary for your platform:

**macOS (Apple Silicon)**
```bash
curl -L https://releases.wildreason.com/tangent/v0.0.1/tangent-macos-arm64 -o tangent
chmod +x tangent
sudo mv tangent /usr/local/bin/
```

**macOS (Intel)**
```bash
curl -L https://releases.wildreason.com/tangent/v0.0.1/tangent-macos-amd64 -o tangent
chmod +x tangent
sudo mv tangent /usr/local/bin/
```

**Linux**
```bash
curl -L https://releases.wildreason.com/tangent/v0.0.1/tangent-linux-amd64 -o tangent
chmod +x tangent
sudo mv tangent /usr/local/bin/
```

**Windows**
```powershell
# Download from: https://releases.wildreason.com/tangent/v0.0.1/tangent-windows-amd64.exe
# Move to a directory in your PATH
```

### Option 2: Build from Source

```bash
git clone https://github.com/wildreason/characters.git
cd characters/cmd/tangent
go build -o tangent .
```

## Quick Start

Simply run:

```bash
tangent
```

You'll see the main menu with options to:

1. **Create new character** - Start a fresh character project
2. **Load character project** - Resume working on a saved project
3. **Browse library characters** - Explore pre-built characters
4. **Preview library character** - See characters in action
5. **View palette** - Reference for all block codes
6. **Exit**

## Features

### ◆ Multi-Frame Session Manager
- Auto-saves your progress
- Resume work anytime
- Manage multiple character projects

### ◆ Visual Builder
- Interactive palette reference
- Line-by-line character construction
- Live preview as you build
- Pattern confirmation at each step

### ◆ Smart Mirroring
Use `X` as a mirror marker to auto-mirror patterns:
```
__R5FX → __R5F5R__ (automatically mirrored!)
```

### ◆ Library Integration
Browse and preview pre-built characters from the Wildreason library:
- `alien` - Animated alien with waving hands (3 frames)

### ◆ Code Export
Export your character as ready-to-use Go code:
```go
spec := characters.NewCharacterSpec("mychar", 11, 3)
    .AddFrame("idle", []string{
        "__R6FFF6L__",
        "_T6FFFFF5T_",
        "___11__22__",
    })
```

## Block Palette Reference

### Basic Blocks
- `F` = █ (Full block)
- `T` = ▀ (Top half)
- `B` = ▄ (Bottom half)
- `L` = ▌ (Left half)
- `R` = ▐ (Right half)

### Quadrants (1-4)
- `1` = ▘ (Upper Left)
- `2` = ▝ (Upper Right)
- `3` = ▖ (Lower Left)
- `4` = ▗ (Lower Right)

### Three-Quadrants (5-8)
- `5` = ▛ (UL+UR+LL) - reverse of `4`
- `6` = ▜ (UL+UR+LR) - reverse of `3`
- `7` = ▙ (UL+LL+LR) - reverse of `2`
- `8` = ▟ (UR+LL+LR) - reverse of `1`

### Shading
- `.` = ░ (Light)
- `:` = ▒ (Medium)
- `#` = ▓ (Dark)

### Diagonals
- `\` = ▚ (Backward diagonal)
- `/` = ▞ (Forward diagonal)

### Special
- `_` = Space
- `X` = Mirror marker

## Example Workflow

1. **Start Tangent**
   ```bash
   tangent
   ```

2. **Create a new character**
   - Choose option `1`
   - Name: `mybot`
   - Dimensions: `11x3`

3. **Add first frame**
   - Frame name: `idle`
   - Enter patterns line by line
   - Confirm each line
   - See progressive preview

4. **Add animation frames**
   - Frame name: `wave_left`
   - Frame name: `wave_right`

5. **Export code**
   - Choose option `4`
   - Copy the generated Go code
   - Use in your project!

## Use Your Character in Go

After exporting, use your character:

```go
package main

import (
    "os"
    "local/characters/pkg/characters"
)

func main() {
    spec := characters.NewCharacterSpec("mybot", 11, 3).
        AddFrame("idle", []string{
            "__R6FFF6L__",
            "_T6FFFFF5T_",
            "___11__22__",
        })
    
    char, _ := spec.Build()
    characters.ShowIdle(os.Stdout, char)
}
```

## Session Storage

All your character projects are saved to `~/.tangent/` as JSON files. You can:
- Resume any project anytime
- Share project files with teammates
- Version control your character designs

## About

Tangent is part of the **Wildreason, Inc** suite.

- **Version**: v0.0.1
- **License**: Proprietary
- **Website**: https://wildreason.com
- **Support**: support@wildreason.com

## Building from Source

```bash
cd cmd/tangent
go build -o tangent .

# Cross-compile for other platforms
GOOS=linux GOARCH=amd64 go build -o tangent-linux-amd64 .
GOOS=darwin GOARCH=arm64 go build -o tangent-macos-arm64 .
GOOS=windows GOARCH=amd64 go build -o tangent-windows-amd64.exe .
```

---

Made with ◆ by Wildreason, Inc

