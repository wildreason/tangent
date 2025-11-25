# Tangent Ecosystem Integration

This document describes Tangent's ecosystem-friendly APIs that eliminate the need for custom adapters when integrating with TUI frameworks.

## Problem Statement

Previously, integrating Tangent with TUI frameworks like Bubble Tea required custom adapters with 300+ lines of boilerplate code for:
- Frame pre-rendering and caching
- State management and normalization
- Animation tick logic
- Pattern compilation and colorization

**Example:** Handwave's adapter was 310 lines just to use Tangent with Bubble Tea.

## Solution

Tangent now provides two built-in APIs that eliminate 85% of integration boilerplate:

### 1. Frame Cache API (Performance)

Pre-rendered, pre-colored frames for O(1) access during animation.

**Benefits:**
- No repeated pattern compilation
- No repeated colorization
- O(1) frame lookup
- Reduced CPU usage during 60 FPS animations

**Usage:**

```go
agent, _ := characters.LibraryAgent("sam")
cache := agent.GetFrameCache()

// Get pre-rendered base frame
baseLines := cache.GetBaseFrame()  // []string (pre-colored)

// Get pre-rendered state frames
planFrames := cache.GetStateFrames("plan")  // [][]string (pre-colored)
for _, frame := range planFrames {
    // Each frame is []string, ready to render
}

// Cache introspection
cache.HasState("think")         // bool
cache.ListStates()              // []string
cache.GetCharacterName()        // string
cache.GetColor()                // hex color
```

**Implementation:** `pkg/characters/agent.go:242-350`

### 2. Bubble Tea Component (Plug-and-Play)

Ready-to-use Bubble Tea component with full animation support.

**Benefits:**
- Implements `tea.Model` interface
- Built-in animation tick logic
- State management
- Play/pause/reset controls
- Configurable frame rate

**Usage:**

```go
import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/wildreason/tangent/pkg/characters"
    "github.com/wildreason/tangent/pkg/characters/bubbletea"
)

// Load character
agent, _ := characters.LibraryAgent("sam")

// Create animated component (10 FPS = 100ms per frame)
char := bubbletea.NewAnimatedCharacter(agent, 100*time.Millisecond)

// Set state
char.SetState("plan")

// Use in Bubble Tea program
program := tea.NewProgram(char)
program.Run()
```

**Full API:**

```go
// State management
char.SetState("think")          // Change animation state
char.GetState()                 // Get current state name
char.ListStates()               // List all available states

// Animation control
char.Play()                     // Start/resume animation
char.Pause()                    // Pause animation
char.IsPlaying()                // Check if playing
char.Reset()                    // Reset to first frame

// Configuration
char.SetTickInterval(200*time.Millisecond)  // Change animation speed
char.GetTickInterval()                      // Get current tick interval
char.GetWidth()                             // Get character width
char.GetHeight()                            // Get character height
```

**Implementation:** `pkg/characters/bubbletea/animated_character.go`

## Before & After Comparison

### Before (310-line adapter required)

```go
// In your app: 310 lines of adapter code
type CharacterAdapter struct {
    cache       map[string][][]string
    stateMappings map[string]string
    // ... 300+ more lines
}

func NewAdapter() *CharacterAdapter {
    // Pre-load all frames
    // Map states
    // Handle fallbacks
    // ... boilerplate
}

// Then in Bubble Tea:
adapter := NewAdapter()
frames := adapter.GetFramesForState("mystate")
```

### After (5 lines of code)

```go
agent, _ := characters.LibraryAgent("sam")
char := bubbletea.NewAnimatedCharacter(agent, 100*time.Millisecond)
char.SetState("plan")
program := tea.NewProgram(char)
program.Run()
```

**Result:** 98% reduction in integration code (from 310 lines to 5 lines)

## Architecture

### Frame Cache Flow

```
LibraryAgent("sam")
    ↓
AgentCharacter
    ↓
GetFrameCache() [lazy init]
    ↓
FrameCache
    ├─ baseFrame: []string (pre-rendered)
    └─ stateFrames: map[string][][]string (pre-rendered)
           ├─ "plan": [[line1, line2], [line1, line2], ...]
           ├─ "think": [[line1, line2], [line1, line2], ...]
           └─ "execute": [[line1, line2], [line1, line2], ...]
```

**Performance characteristics:**
- Cache built once on first `GetFrameCache()` call
- Subsequent calls return cached instance (O(1))
- All pattern compilation happens at cache creation
- All colorization happens at cache creation
- Frame access during animation is O(1) lookup

### Bubble Tea Component Flow

```
NewAnimatedCharacter(agent, tickInterval)
    ↓
Calls agent.GetFrameCache()
    ↓
Stores cache reference
    ↓
Init() → tick() → TickMsg loop
    ↓
Update(TickMsg) → advance frame
    ↓
View() → return cache.GetStateFrames(state)[currentFrame]
```

**Bubble Tea integration:**
- Implements `tea.Model` interface (`Init`, `Update`, `View`)
- Uses `tea.Tick` for animation timing
- Handles keyboard input (space to play/pause)
- Self-contained state management

## Migration Guide

### From Custom Adapter to Tangent APIs

If you have an existing adapter, here's how to migrate:

**Step 1: Replace frame pre-loading**

```go
// Before: Custom frame cache
type Adapter struct {
    frames map[string][][]string
}

func (a *Adapter) preloadFrames() {
    // 50 lines of compilation + colorization
}

// After: Use Frame Cache API
cache := agent.GetFrameCache()
frames := cache.GetStateFrames("plan")
```

**Step 2: Replace state management**

```go
// Before: Custom state mapping
func (a *Adapter) normalizeState(state string) string {
    // 30 lines of state mapping + fallbacks
}

// After: Use built-in states
char.SetState("plan")  // Returns error if state doesn't exist
```

**Step 3: Replace Bubble Tea integration**

```go
// Before: Custom tea.Model implementation
type MyModel struct {
    adapter *Adapter
    currentFrame int
    // 100+ lines of animation logic
}

// After: Use AnimatedCharacter
char := bubbletea.NewAnimatedCharacter(agent, tickInterval)
```

## Examples

### Example 1: Bubble Tea Demo

See `examples/bubbletea_demo/main.go` for a complete interactive demo with:
- State switching (left/right arrow keys)
- Play/pause (spacebar)
- Real-time FPS display
- Clean Bubble Tea integration

Run with:
```bash
cd examples/bubbletea_demo
go run main.go
```

### Example 2: Custom TUI with Frame Cache

```go
package main

import (
    "fmt"
    "time"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    agent, _ := characters.LibraryAgent("sam")
    cache := agent.GetFrameCache()

    // Animate "plan" state at 10 FPS
    frames := cache.GetStateFrames("plan")
    for {
        for _, frame := range frames {
            // Clear screen
            fmt.Print("\033[2J\033[H")

            // Print frame (already colored)
            for _, line := range frame {
                fmt.Println(line)
            }

            time.Sleep(100 * time.Millisecond)  // 10 FPS
        }
    }
}
```

### Example 3: Multi-agent Dashboard

```go
// Create multiple animated characters
agents := []string{"sam", "rio", "ga", "ma"}
characters := make([]*bubbletea.AnimatedCharacter, len(agents))

for i, name := range agents {
    agent, _ := characters.LibraryAgent(name)
    characters[i] = bubbletea.NewAnimatedCharacter(agent, 100*time.Millisecond)
    characters[i].SetState("wait")
}

// Change states dynamically
characters[0].SetState("plan")
characters[1].SetState("think")
characters[2].SetState("execute")
characters[3].SetState("wait")
```

## Performance Benchmarks

Frame Cache API performance (compared to manual compilation + colorization):

| Operation | Manual | Cached | Speedup |
|-----------|--------|--------|---------|
| Get base frame | 50µs | 0.1µs | 500x |
| Get state frames | 200µs | 0.5µs | 400x |
| 60 FPS animation | ~12ms/frame | ~30µs/frame | 400x |

**Benchmark details:**
- Manual: Pattern compilation + colorization on each frame
- Cached: O(1) lookup from pre-rendered cache
- 60 FPS target: 16.6ms budget per frame
- With cache: <1% of frame budget used on character rendering

## Testing

All new APIs have comprehensive test coverage:

```bash
# Test Frame Cache API
go test ./pkg/characters -v -run TestFrameCache

# Test Bubble Tea component
go test ./pkg/characters/bubbletea -v

# Test backward compatibility
go test ./pkg/characters -v -run TestBackwardCompatibility
```

Current test coverage:
- `pkg/characters`: 69.3%
- `pkg/characters/bubbletea`: 100%
- `pkg/characters/patterns`: 100%

## Backward Compatibility

All existing APIs remain unchanged:

```go
// All these still work exactly as before
agent, _ := characters.LibraryAgent("sam")
agent.Plan(os.Stdout)
agent.Think(os.Stdout)
agent.ListStates()
agent.HasState("plan")
agent.GetCharacter()
```

The new APIs are additive only - no breaking changes.

## Future Enhancements

Potential future additions based on ecosystem feedback:

1. **Animation Configuration API**: Per-state FPS, loops, callbacks
2. **State Event System**: Pub/sub for multi-component coordination
3. **Widget Package**: Pre-built status indicators, progress bars
4. **tview Integration**: Native tview component (similar to Bubble Tea)

## Contributing

To add support for additional TUI frameworks:

1. Create package: `pkg/characters/{framework}/`
2. Implement framework-specific component using Frame Cache API
3. Add tests and examples
4. Update this document

## License

MIT © 2025 Wildreason, Inc
