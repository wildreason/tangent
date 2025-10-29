# Tangent Character Specification Contract

**Version:** 1.0.0
**Status:** Active
**Last Updated:** 2025-10-24

---

## Overview

This document defines the contract between **Tangent** (terminal character library) and its consumers, primarily **Handwave** (agent TUI framework).

**Core Principle:** Tangent provides reusable, state-aware ASCII art characters. Consumers request characters by name and state, Tangent delivers animation frames.

**Integration Architecture:**
```
Murmur (orchestration)
    ↓ emits states
Handwave (visualization) → see Handwave CONTRACT.md
    ↓ requests avatars
Tangent (characters) → THIS CONTRACT
```

**Related Contracts:**
- **[Handwave CONTRACT.md](https://github.com/wildreason/handwave/blob/main/CONTRACT.md)** - Murmur ↔ Handwave integration
- **This contract** - Tangent character specification for consumers

---

## 1. Responsibilities

### 1.1 Tangent Responsibilities (Character Provider)

**Character Library Management**
- ✅ Maintain library of terminal characters (ASCII art)
- ✅ Support 14 standard agent states per character
- ✅ Provide multi-frame animations for visual interest
- ✅ Use consistent pattern code system for colors
- ✅ Version and release character library as Go module

**API Stability**
- ✅ Maintain stable public API for character loading
- ✅ Support backward compatibility (don't remove characters)
- ✅ Follow semantic versioning for releases
- ✅ Document character availability and requirements

**Quality Standards**
- ✅ All characters MUST support 14 required states
- ✅ Frames MUST use valid pattern codes (R, L, 1-6, F, etc.)
- ✅ Dimensions MUST be consistent per character
- ✅ Animations MUST loop smoothly (first/last frame compatible)

---

### 1.2 Consumer Responsibilities (Handwave, etc.)

**Character Loading**
- ✅ Request characters via `LibraryAgent(name)` API
- ✅ Handle missing characters gracefully (fallback to default)
- ✅ Cache loaded characters for performance
- ✅ Respect character dimensions (don't resize)

**State Mapping**
- ✅ Map application states to Tangent's 14 standard states
- ✅ Handle missing state gracefully (use fallback)
- ✅ Request frames via state name (string)

**Attribution**
- ✅ Credit Tangent in documentation/about screens
- ✅ Link to Tangent repository in README
- ✅ Preserve author metadata from characters

---

## 2. API Contract

### 2.1 Public API

**Package:** `github.com/wildreason/tangent/pkg/characters`

**Primary Entry Point:**
```go
// LibraryAgent loads a character from the built-in library
func LibraryAgent(name string) (*AgentCharacter, error)
```

**Character Access:**
```go
type AgentCharacter struct {
    // GetCharacter returns the underlying character definition
    GetCharacter() *domain.Character
}

type Character struct {
    Name        string
    Description string
    Author      string
    BaseFrame   domain.Frame
    States      map[string]domain.CharacterState
}

type CharacterState struct {
    Name        string
    Description string
    Frames      []domain.Frame
}

type Frame struct {
    Lines []string  // Each line is pattern-coded ASCII art
}
```

**Example Usage:**
```go
// Load fire character
agent, err := characters.LibraryAgent("fire")
if err != nil {
    log.Fatal(err)
}

// Get character definition
char := agent.GetCharacter()

// Access state
waitState := char.States["wait"]
firstFrame := waitState.Frames[0]

// Render frame
for _, line := range firstFrame.Lines {
    fmt.Println(line)
}
```

---

### 2.2 Character Library

**Current Characters (v1.0.0):**

| Name | Dimensions | States | Description |
|------|------------|--------|-------------|
| `fire` | 11x3 | 14 | Orange/red flame avatar |

**Planned Characters (v1.1.0+):**

| Name | Dimensions | States | Color Theme | Agent Mapping |
|------|------------|--------|-------------|---------------|
| `fire` | 11x3 | 14 | Orange/red | `sa` (default) |
| `mercury` | 11x3 | 14 | Silver/white | `ri` |
| `neptune` | 11x3 | 14 | Cyan/blue | `ga` |
| `mars` | 11x3 | 14 | Red/crimson | `ma` |
| `jupiter` | 11x3 | 14 | Gold/yellow | `pa` |
| `saturn` | 11x3 | 14 | Purple/violet | `da` |
| `uranus` | 11x3 | 14 | Teal/aqua | `ni` |

---

## 3. State Specification

### 3.1 Required States

Every character MUST provide these 14 states:

**Core States (6):**
| State | Purpose | Animation Style |
|-------|---------|-----------------|
| `wait` | Agent idle | Gentle pulse/breathing |
| `think` | Agent reasoning | Focused intensity |
| `plan` | Agent planning | Organized movement |
| `execute` | Agent acting | Dynamic action |
| `error` | Agent failed | Alert/warning indication |
| `success` | Agent completed | Celebration/confirmation |

**Tool-Specific States (8):**
| State | Purpose | Animation Style |
|-------|---------|-----------------|
| `read` | Reading files | Scanning motion |
| `search` | Searching code | Sweeping/seeking |
| `write` | Writing files | Creating/building |
| `bash` | Running commands | Executing/processing |
| `build` | Compiling/testing | Constructing |
| `communicate` | Agent messaging | Transmitting/sending |
| `block` | Blocking operation | Waiting/synchronizing |
| `blocked` | Waiting on dependency | Paused/holding |

---

### 3.2 Animation Requirements

**Frame Count:**
- Minimum: 1 frame per state (static)
- Recommended: 3-10 frames per state (animated)
- Maximum: No limit (but consider performance)

**Frame Timing:**
- Consumers typically render at 10 FPS (100ms per frame)
- Design animations to loop smoothly
- First and last frame should transition well

**Example:**
```go
// Good: Smooth loop (3 frames)
wait_1 → wait_2 → wait_3 → wait_1 (loops seamlessly)

// Bad: Jarring transition
wait_1 (dim) → wait_2 (bright) → wait_1 (dim = sudden jump)
```

---

## 4. Pattern Code System

### 4.1 Color Codes

Tangent uses single-character pattern codes for colors:

**Standard Codes:**
```
R = Red
L = Blue
F = Bright/white
1 = Dark shade 1
2 = Dark shade 2
3 = Medium shade 3
4 = Medium shade 4
5 = Medium shade 5
6 = Light shade 6
_ = Space/empty
```

**Color Palette per Character:**
Each character defines its own color mapping:

```go
// Example: Fire character
Palette: map[rune]string{
    'R': "#FF4500", // Orange-red
    'L': "#FF6347", // Tomato
    'F': "#FFD700", // Gold
    '1': "#8B0000", // Dark red
    '2': "#A52A2A", // Brown
    // ...
}
```

---

### 4.2 Pattern Code Examples

**Fire Character (11x3):**
```
__R5FFF6L__
_26FFFFF51_
___11_22___
```

**Decoded:**
- Line 1: `__R5FFF6L__` = 2 spaces, red, shade5, 3x bright, shade6, blue, 2 spaces
- Line 2: `_26FFFFF51_` = space, shade2, shade6, 5x bright, shade5, shade1, space
- Line 3: `___11_22___` = 3 spaces, 2x shade1, space, 2x shade2, 3 spaces

**Mercury Character (silver/white - future):**
```
__F5FFF5F__
_46FFFFF64_
___66_66___
```

---

## 5. Character Creation Guidelines

### 5.1 Dimensions

**Standard Sizes:**
- **Small:** 11x3 (current standard)
- **Medium:** 20x5 (future)
- **Large:** 30x8 (future)

**Consistency:**
- All frames within a character MUST have same dimensions
- All states within a character MUST have same dimensions
- Characters of same "size class" SHOULD use same dimensions

---

### 5.2 Creating a New Character

**Step 1: Define Character Metadata**
```go
var mercuryCharacter = LibraryCharacter{
    Name:        "mercury",
    Description: "mercury - AI agent avatar (14 states)",
    Author:      "Wildreason, Inc",
    Width:       11,
    Height:      3,
    Patterns:    []Frame{ /* ... */ },
}
```

**Step 2: Create Base Frame**
```go
{
    Name: "base",
    Lines: []string{
        "__F5FFF5F__",
        "_46FFFFF64_",
        "___66_66___",
    },
},
```

**Step 3: Create All 14 States**
```go
// Wait state (3 frames for animation)
{
    Name: "wait_1",
    Lines: []string{
        "__F5FFF5F__",
        "_46FFFFF64_",
        "___66_66___",
    },
},
{
    Name: "wait_2",
    Lines: []string{
        "__F5FFFF5__",
        "_46FFFFF64_",
        "___66_66___",
    },
},
{
    Name: "wait_3",
    Lines: []string{
        "__FFFFF5F__",
        "_46FFFFF64_",
        "___66_66___",
    },
},

// Repeat for: plan, think, execute, error, success,
//             read, search, write, bash, build,
//             communicate, block, blocked
```

**Step 4: Register Character**
```go
// In library.go
func init() {
    register(mercuryCharacter)
}
```

---

### 5.3 Testing Checklist

Before submitting a new character:

- [ ] All 14 states implemented
- [ ] Each state has at least 1 frame
- [ ] Dimensions consistent across all frames
- [ ] Pattern codes valid (R, L, F, 1-6, _)
- [ ] Animations loop smoothly (first ↔ last frame)
- [ ] Metadata complete (name, description, author)
- [ ] Character registered in library
- [ ] Manual visual test completed (render each state)
- [ ] Added to character catalog in README

---

## 6. Versioning & Releases

### 6.1 Semantic Versioning

Tangent follows semantic versioning (MAJOR.MINOR.PATCH):

**MAJOR (v2.0.0):**
- Breaking API changes
- Remove/rename public functions
- Change character format (incompatible)

**MINOR (v1.1.0):**
- Add new characters to library
- Add new optional fields to Character struct
- New features (backward compatible)

**PATCH (v1.0.1):**
- Fix character rendering bugs
- Fix pattern code issues
- Documentation updates

---

### 6.2 Release Process

**Adding New Characters (MINOR version):**

1. Create character file (e.g., `mercury.go`)
2. Implement all 14 states
3. Add to library registry
4. Update CHANGELOG.md
5. Update README.md character catalog
6. Tag release: `git tag v1.1.0`
7. Push tag: `git push origin v1.1.0`

**Consumers update:**
```bash
go get github.com/wildreason/tangent@v1.1.0
```

---

## 7. Consumer Integration

### 7.1 Handwave Integration

**Design:** Handwave maintains an adapter layer to handle application-specific logic:

- **Tangent's responsibility:** Provide raw character data (art library)
- **Handwave's responsibility:** Application logic (fallbacks, legacy mapping, frame formatting)

**Handwave's adapter layer:**
```go
// internal/characters/adapter.go
func GetFramesForState(avatarName string, state State) ([]string, error) {
    // Load from Tangent
    agent, err := characters.LibraryAgent(avatarName)
    if err != nil {
        // Fallback to "fire" if requested avatar doesn't exist
        agent, err = characters.LibraryAgent("fire")
        if err != nil {
            return nil, fmt.Errorf("failed to load fallback avatar: %w", err)
        }
    }

    // Normalize legacy states (handwave's business logic)
    state = NormalizeLegacyState(state)  // idle → wait, progress → think

    // Get state frames
    char := agent.GetCharacter()
    stateData := char.States[string(state)]

    // Handle missing states with hierarchical fallback (handwave's policy)
    if stateData == nil {
        state = getFallbackState(state)  // read → think → plan → wait
        stateData = char.States[string(state)]
    }

    // Convert to string slices (handwave's format)
    frames := make([]string, len(stateData.Frames))
    for i, frame := range stateData.Frames {
        frames[i] = joinLines(frame.Lines)
    }

    return frames, nil
}
```

**Why separation?**
- Tangent stays reusable (no application-specific assumptions)
- Handwave adds policies for its specific use case
- Other consumers can implement their own adapter logic

**Handwave's go.mod:**
```go
require (
    github.com/wildreason/tangent v1.1.0
    github.com/charmbracelet/bubbletea v0.27.2
)
```

---

### 7.2 Other Potential Consumers

Tangent can serve any terminal UI application:

**Examples:**
- **Agent dashboards** - Show agent status with avatars
- **CLI tools** - Animated loading indicators
- **TUI applications** - Character-based interfaces
- **Terminal games** - Character sprites

**Integration pattern:**
```go
import "github.com/wildreason/tangent/pkg/characters"

// Load character
agent, _ := characters.LibraryAgent("fire")
char := agent.GetCharacter()

// Animate in your TUI
for {
    for _, frame := range char.States["wait"].Frames {
        renderFrame(frame)
        time.Sleep(100 * time.Millisecond)
    }
}
```

---

## 8. Performance Contract

### 8.1 Character Loading

**Tangent guarantees:**
- Character loading < 10ms (library lookup)
- Characters cached after first load
- No disk I/O after initialization (embedded in binary)

**Consumer best practices:**
- Load characters once at startup
- Cache AgentCharacter instances
- Pre-load all needed characters upfront

---

### 8.2 Memory Footprint

**Per Character:**
- ~14 states × 5 frames avg = 70 frames
- ~11x3 = 33 characters per frame
- ~70 × 33 = ~2.3 KB per character

**Full Library (7 characters):**
- ~7 × 2.3 KB = ~16 KB total
- Negligible memory impact

---

## 9. Examples

### 9.1 Loading and Rendering

**Basic Usage:**
```go
package main

import (
    "fmt"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    // Load fire character
    agent, err := characters.LibraryAgent("fire")
    if err != nil {
        panic(err)
    }

    char := agent.GetCharacter()
    fmt.Println("Character:", char.Name)
    fmt.Println("Author:", char.Author)

    // Render wait state (first frame)
    waitState := char.States["wait"]
    frame := waitState.Frames[0]

    for _, line := range frame.Lines {
        fmt.Println(line)
    }
}
```

**Output:**
```
Character: fire
Author: Wildreason, Inc
__R5FFF6L__
_26FFFFF51_
___11_22___
```

---

### 9.2 Animated Loop

**Animate a State:**
```go
package main

import (
    "fmt"
    "time"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    agent, _ := characters.LibraryAgent("fire")
    char := agent.GetCharacter()

    planState := char.States["plan"]

    // Loop forever
    for {
        for _, frame := range planState.Frames {
            // Clear screen (ANSI escape)
            fmt.Print("\033[H\033[2J")

            // Render frame
            for _, line := range frame.Lines {
                fmt.Println(line)
            }

            // 100ms per frame = 10 FPS
            time.Sleep(100 * time.Millisecond)
        }
    }
}
```

---

## 10. Character Catalog

### 10.1 Current Characters

**fire (v1.0.0)**
- Dimensions: 11x3
- States: 14
- Theme: Orange/red flames
- Animation: Flickering flame effect
- Best for: High-energy, active agents

---

### 10.2 Planned Characters (v1.1.0)

**mercury**
- Dimensions: 11x3
- States: 14
- Theme: Silver/white liquid metal
- Animation: Flowing mercury effect
- Best for: Fast, fluid agents

**neptune**
- Dimensions: 11x3
- States: 14
- Theme: Cyan/blue ocean waves
- Animation: Wave motion effect
- Best for: Deep-thinking, analytical agents

**mars**
- Dimensions: 11x3
- States: 14
- Theme: Red/crimson war theme
- Animation: Pulsing energy effect
- Best for: Aggressive, action-oriented agents

**jupiter**
- Dimensions: 11x3
- States: 14
- Theme: Gold/yellow storm clouds
- Animation: Swirling storm effect
- Best for: Powerful, commanding agents

**saturn**
- Dimensions: 11x3
- States: 14
- Theme: Purple/violet rings
- Animation: Orbital ring rotation
- Best for: Organized, systematic agents

**uranus**
- Dimensions: 11x3
- States: 14
- Theme: Teal/aqua ice crystals
- Animation: Crystalline shimmer
- Best for: Cool, methodical agents

---

## 11. Support & Contribution

### 11.1 Issues & Requests

**Bug Reports:**
- File at: https://github.com/wildreason/tangent/issues
- Include: Character name, state name, expected vs actual output

**Character Requests:**
- File feature request with character concept
- Include: Theme, color palette, animation ideas
- Community can vote on priority

---

### 11.2 Contributing Characters

**Submission Process:**

1. Fork tangent repository
2. Create character file (`pkg/characters/library/yourcharacter.go`)
3. Implement all 14 states
4. Test thoroughly (see checklist §5.3)
5. Submit PR with:
   - Character implementation
   - Updated README.md catalog
   - Example renders (screenshots)
   - CHANGELOG.md entry

**Review Criteria:**
- All 14 states present
- Visual quality (clear, recognizable)
- Animation smoothness
- Code quality (follows existing patterns)
- Documentation completeness

---

## 12. Future Roadmap

### 12.1 v1.1.0 - Seven Characters

**Goal:** Support murmur's 7-agent system

**Deliverables:**
- 6 new characters (mercury, neptune, mars, jupiter, saturn, uranus)
- All 11x3 dimensions
- All 14 states per character
- Color variants of fire base design

**Timeline:** 2-3 weeks

---

### 12.2 v1.2.0 - Character Customization

**Features:**
- Custom color palettes (user-defined)
- Character size variants (small/medium/large)
- State aliasing (map custom states to 14 standard)

---

### 12.3 v2.0.0 - Advanced Characters

**Features:**
- Multi-layer characters (foreground/background)
- Particle effects
- Sound integration (beep codes)
- Dynamic color shifting

---

## Appendix A: Complete State List

Reference for character creators:

```
1.  wait       - Agent idle, no task assigned
2.  think      - Agent analyzing/reasoning (LLM inference)
3.  plan       - Agent breaking down task into steps
4.  execute    - Agent running commands/taking action
5.  error      - Agent encountered failure
6.  success    - Agent completed task
7.  read       - Agent examining files/docs
8.  search     - Agent searching codebase (Grep/Glob)
9.  write      - Agent creating/editing files
10. bash       - Agent executing bash commands
11. build      - Agent compiling/testing
12. communicate - Agent messaging other agents (MCP)
13. block      - Agent blocking/synchronizing
14. blocked    - Agent waiting on dependency
```

---

**Contract Version:** 1.0.0
**Effective Date:** 2025-10-24
**Maintained by:** Wildreason, Inc
