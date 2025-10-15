# Pattern Reference

Quick reference for hex-style block element patterns.

## Pattern Codes

| Code | Block | Name | Code | Block | Name |
|------|-------|------|------|-------|------|
| `F` | █ | Full | `5` | ▛ | 3-Quad (reverse of 4) |
| `T` | ▀ | Top | `6` | ▜ | 3-Quad (reverse of 3) |
| `B` | ▄ | Bottom | `7` | ▙ | 3-Quad (reverse of 2) |
| `L` | ▌ | Left | `8` | ▟ | 3-Quad (reverse of 1) |
| `R` | ▐ | Right | `1` | ▘ | Quad UL |
| `.` | ░ | Light | `2` | ▝ | Quad UR |
| `:` | ▒ | Medium | `3` | ▖ | Quad LL |
| `#` | ▓ | Dark | `4` | ▗ | Quad LR |
| `_` | (space) | Empty | `\` | ▚ | Backward diagonal |
|   |   |   | `/` | ▞ | Forward diagonal |

**Pattern Logic:**
- **Quadrants 1-4**: Single quadrants
- **Quadrants 5-8**: Three quadrants (reverse of 1-4)
  - `1` ↔ `8` (reverse pair)
  - `2` ↔ `7` (reverse pair)
  - `3` ↔ `6` (reverse pair)
  - `4` ↔ `5` (reverse pair)
- **Diagonals**: `\` = backward, `/` = forward

## Usage

```go
// Pattern format: "__R6FFF6L__"
alien := characters.NewCharacterSpec("alien", 11, 3).
    AddFrame("idle", []string{
        "__R6FFF6L__",
        "_T6FFFFF5T_",
        "___11__22__",
    })
```

## Visual Builder

Use the interactive CLI builder to design characters visually:

```bash
go run tools/builder/main.go
```

The builder lets you:
- Drag and drop block elements like Lego
- See live preview
- Use `X` for auto-mirroring
- Export as pattern code
- Copy/paste ready for use

## Examples

### Simple Character
```
Pattern: "L6FFF6R"
Result:  ▌▜███▜▐
```

### Alien (3 frames)
```
Frame 1:  __R6FFF6L__    →    ▐▜███▜▌  
          _T6FFFFF5T_    →   ▀▜█████▛▀ 
          ___11__22__    →     ▘▘ ▝▝  

Frame 2:  __R6FFF6L__
          5T6FFFFF5T_
          __11__22___

Frame 3:  __R6FFF6L__
          _T6FFFFF5T6
          ___11__22__
```

## Library Characters

Use pre-built characters without creating patterns:

```go
alien, _ := characters.Library("alien")
characters.Animate(os.Stdout, alien, 4, 2)
```

See `LIBRARY.md` for available characters.
