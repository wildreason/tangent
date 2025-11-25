# Tangent API

Go library for terminal avatars in AI agents.

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

## Quick Start

```go
package main

import (
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    // Load agent
    agent, err := characters.LibraryAgent("sam")
    if err != nil {
        panic(err)
    }

    // Show states
    agent.Think(os.Stdout)
    agent.Execute(os.Stdout)
}
```

## Core API

### Load Agent

```go
agent, err := characters.LibraryAgent(name)
```

**Available characters**: sam, rio, ga, ma, pa, da, ni

### State Methods

Standard agent states:

```go
agent.Plan(writer)      // Planning
agent.Think(writer)     // Thinking
agent.Execute(writer)   // Executing
agent.Wait(writer)      // Waiting/idle
agent.Error(writer)     // Error state
```

Custom states:

```go
agent.ShowState(writer, "arise")      // Awakening
agent.ShowState(writer, "read")       // Reading
agent.ShowState(writer, "write")      // Writing
agent.ShowState(writer, "search")     // Searching
agent.ShowState(writer, "approval")   // Approval
agent.ShowState(writer, "bash")       // Running bash
agent.ShowState(writer, "build")      // Building
agent.ShowState(writer, "communicate")// Communicating
agent.ShowState(writer, "block")      // Blocking
agent.ShowState(writer, "blocked")    // Blocked
agent.ShowState(writer, "resting")    // Resting/complete
```

### Introspection

```go
// List all states
states := agent.ListStates()  // []string

// Check if state exists
has := agent.HasState("think")  // bool

// Get character
char := agent.GetCharacter()  // *domain.Character
```

## Themes

Global color themes for all agents.

### Set Theme

```go
// Set theme (applies to all agents created after)
err := characters.SetTheme("latte")  // Default
err := characters.SetTheme("bright")
err := characters.SetTheme("garden")
err := characters.SetTheme("cozy")
```

### Get Theme

```go
// Get current theme
current := characters.GetCurrentTheme()  // "latte"

// List all themes
themes := characters.ListThemes()  // ["bright", "cozy", "garden", "latte"]
```

### Theme Colors

**latte** (default) - Catppuccin-inspired warm pastels, GUI-friendly
```
sam: #E78284  rio: #EF9F76  ga: #E5C890  ma: #A6D189
pa: #85C1DC  da: #CA9EE6  ni: #F4B8E4
```

**bright** - Original bright colors, 100% saturation
```
sam: #FF0000  rio: #FF8800  ga: #FFD700  ma: #00FF00
pa: #0088FF  da: #8800FF  ni: #FF0088
```

**garden** - Earthy natural colors, reduces terminal intimidation
```
sam: #D4787D  rio: #D89C6A  ga: #C9B68C  ma: #8FB378
pa: #7CA8B8  da: #A888BA  ni: #C895A8
```

**cozy** - Modern GUI hybrid with professional warmth
```
sam: #E18B8B  rio: #E5A679  ga: #E6CC94  ma: #99C794
pa: #78AED4  da: #B592D4  ni: #DE99B8
```

## Advanced Usage

### Bubble Tea Integration (Recommended)

Use the pre-built Bubble Tea component for plug-and-play animation:

```go
import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/wildreason/tangent/pkg/characters"
    "github.com/wildreason/tangent/pkg/characters/bubbletea"
)

// Load character
agent, _ := characters.LibraryAgent("sam")

// Create animated component (10 FPS)
char := bubbletea.NewAnimatedCharacter(agent, 100*time.Millisecond)

// Set initial state
char.SetState("plan")

// Use in Bubble Tea program
program := tea.NewProgram(char)
program.Run()
```

**AnimatedCharacter API:**

```go
// State management
char.SetState("think")          // Change state
char.GetState()                 // Get current state
char.ListStates()               // List available states

// Animation control
char.Play()                     // Start animation
char.Pause()                    // Pause animation
char.IsPlaying()                // Check if playing
char.Reset()                    // Reset to first frame

// Configuration
char.SetTickInterval(200*time.Millisecond)  // Change speed
char.GetTickInterval()                      // Get current speed
char.GetWidth()                             // Get character width
char.GetHeight()                            // Get character height
```

### Frame Cache API (Performance)

Pre-render all frames for O(1) access during animation:

```go
agent, _ := characters.LibraryAgent("sam")

// Get frame cache (pre-rendered and pre-colored)
cache := agent.GetFrameCache()

// Get base frame ([]string)
baseLines := cache.GetBaseFrame()
for _, line := range baseLines {
    fmt.Println(line)  // Already compiled and colored
}

// Get state frames ([][]string)
planFrames := cache.GetStateFrames("plan")
for frameIdx, frame := range planFrames {
    for _, line := range frame {
        fmt.Println(line)  // Already compiled and colored
    }
}

// Cache introspection
cache.HasState("think")         // Check state exists
cache.ListStates()              // List all states
cache.GetCharacterName()        // Get character name
cache.GetColor()                // Get hex color
```

**Performance benefit:** Pre-rendering eliminates pattern compilation and colorization during animation, reducing CPU usage during 60 FPS animations.

### Custom TUI Integration

For custom TUI frameworks (tview, etc.):

```go
import "github.com/wildreason/tangent/pkg/characters/domain"

// Option 1: Use Frame Cache (recommended)
cache := agent.GetFrameCache()
frames := cache.GetStateFrames("think")
// frames[i] is []string (pre-colored lines)

// Option 2: Manual frame access
char := agent.GetCharacter()
state := char.States["think"]
frames := state.Frames  // []domain.Frame

// Colorize frames manually
for _, frame := range frames {
    coloredLines := characters.ColorizeFrame(frame, char.Color)
    // Use coloredLines in your TUI
}
```

### Frame Extraction

```go
// Extract single frame as string
frame := char.States["think"].Frames[0]
coloredLines := characters.ColorizeFrame(frame, "#FF0000")
output := strings.Join(coloredLines, "\n")
```

## Examples

### AI Agent Workflow

```go
agent, _ := characters.LibraryAgent("sam")

agent.Wait(os.Stdout)           // Idle
agent.Think(os.Stdout)          // Processing
agent.Read(os.Stdout)           // Reading input
agent.Execute(os.Stdout)        // Running task
agent.ShowState(os.Stdout, "resting")  // Complete
```

### Error Handling

```go
agent, err := characters.LibraryAgent("unknown")
if err != nil {
    log.Fatalf("Failed to load agent: %v", err)
}

if !agent.HasState("custom_state") {
    agent.Wait(os.Stdout)  // Fallback to standard state
}
```

### Theme Configuration

```go
// Set theme at startup
func init() {
    characters.SetTheme("latte")  // Default is already latte
}

// Or allow user config
if cfg.Theme != "" {
    if err := characters.SetTheme(cfg.Theme); err != nil {
        log.Printf("Invalid theme %q, using default", cfg.Theme)
    }
}
```

## Reference

**7 characters** × **16 states** × **4 themes** = 448 combinations

**States**: arise, wait, think, plan, execute, error, read, search, write, bash, build, communicate, block, blocked, resting, approval

**Themes**: latte (default), bright, garden, cozy

**License**: MIT © 2025 Wildreason, Inc
