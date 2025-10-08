# Tangent

**Terminal Character Design System for Go**

Design animated Unicode block characters with an intuitive pattern system, visual builder, and pre-built library. Works with any TUI framework.

---

## Why Tangent?

✅ **Best-in-class pattern system** - Simple, memorable character design  
✅ **Visual + CLI builders** - Design without coding  
✅ **Framework-agnostic** - Works with Bubble Tea, raw stdout, any TUI  
✅ **Pre-built library** - Quality characters ready to use  
✅ **Zero core dependencies** - Pure Go stdlib for core functionality  
✅ **AI-friendly** - Non-interactive CLI for agents  

---

## Install

**One command for everyone:**

```bash
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash
```

This installs the `tangent` CLI tool to `~/.local/bin`.

**For Go developers:** To use the package in your code, just import it:

```go
import "github.com/wildreason/tangent/pkg/characters"
```

Then run `go mod tidy` in your project. No separate installation needed.

---

## Two Ways to Use

### 1. Simple CLIs - Built-in Animation

For standalone tools, use the built-in animation engine:

```go
import "github.com/wildreason/tangent/pkg/characters"

alien, _ := characters.Library("alien")
characters.Animate(os.Stdout, alien, 5, 3)  // Done!
```

### 2. Complex TUIs - Frame Extraction

For Bubble Tea and other frameworks, extract frames:

```go
import (
    "github.com/wildreason/tangent/pkg/characters"
    "github.com/wildreason/tangent/pkg/adapters/bubbletea"
)

// One-line Bubble Tea integration
spinner, _ := bubbletea.LibrarySpinner("alien", 5)

// Or extract frames manually
alien, _ := characters.Library("alien")
frames := characters.ExtractFrames(alien)
// Use frames in your TUI framework
```

**📘 Bubble Tea:** See [`docs/BUBBLETEA_INTEGRATION.md`](docs/BUBBLETEA_INTEGRATION.md)

---

## Quick Start

### Design with Visual Builder

```bash
tangent  # Opens interactive builder
```

- Design frame by frame with visual palette
- Preview animation live
- Export Go code
- Save/load sessions

### Design with CLI (AI Agents)

```bash
tangent create --name robot --width 11 --height 3 \
  --frame idle '__R6FFF6L__,_T5FFFFF6T_,___11_22___' \
  --output robot.go --package agent
```

**📘 AI Agents:** See [`AGENTS-README.md`](AGENTS-README.md)

### Browse Library

```bash
tangent gallery  # See all pre-built characters
```

---

## Pattern System

Design characters with single-character codes:

| Code | Block | Code | Block | Code | Block |
|------|-------|------|-------|------|-------|
| `F` | █ Full | `1` | ▘ Quad UL | `.` | ░ Light |
| `T` | ▀ Top | `2` | ▝ Quad UR | `:` | ▒ Medium |
| `B` | ▄ Bottom | `3` | ▖ Quad LL | `#` | ▓ Dark |
| `L` | ▌ Left | `4` | ▗ Quad LR | `_` | Space |
| `R` | ▐ Right | `5` | ▛ 3-Quad | `X` | Mirror |

**Example pattern:**
```
'R6FFF6L,T5FFF6T,_1_2_'  // 3 lines: right-body-left, top-body-top, eyes
```

**Full guide:** [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md)

---

## Library Characters

Pre-built, ready-to-use characters:

- **alien** (7x3, 3 frames) - Waving hands animation
- **pulse** (9x5, 3 frames) - Heartbeat/thinking indicator
- **wave** (11x5, 5 frames) - Friendly greeting bot
- **rocket** (7x7, 4 frames) - Launch sequence

```bash
tangent gallery  # Browse with visual previews
```

```go
// Use in code
alien, _ := characters.Library("alien")
characters.Animate(os.Stdout, alien, 5, 3)
```

**Full library:** [`docs/LIBRARY.md`](docs/LIBRARY.md)

---

## Integration Examples

### Bubble Tea (Recommended for TUIs)

```go
import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/wildreason/tangent/pkg/adapters/bubbletea"
)

type model struct {
    spinner spinner.Model
}

func initialModel() model {
    s, _ := bubbletea.LibrarySpinner("wave", 6)
    return model{spinner: s}
}

func (m model) View() string {
    return m.spinner.View()  // Bubble Tea controls rendering
}
```

**See:** [`examples/bubbletea/`](examples/bubbletea/)

### Simple CLI

```go
import "github.com/wildreason/tangent/pkg/characters"

func main() {
    pulse, _ := characters.Library("pulse")
    characters.Animate(os.Stdout, pulse, 8, 5)
}
```

**See:** [`examples/demo/`](examples/demo/)

---

## CLI Commands

### `tangent` (interactive)
Visual character designer with:
- Pattern palette
- Frame-by-frame editor
- Live preview
- Export options

### `tangent create`
Create character from CLI:
```bash
tangent create --name bot --width 7 --height 3 \
  --frame idle 'R6F6L,T5F6T,_1_2_' \
  --output bot.go
```

### `tangent animate`
Preview animations:
```bash
tangent animate --name alien --fps 5 --loops 3
```

### `tangent gallery`
Browse library characters:
```bash
tangent gallery
```

### `tangent export`
Export saved sessions:
```bash
tangent export --session mychar --output mychar.go --package agent
```

**Full CLI guide:** `tangent help`

---

## Frame Extraction API

For TUI frameworks, extract frames without animation:

```go
import "github.com/wildreason/tangent/pkg/characters"

char, _ := characters.Library("alien")

// Get all frames
frames := char.GetFrames()
for _, frame := range frames {
    fmt.Println(frame.Content)  // Multi-line string
    fmt.Println(frame.Lines)    // []string for easy manipulation
}

// Normalize frames (prevent jitter)
normalized := char.Normalize()

// Quick extraction for spinners
spinnerFrames := characters.ToSpinnerFrames(char)

// Frame statistics
stats := char.Stats()
fmt.Printf("Frames: %d, Size: %dx%d\n", 
    stats.TotalFrames, stats.Width, stats.Height)
```

---

## Adapters

### Bubble Tea Adapter

```go
import "github.com/wildreason/tangent/pkg/adapters/bubbletea"

// One-liner: library character → spinner
s, _ := bubbletea.LibrarySpinner("wave", 6)

// Or with custom character
s := bubbletea.SpinnerFromCharacter(myChar, 5)

// Normalized (no jitter)
s := bubbletea.NormalizedSpinner(myChar, 5)

// Batch create multiple
spinners := bubbletea.MultiCharacterSpinners(chars, 5)
```

**Documentation:** [`docs/BUBBLETEA_INTEGRATION.md`](docs/BUBBLETEA_INTEGRATION.md)

---

## Features

- **Character Design**
  - Pattern-based system (F, T, B, L, R, 1-8, etc.)
  - Visual builder (interactive mode)
  - CLI builder (for AI agents)
  - Mirror support for symmetry

- **Library**
  - Pre-built characters (alien, pulse, wave, rocket)
  - Gallery command with visual previews
  - Easily extensible

- **Integration**
  - Works with any TUI framework
  - Bubble Tea adapter included
  - Frame extraction API
  - Built-in animation for simple CLIs

- **Tooling**
  - Session management
  - Multi-frame editing
  - Live preview
  - Export to Go code

- **Quality**
  - Frame normalization (no jitter)
  - Consistent dimensions
  - Zero core dependencies
  - Well-tested

---

## Use Cases

| Use Case | Integration | Example |
|----------|-------------|---------|
| **Simple CLI tools** | Built-in animation | Loading indicator |
| **Bubble Tea apps** | Adapter | TUI with animated characters |
| **AI Agents** | CLI commands | Generate characters on-the-fly |
| **Terminal games** | Frame extraction | Sprite animations |
| **DevOps tools** | Library characters | Build/deploy feedback |
| **Status indicators** | Pulse/rocket | Background processing |

---

## Examples

- [`examples/demo/`](examples/demo/) - Library + custom characters
- [`examples/bubbletea/`](examples/bubbletea/) - Bubble Tea integration
- [`examples/tokyo/`](examples/tokyo/) - Custom character export

```bash
cd examples/demo
go run main.go
```

---

## Documentation

- [`AGENTS-README.md`](AGENTS-README.md) - **Complete guide for AI agents**
- [`docs/BUBBLETEA_INTEGRATION.md`](docs/BUBBLETEA_INTEGRATION.md) - **Bubble Tea integration**
- [`docs/PATTERN_GUIDE.md`](docs/PATTERN_GUIDE.md) - Pattern codes reference
- [`docs/LIBRARY.md`](docs/LIBRARY.md) - Pre-built characters
- [`CONTRIBUTING.md`](CONTRIBUTING.md) - How to contribute
- [`CHANGELOG.md`](CHANGELOG.md) - Version history
- [`ROADMAP.md`](ROADMAP.md) - Future plans

---

## Architecture

```
┌─────────────────────────────────────────────────┐
│ Your TUI Framework (Bubble Tea, etc)            │
│ • Event loop                                     │
│ • Rendering                                      │
│ • Layout (Lip Gloss)                             │
├─────────────────────────────────────────────────┤
│ Tangent Adapter (optional)                       │
│ • Frame extraction                               │
│ • Framework-specific helpers                     │
├─────────────────────────────────────────────────┤
│ Tangent Core                                     │
│ • Character design system                        │
│ • Pattern compiler                               │
│ • Library characters                             │
│ • Optional built-in animation (simple CLIs)      │
└─────────────────────────────────────────────────┘
```

**Tangent = Character Design System**  
**Your Framework = Render Engine**  
**Clear separation of concerns**

---

## Requirements

- **Core:** Go 1.21+ only
- **Bubble Tea adapter:** Adds charmbracelet dependencies
- **Terminal:** Unicode block element support

---

## API Reference

### Character Design
```go
spec := characters.NewCharacterSpec("name", width, height).
    AddFrame("idle", []string{"pattern..."}).
    AddFrame("move", []string{"pattern..."})

char, err := spec.Build()
```

### Frame Extraction
```go
frames := char.GetFrames()           // Full metadata
lines := char.GetFrameLines()        // [][]string
spinnerFrames := char.ToSpinnerFrames()  // []string
normalized := char.Normalize()       // No jitter
```

### Built-in Animation (Simple CLIs)
```go
characters.Animate(os.Stdout, char, fps, loops)
characters.ShowIdle(os.Stdout, char)
```

### Library
```go
char, _ := characters.Library("alien")
names := characters.ListLibrary()
info, _ := characters.LibraryInfo("alien")
```

### Bubble Tea Adapter
```go
s, _ := bubbletea.LibrarySpinner("wave", 6)
s := bubbletea.SpinnerFromCharacter(char, 5)
frames := bubbletea.FramesFromCharacter(char)
```

---

## Contributing

Contributions welcome! See [`CONTRIBUTING.md`](CONTRIBUTING.md).

**Want to add a character to the library?**
1. Design with `tangent`
2. Export your design
3. Submit a PR

---

## License

MIT License © 2025 Wildreason, Inc

See [`LICENSE`](LICENSE) for details.

---

## Links

- **Repository:** https://github.com/wildreason/tangent
- **Issues:** https://github.com/wildreason/tangent/issues
- **Releases:** https://github.com/wildreason/tangent/releases

---

**Built for terminal developers**  
**Designed for the Charmbracelet ecosystem**  
**Works anywhere**
