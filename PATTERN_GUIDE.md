# Hex-Style Pattern Guide

## Overview

Create block element characters using compact, hex-color-like patterns!

```go
// Like hex colors:  #F8394839
// But for blocks:   00R9FFF9L00
```

## Why Hex-Style?

### Before (Comma-Separated)
```go
"0,0,rf,comp1,fb,fb,fb,comp2,lf,0,0"  // Hard to read, verbose
```

### After (Hex-Style)
```go
"00R9FFF9L00"  // Clean, compact, like a color code!
```

## Pattern Legend

### Basic Blocks
```
F = █  Full block
T = ▀  Top half (Upper)
B = ▄  Bottom half (Lower)
L = ▌  Left half
R = ▐  Right half
```

### Shading
```
. = ░  Light shade (think: light dot)
: = ▒  Medium shade (think: colon density)
# = ▓  Dark shade (think: hash density)
```

### Single Quadrants
```
1 = ▘  Quadrant 1 (Upper Left)
2 = ▝  Quadrant 2 (Upper Right)
3 = ▖  Quadrant 3 (Lower Left)
4 = ▗  Quadrant 4 (Lower Right)
```

### Three-Quadrant Composites
```
7 = ▛  UL+UR+LL (looks like a 7)
9 = ▜  UL+UR+LR (looks like a 9)
6 = ▙  UL+LL+LR (looks like a 6)
8 = ▟  UR+LL+LR (looks like an 8)
```

### Diagonals
```
/ = ▚  Forward slash pattern
\ = ▞  Backslash pattern
```

### Space
```
0 = Space (like hex: #000000)
_ = Space (for readability)
```

## Examples

### Simple Character
```go
// Pattern:  "L9FFF9R"
// Result:   ▌▜███▜▐

char := characters.NewCharacterSpec("simple", 7, 1).
    AddFrame("idle", []string{"L9FFF9R"})
```

### Alien Character
```go
// Frame 1: Idle
"00R9FFF9L00"  →   ▐▜███▜▌  
"0T9FFFFF7T0"  →  ▀▜█████▛▀ 
"00011000220"  →    ▘▘   ▝▝ 

alien := characters.NewCharacterSpec("alien", 11, 3).
    AddFrame("idle", []string{
        "00R9FFF9L00",
        "0T9FFFFF7T0",
        "00011000220",
    })
```

### Animation (Easy Editing!)
```go
// Just change ONE character per frame!
alien := characters.NewCharacterSpec("alien", 11, 3).
    AddFrame("idle", []string{
        "00R9FFF9L00",
        "0T9FFFFF7T0",  // ← Middle starts with 0
        "00011000220",
    }).
    AddFrame("left", []string{
        "00R9FFF9L00",
        "7T9FFFFF7T0",  // ← Changed 0 to 7 (one char!)
        "00011000220",
    }).
    AddFrame("right", []string{
        "00R9FFF9L00",
        "0T9FFFFF7T9",  // ← Changed last 0 to 9 (one char!)
        "00011000220",
    })
```

### With Shading
```go
// Use . : # for shading effects
char := characters.NewCharacterSpec("shaded", 9, 3).
    AddFrame("idle", []string{
        "0.:#F#:.0",  // ░▒▓█▓▒░
        "00:#F#:00",
        "000#F#000",
    })
```

## Editing Tips

### 1. Visual Pattern
The pattern code itself shows the shape:
```
"00R9FFF9L00"  ← Two spaces, right, composite, fills, composite, left, spaces
  ▐▜███▜▌      ← Result looks like the pattern!
```

### 2. Quick Changes
Need to make the head wider? Just add more F's:
```
OLD: "00R9FFF9L00"
NEW: "00R9FFFF9L0"  ← Added one F, removed one 0
```

### 3. Animation Frames
Change just one character to animate:
```
Frame 1: "0T9FFFFF7T0"
Frame 2: "7T9FFFFF7T0"  ← Changed first 0 to 7
Frame 3: "0T9FFFFF7T9"  ← Changed last 0 to 9
```

### 4. Symmetry
Use the same code on both sides for symmetry:
```
"00R9FFF9L00"
  └─┘   └─┘
  Same on both sides!
```

## Validation

The pattern compiler validates your patterns:
```go
compiler := characters.NewPatternCompiler()

// Valid
compiler.ValidatePattern("00R9FFF9L00")  // ✓

// Invalid
compiler.ValidatePattern("00X9FFF9L00")  // ✗ Unknown char 'X'
```

## Complete Example

```go
package main

import (
    "os"
    "local/characters/pkg/characters"
)

func main() {
    // Create alien with hex-style patterns
    alien := characters.NewCharacterSpec("alien", 11, 3).
        AddFrame("idle", []string{
            "00R9FFF9L00",
            "0T9FFFFF7T0",
            "00011000220",
        }).
        AddFrame("wave", []string{
            "00R9FFF9L00",
            "7T9FFFFF7T0",  // Just changed one character!
            "00011000220",
        })

    char, _ := alien.Build()
    characters.Register(char)
    
    // Animate!
    characters.Animate(os.Stdout, char, 4, 3)
}
```

## Quick Reference Card

```
┌─────────────────────────────────────────┐
│  Hex-Style Block Element Pattern Codes  │
├─────────────────────────────────────────┤
│  F=█  T=▀  B=▄  L=▌  R=▐               │
│  7=▛  9=▜  6=▙  8=▟                    │
│  1=▘  2=▝  3=▖  4=▗                    │
│  .=░  :=▒  #=▓  0=Space                 │
├─────────────────────────────────────────┤
│  Example: "00R9FFF9L00" → "  ▐▜███▜▌  " │
└─────────────────────────────────────────┘
```

## Why It Works

1. **Like hex colors** - Familiar paradigm (`#F8394839`)
2. **Single characters** - Easy to type and edit
3. **No commas** - Continuous string like color codes
4. **Visual mapping** - F=Full, T=Top, L=Left makes sense
5. **Number patterns** - 7,9,6,8 look like the blocks they represent
6. **Quick editing** - Change one char to modify the design

Happy character building! 🎨
