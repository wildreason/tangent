# Tangent

Terminal avatars for AI agents. Go library.

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

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

## License

MIT © 2025 Wildreason, Inc
