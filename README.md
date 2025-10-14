# Tangent

**Terminal Character Library for AI Agents**

Create expressive terminal characters for AI agents with state-based animations. Characters represent agent behaviors like planning, thinking, and executing through visual states.

---

## Quick Start

### Agent State API (Recommended)

**Use characters with agent states:**

```go
import "github.com/wildreason/tangent/pkg/characters"

// Get character with agent state API
agent, _ := characters.LibraryAgent("mercury")

// Use agent states
agent.Plan(os.Stdout)      // Planning animation
agent.Think(os.Stdout)     // Thinking animation
agent.Execute(os.Stdout)   // Execution animation
agent.Success(os.Stdout)   // Success animation
```

### Legacy API

**Simple animation (backward compatible):**

```go
// Get pre-built character (legacy API)
mercury, _ := characters.Library("mercury")

// Use it immediately
characters.Animate(os.Stdout, mercury, 5, 3)  // Done!
```

**That's it!** No configuration, no setup, no complexity.

---

## Agent States

Characters support agent behavioral states:

- **plan** - Agent analyzing and planning
- **think** - Agent processing information
- **execute** - Agent performing actions
- **wait** - Agent waiting for input
- **error** - Agent handling errors
- **success** - Agent celebrating success

**Example workflow:**

```go
agent, _ := characters.LibraryAgent("mercury")

agent.Wait(os.Stdout)      // Waiting for task
agent.Plan(os.Stdout)      // Analyzing task
agent.Think(os.Stdout)     // Processing solution
agent.Execute(os.Stdout)   // Performing action
agent.Success(os.Stdout)   // Task complete!
```

See [Agent States Documentation](docs/AGENT_STATES.md) for complete guide.

---

## Core API (3 Functions Only)

### 1. Get Library Character
```go
mercury, _ := characters.LibraryAgent("mercury")
venus, _ := characters.LibraryAgent("venus")
earth, _ := characters.LibraryAgent("earth")
```

### 2. Create Custom Character
```go
robot := characters.NewCharacterSpec("my-robot", 8, 4).
    AddFrame("idle", []string{"FRF", "LRL", "FRF", "LRL"}).
    Build()
```

### 3. Use Character
```go
mercury.Plan(os.Stdout)      // Agent state
mercury.Think(os.Stdout)     // Agent state
mercury.Execute(os.Stdout)   // Agent state
```

---

## Library Characters

**Planet Series** - Professional agent characters with state-based animations:

| Character | Personality | Use Case | Example |
|-----------|-------------|----------|---------|
| **mercury** | Efficient | Fast, direct agent | `characters.LibraryAgent("mercury")` |
| **venus** | Friendly | Warm, welcoming agent | `characters.LibraryAgent("venus")` |
| **earth** | Balanced | Versatile, all-purpose agent | `characters.LibraryAgent("earth")` |
| **mars** | Action-oriented | Dynamic, energetic agent | `characters.LibraryAgent("mars")` |
| **jupiter** | Powerful | Large-scale, commanding agent | `characters.LibraryAgent("jupiter")` |
| **saturn** | Analytical | Methodical, precise agent | `characters.LibraryAgent("saturn")` |
| **uranus** | Creative | Innovative, exploratory agent | `characters.LibraryAgent("uranus")` |
| **neptune** | Calm | Smooth, flowing agent | `characters.LibraryAgent("neptune")` |

**All characters include:**
- Base (idle) state
- Required states: `plan`, `think`, `execute`
- Optional states: `wait`, `error`, `success`

**Browse characters:**
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

### AI Agent Applications
```go
// Agent workflow
mercury, _ := characters.LibraryAgent("mercury")
mercury.Plan(os.Stdout)
mercury.Think(os.Stdout)
mercury.Execute(os.Stdout)
mercury.Success(os.Stdout)
```

### TUI Applications
```go
// Extract state frames for your TUI framework
mercury, _ := characters.LibraryAgent("mercury")
state := mercury.GetCharacter().States["plan"]
for _, frame := range state.Frames {
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
│ • characters.LibraryAgent("mercury")            │
│ • agent.Plan() / Think() / Execute()            │
├─────────────────────────────────────────────────┤
│ Tangent Core                                    │
│ • Planet Series characters                      │
│ • State-based API                               │
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
    // Get character with agent states
    mercury, _ := characters.LibraryAgent("mercury")
    
    // Use agent states
    mercury.Plan(os.Stdout)
    mercury.Think(os.Stdout)
    mercury.Execute(os.Stdout)
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
    // Create custom character with CLI
    // tangent -> Create new character -> Design states
    
    // Or use Planet Series characters
    venus, _ := characters.LibraryAgent("venus")
    venus.Plan(os.Stdout)
    venus.Execute(os.Stdout)
    venus.Success(os.Stdout)
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
