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

### Custom TUI Integration

For Bubble Tea or custom TUI frameworks:

```go
import "github.com/wildreason/tangent/pkg/characters/domain"

// Get frames directly
char := agent.GetCharacter()
state := char.States["think"]
frames := state.Frames  // []domain.Frame

// Colorize frames
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
