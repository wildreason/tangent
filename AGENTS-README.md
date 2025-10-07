# Tangent for AI Agents

**Command-line character designer for terminal agents**

This guide is for AI agents that want to create animated terminal characters.

---

## Installation

```bash
# Install tangent CLI
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash

# Verify installation
tangent version
```

---

## Quick Start

### Create a Simple Character

```bash
tangent create --name robot --width 7 --height 3 \
  --frame idle 'R6FFF6L,T5FFF6T,_11_22_' \
  --output robot.go --package agent
```

This creates `robot.go` with a `GetRobot()` function you can import.

---

## Pattern System

Use single-character codes to build characters:

| Code | Shape | Code | Shape | Code | Shape |
|------|-------|------|-------|------|-------|
| `F`  | █ Full block | `1` | ▘ Upper-left quad | `.` | ░ Light shade |
| `T`  | ▀ Top half | `2` | ▝ Upper-right quad | `:` | ▒ Medium shade |
| `B`  | ▄ Bottom half | `3` | ▖ Lower-left quad | `#` | ▓ Dark shade |
| `L`  | ▌ Left half | `4` | ▗ Lower-right quad | `_` | Space |
| `R`  | ▐ Right half | `5` | ▛ Three-quad (no LR) | `X` | Mirror line |
|      |   | `6` | ▜ Three-quad (no LL) |     |       |
|      |   | `7` | ▙ Three-quad (no UR) |     |       |
|      |   | `8` | ▟ Three-quad (no UL) |     |       |

**Rules:**
- Each line is a string of pattern codes
- Separate lines with commas
- `X` mirrors everything before it on that line
- Width = number of characters per line
- Height = number of lines

---

## Commands

### `tangent create`

Create character from patterns:

```bash
tangent create --name NAME --width W --height H \
  --frame FRAME_NAME "line1,line2,line3" \
  [--frame FRAME_NAME "line1,line2,line3"] \
  [--output file.go] \
  [--package pkg]
```

**Options:**
- `--name` - Character identifier (required)
- `--width` - Character width in blocks (required)
- `--height` - Character height in blocks (required)
- `--frame` - Add frame (can repeat for animation)
- `--output` - Save to Go file (optional, shows preview if omitted)
- `--package` - Go package name (default: `main`)

**Examples:**

```bash
# Single frame, preview only
tangent create --name bot --width 5 --height 2 \
  --frame idle 'R6F6L,T5F6T'

# Multi-frame with export
tangent create --name alien --width 7 --height 3 \
  --frame idle 'L6FFF6R,T5FFF6T,_1_2_' \
  --frame wave 'R6FFF6L,T5FFF6T,_2_1_' \
  --output alien.go --package myagent

# Using mirror for symmetry
tangent create --name face --width 9 --height 3 \
  --frame smile '___R6FX6L___,__T5FFFX5T__,_________' \
  --output face.go
```

---

### `tangent animate`

Preview animation in terminal:

```bash
# Animate library character
tangent animate --name alien --fps 10 --loops 5

# Animate from saved session (interactive mode)
tangent animate --session mychar --fps 5 --loops 3
```

**Options:**
- `--name` - Library character name
- `--session` - Saved session from interactive mode
- `--fps` - Frames per second (default: 5)
- `--loops` - Number of animation loops (default: 3)

---

### `tangent export`

Export interactive session to Go code:

```bash
tangent export --session mychar \
  --output mychar.go \
  --package agent
```

**Options:**
- `--session` - Session name from interactive mode (required)
- `--output` - Output file (default: stdout)
- `--package` - Go package name (default: `main`)

---

## Example Workflow

### 1. Design Character

```bash
# Create a thinking indicator
tangent create --name thinking --width 11 --height 3 \
  --frame dots1 '___________,_F_________,___________' \
  --frame dots2 '___________,___F_______,___________' \
  --frame dots3 '___________,_____F_____,___________' \
  --frame dots4 '___________,_______F___,___________' \
  --frame dots5 '___________,_________F_,___________' \
  --output thinking.go --package indicators
```

### 2. Use Generated Code

The generated `thinking.go` will contain:

```go
package indicators

import (
	"github.com/wildreason/tangent/pkg/characters"
)

// GetThinking returns the thinking character
func GetThinking() (*characters.Character, error) {
	spec := characters.NewCharacterSpec("thinking", 11, 3)
	spec = spec.AddFrame("dots1", []string{
		"___________",
		"_F_________",
		"___________",
	})
	// ... more frames
	return spec.Build()
}
```

### 3. Integrate in Your Agent

```go
package main

import (
	"os"
	"yourapp/indicators"
	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	thinking, _ := indicators.GetThinking()
	
	// Show animation while processing
	go characters.Animate(os.Stdout, thinking, 10, 0) // infinite loops
	
	// Do your processing
	processTask()
	
	// Stop animation (Ctrl+C or program exit)
}
```

---

## Design Patterns

### Simple Face

```bash
tangent create --name face --width 7 --height 3 \
  --frame smile 'R6FFF6L,T5FFF6T,_1_2_'
```

### Symmetric Character (using X mirror)

```bash
tangent create --name bot --width 8 --height 4 \
  --frame idle '__R6FX6L__,_T5FFFX5T_,__11_X_22__,____FFFF____'
```

### Loading Spinner

```bash
tangent create --name spinner --width 3 --height 3 \
  --frame f1 '_T_,___,___' \
  --frame f2 '__R,___,___' \
  --frame f3 '___,___,_B_' \
  --frame f4 '___,___,L__' \
  --output spinner.go
```

### Progress Bar

```bash
tangent create --name progress --width 10 --height 1 \
  --frame p0 '__________' \
  --frame p1 'R_________' \
  --frame p2 'RR________' \
  --frame p3 'RRR_______' \
  --frame p4 'RRRR______' \
  --frame p5 'RRRRR_____' \
  --output progress.go
```

### Agent Avatar

```bash
tangent create --name agent --width 9 --height 5 \
  --frame idle '___R6F6L___,__T5FFF6T__,_1_______2_,___________,___________' \
  --frame talk '___R6F6L___,__T5FFF6T__,_1_______2_,____BBB____,___________' \
  --output agent.go
```

---

## Pattern Tips

### Composing Shapes

**Head/Body:**
- Use `6` (▜) and `5` (▛) for rounded tops
- Use `T` (▀) for flat tops
- Use `F` (█) for solid blocks

**Eyes:**
- Use `1` (▘) and `2` (▝) for simple eyes
- Use `.` (░) for sleepy/closed eyes
- Use `F` for open/alert eyes

**Symmetry:**
- Use `X` to mirror: `R6FFFX` → `R6FFF6L`
- Saves typing and ensures balance

**Animation:**
- Create multiple frames with slight variations
- Eye blinks: frame1 uses `1_2`, frame2 uses `._.'
- Movement: shift characters left/right by changing `_` placement

---

## Error Handling

### Common Issues

**Frame line count mismatch:**
```bash
Error: frame 'idle' has 2 lines, expected 3
```
→ Fix: Ensure pattern has exactly `--height` lines (comma-separated)

**Invalid width:**
```bash
Error building character: line 0 has 5 blocks, expected 7
```
→ Fix: Each line must have exactly `--width` characters

**Missing arguments:**
```bash
Error: --name is required
```
→ Fix: Provide all required flags: `--name`, `--width`, `--height`, `--frame`

---

## Library Characters

Pre-built characters you can use immediately:

```bash
# List available (currently: alien)
tangent animate --name alien --fps 5 --loops 3
```

Use in code:
```go
import "github.com/wildreason/tangent/pkg/characters"

alien, _ := characters.Library("alien")
characters.Animate(os.Stdout, alien, 5, 3)
```

---

## Advanced Usage

### Multi-Package Organization

```bash
# Create characters for different contexts
tangent create --name loader --width 5 --height 2 \
  --frame spin1 'R6F6L,T5F6T' \
  --output loaders/loader.go --package loaders

tangent create --name success --width 7 --height 3 \
  --frame check 'R6FFF6L,T5FFF6T,_11_22_' \
  --output feedback/success.go --package feedback
```

### Animation Control

```go
// Animate at 10 FPS for 5 loops
characters.Animate(os.Stdout, char, 10, 5)

// Show single idle frame
characters.ShowIdle(os.Stdout, char)

// Infinite animation (use with goroutine)
characters.Animate(os.Stdout, char, 5, 0)
```

### Character Registry

```go
// Register characters globally
characters.Register(myChar)

// Retrieve anywhere in your app
char, _ := characters.Get("mychar")
```

---

## Shell Integration

### Bash Script Example

```bash
#!/bin/bash
# generate-characters.sh

# Create multiple characters
tangent create --name bot1 --width 7 --height 3 \
  --frame idle 'R6FFF6L,T5FFF6T,_1_2_' \
  --output chars/bot1.go --package characters

tangent create --name bot2 --width 7 --height 3 \
  --frame idle 'L6FFF6R,T5FFF6T,_2_1_' \
  --output chars/bot2.go --package characters

echo "✓ Characters generated in chars/"
```

### Python Integration

```python
import subprocess
import json

def create_character(name, width, height, frames, output=None):
    """Create character using tangent CLI"""
    cmd = [
        "tangent", "create",
        "--name", name,
        "--width", str(width),
        "--height", str(height),
    ]
    
    for frame_name, pattern in frames.items():
        cmd.extend(["--frame", frame_name, pattern])
    
    if output:
        cmd.extend(["--output", output, "--package", "characters"])
    
    result = subprocess.run(cmd, capture_output=True, text=True)
    return result.returncode == 0

# Usage
create_character(
    name="assistant",
    width=7,
    height=3,
    frames={
        "idle": "R6FFF6L,T5FFF6T,_1_2_",
        "active": "L6FFF6R,T5FFF6T,_2_1_"
    },
    output="assistant.go"
)
```

---

## Zero Dependencies

Tangent uses **only Go standard library**. No external dependencies.

**Requirements:**
- Go 1.21+ (if building from source)
- Terminal with Unicode support
- That's it!

---

## Help & Documentation

```bash
# Show all commands
tangent help

# Check version
tangent version

# Interactive mode (visual builder)
tangent
```

**Full docs:** https://github.com/wildreason/tangent

---

## Quick Reference

```bash
# CREATE: Build from CLI
tangent create --name NAME --width W --height H \
  --frame FRAME "pattern" --output file.go

# ANIMATE: Preview in terminal
tangent animate --name LIBRARY_CHAR --fps 10 --loops 5

# EXPORT: Save session
tangent export --session NAME --output file.go

# HELP: Show usage
tangent help
```

---

**Built for AI agents and terminal developers**

MIT License © 2025 Wildreason, Inc

