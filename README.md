# Tangent

Terminal avatars for AI agents. Go library.

**v0.3.0:** TangentClient API for framework-agnostic animation control. Compact 8x2 micro avatars.

## Features

- **7 characters** × **16 animated states** × **4 color themes** = 448 combinations
- **TangentClient API** - Thread-safe animation controller for any TUI framework
- **Micro Avatars** - Compact 8x2 format for status bars
- **Frame Cache API** - Pre-rendered frames for O(1) access (500x faster)
- **Bubble Tea Component** - Built-in AnimatedCharacter (5 lines vs 310-line adapter)
- **State Aliases** - Built-in mappings (think->resting, bash->read, etc.)
- **Theme System** - Global color themes (latte, bright, garden, cozy)

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

## What's New in v0.3.0

### TangentClient API - Framework-Agnostic Animation

Thread-safe animation controller for tview, tcell, or any TUI framework:

```go
import "github.com/wildreason/tangent/pkg/characters/client"

tc, _ := client.NewMicro("sam")
tc.SetStateFPS("resting", 2)
tc.SetStateFPS("write", 8)
tc.Start()

// In render loop
frame := tc.GetFrameRaw()

// On state change - aliases handle mapping
tc.SetState("bash")  // automatically maps to "read" animation
```

**Features:**
- Per-state FPS configuration
- Built-in state aliases (think->resting, bash->read, etc.)
- Auto-tick background animation
- State queue and callbacks

### Compact Micro Avatars (8x2)

Micro avatars reduced from 10x2 to 8x2 for tighter layouts:

```go
tc, _ := client.NewMicro("sam")
w, h := tc.GetDimensions()  // 8, 2
```

See [CHANGELOG.md](CHANGELOG.md) for full v0.3.0 details.

## Usage

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sam")
agent.Plan(os.Stdout)
agent.Think(os.Stdout)
agent.Execute(os.Stdout)
```

## API

```go
// Get agent
agent, err := characters.LibraryAgent(name)

// State methods
agent.Plan(writer)
agent.Think(writer)
agent.Execute(writer)
agent.Wait(writer)
agent.Error(writer)

// Custom states
agent.ShowState(writer, "arise")
agent.ShowState(writer, "approval")

// Introspection
states := agent.ListStates()
hasState := agent.HasState("think")

// Themes - set color theme for all agents
characters.SetTheme("latte")        // Switch to latte theme
themes := characters.ListThemes()   // ["bright", "cozy", "garden", "latte"]
current := characters.GetCurrentTheme()  // "latte"

// Frame Cache API - pre-rendered frames for performance
cache := agent.GetFrameCache()
baseLines := cache.GetBaseFrame()      // []string (pre-colored)
planFrames := cache.GetStateFrames("plan")  // [][]string (pre-colored)
```

## TUI Framework Integration

### Bubble Tea (Plug-and-Play)

```go
import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/wildreason/tangent/pkg/characters"
    "github.com/wildreason/tangent/pkg/characters/bubbletea"
)

// Load character
agent, _ := characters.LibraryAgent("sam")

// Create animated component
char := bubbletea.NewAnimatedCharacter(agent, 100*time.Millisecond)
char.SetState("plan")

// Run in Bubble Tea
program := tea.NewProgram(char)
program.Run()
```

See [examples/bubbletea_demo](examples/bubbletea_demo/main.go) for full demo.

### Custom TUI Frameworks (tview, etc.)

```go
// Use Frame Cache for O(1) frame access
cache := agent.GetFrameCache()
frames := cache.GetStateFrames("think")
// frames[i] is []string (pre-colored, ready to render)
```

## Avatars

7 characters × 16 states × 4 themes

## States

arise, wait, think, plan, execute, error, read, search, write, bash, build, communicate, block, blocked, resting, approval

## Character Management

Characters are auto-generated from `pkg/characters/library/constants.go`.

### Renaming a Character

```bash
# 1. Edit constants.go
vim pkg/characters/library/constants.go
# Change: CharacterGa = "ga" → CharacterGa = "gabe"

# 2. Regenerate
make generate

# 3. Test and commit
make test
git add .
git commit -m "Rename ga to gabe"
```

All character files and theme mappings are updated automatically.

### Adding a Character

1. Add character constants to `constants.go`:
   - Name constant (e.g., `CharacterNew = "newchar"`)
   - Color constant (e.g., `ColorNew = "#ABCDEF"`)
   - Theme colors for all 4 themes
2. Run `make generate`
3. New character file created automatically

See [CODEGEN.md](pkg/characters/library/CODEGEN.md) for details.

## Documentation

- [API Reference](docs/API.md) - Complete API documentation
- [Ecosystem Guide](docs/ECOSYSTEM.md) - TUI framework integration
- [Code Generation](pkg/characters/library/CODEGEN.md) - Character management
- [CHANGELOG](CHANGELOG.md) - Version history

## Examples

- [Bubble Tea Demo](examples/bubbletea_demo/main.go) - Interactive animated demo

## Version

Current: **v0.3.0** (2025-12-11)
- v0.3.0: TangentClient API + Compact 8x2 Micro Avatars
- v0.2.0: Frame Cache API + Bubble Tea + Code Generation
- v0.1.1: Character name updates (sam, rio)
- v0.1.0: Initial stable release

## License

MIT © 2025 Wildreason, Inc
