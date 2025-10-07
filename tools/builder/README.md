# Character Builder CLI

Interactive visual builder for creating block element characters.

## Usage

```bash
go run tools/builder/main.go
```

Or build and run:

```bash
cd tools/builder
go build
./builder
```

## Features

- **Visual Palette** - See all available block elements with codes (with line spacing!)
- **Input Validation** - Dimensions must be valid integers (1-100 width, 1-50 height)
- **Line-by-line Entry** - Build your character one line at a time
- **Progressive Preview** - See your character building up as you add each line
- **Per-Line Preview** - See what each line looks like before accepting
- **Frame Mirroring** - Use `X` to auto-mirror and create symmetry
- **Line Confirmation** - Accept or retry each line (y/N prompt)
- **Pattern Export** - Get the hex-style pattern code ready to copy
- **Go Code Export** - Get complete Go code ready to use

## Example Session

```
╔══════════════════════════════════════════╗
║  Block Character Builder v0.0.1         ║
╚══════════════════════════════════════════╝

╔══════════════════════════════════════════╗
║  BLOCK PALETTE                           ║
╚══════════════════════════════════════════╝

  Basic:    F=█  T=▀  B=▄  L=▌  R=▐
  Shading:  .=░  :=▒  #=▓
  Quads:    1=▘  2=▝  3=▖  4=▗
  3-Quads:  7=▛  9=▜  6=▙  8=▟
  Diagonal: /=▚  \=▞
  Empty:    0=Space  _=Space

◢ Enter width (e.g., 11): 7
◢ Enter height (e.g., 3): 1

▢ Line 1 (enter 7 codes, e.g., 00R9FFF9L00):
  Pattern: L9FFF9R
  Preview: ▌▜███▜▐

╔══════════════════════════════════════════╗
║  PREVIEW                                 ║
╚══════════════════════════════════════════╝
  ▌▜███▜▐

╔══════════════════════════════════════════╗
║  PATTERN CODE (copy this!)               ║
╚══════════════════════════════════════════╝
  "L9FFF9R"

╔══════════════════════════════════════════╗
║  GO CODE (ready to use!)                 ║
╚══════════════════════════════════════════╝

char := characters.NewCharacterSpec("mychar", 7, 1).
    AddFrame("idle", []string{
        "L9FFF9R"
    })

Done! Copy the pattern code above.
```

## Workflow

1. **Run the builder** - `go run tools/builder/main.go`
2. **Set dimensions** - Enter width and height
3. **Build line by line** - Enter pattern codes for each line
4. **See live preview** - Watch your character take shape
5. **Copy the code** - Use the generated pattern in your Go code

## Tips

- Use `0` or `_` for empty spaces
- Use `X` for mirroring - Example: `00R9FX` becomes `00R9F9R00` (auto-mirrored!)
- The palette shows all available blocks with line spacing for easy reading
- Pattern codes are case-sensitive
- Line-by-line preview helps catch mistakes
- Press `n` to retry any line if you make a mistake
- Press `y` or Enter to accept and move to next line
- Auto-pads if you enter too few characters
- Auto-truncates if you enter too many

## Integration

The builder outputs pattern code compatible with the characters package:

```go
import "local/characters/pkg/characters"

// Use the generated pattern
char := characters.NewCharacterSpec("mychar", 11, 3).
    AddFrame("idle", []string{
        "00R9FFF9L00",
        "0T9FFFFF7T0",
        "00011000220",
    })
```

