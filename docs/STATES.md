# Agent States

Tangent characters support agent states, allowing AI agents to express their current behavioral state through visual representation.

## Overview

Agent states represent what an AI agent is doing at any given moment. Every character must have a minimum of 3 required states, with optional additional states for enhanced expressiveness.

## Required States

All characters **must** have these three core states:

### plan
**Purpose:** Show the agent analyzing a problem and formulating a strategy

**Visual Guidelines:**
- Use question marks (?) for inquiry
- Use dots (.) for analysis  
- Show contemplative posture
- Suggest forward thinking

**Frame Requirements:** Minimum 3 frames for animation

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   ?  ?    
```

### think
**Purpose:** Show the agent processing information and reasoning

**Visual Guidelines:**
- Use dots (...) for processing
- Use brain/thought symbols
- Show concentration
- Indicate mental activity

**Frame Requirements:** Minimum 3 frames for animation

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   . . .   
```

### execute
**Purpose:** Show the agent performing actions and implementing decisions

**Visual Guidelines:**
- Use action indicators (arrows →)
- Show movement or activity
- Indicate progress
- Display confidence

**Frame Requirements:** Minimum 3 frames for animation

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   → → →   
```

## Optional States

Characters may include these additional states for richer interaction:

### wait
**Purpose:** Show the agent in an idle or waiting state

**Visual Guidelines:**
- Minimal movement
- Blinking or breathing effects
- Calm, patient appearance
- Low-energy animation

### error
**Purpose:** Show the agent handling an error condition

**Visual Guidelines:**
- Use error symbols (✗, !)
- Show disruption or concern
- Alert colors/patterns
- Clear indication of problem state

### success
**Purpose:** Show the agent celebrating successful completion

**Visual Guidelines:**
- Use success symbols (✓, ★)
- Show joy or satisfaction
- Positive, energetic movement
- Celebration indicators

## Animation Properties

### Frame Count
- **Minimum:** 3 frames per state (required for animation)
- **Recommended:** 3-5 frames for smooth animation
- **Maximum:** No hard limit, but keep performance in mind

### Frame Rate (FPS)
- **Default:** 5 FPS
- **Range:** 1-30 FPS
- **Recommendation:** 4-8 FPS for terminal animations

### Loops
- **Default:** 1 loop
- **Range:** 1-∞ (infinite with 0)
- **Recommendation:** 1-3 loops for state indicators

## Creating State Animations

### Using the TUI

1. **Start Creation:**
   ```bash
   tangent create
   ```

2. **Create Base Character:**
   - This is your idle/immutable foundation
   - All states build from this base

3. **Add Agent States:**
   - Choose state name (plan, think, execute, etc.)
   - Create 3+ frames per state
   - Option to start from base or create from scratch

4. **Preview:**
   - Use "Animate all states" to see each state
   - Verify animations look smooth and clear

5. **Export:**
   - Export JSON and README
   - Submit as GitHub PR

### Frame Design Tips

**1. Consistency:**
- Keep character recognizable across states
- Maintain size and proportions
- Use consistent visual language

**2. Clarity:**
- Each state should be distinctly different
- Viewers should understand state at a glance
- Avoid overly subtle differences

**3. Smoothness:**
- Create natural transitions between frames
- Avoid jarring changes
- Test at target FPS (5 recommended)

**4. Readability:**
- Terminal characters must be legible
- Consider different terminal themes
- Test on light and dark backgrounds

## API Usage

### Load and Use States

```go
import "github.com/wildreason/tangent/pkg/characters"

// Load agent
agent, _ := characters.LibraryAgent("mercury")

// Use required states (guaranteed to exist)
agent.Plan(os.Stdout)    // Planning state
agent.Think(os.Stdout)   // Thinking state
agent.Execute(os.Stdout) // Executing state

// Use optional states (check first)
states := agent.ListStates()
if contains(states, "wait") {
    agent.Wait(os.Stdout)
}
```

### Custom Animation

```go
// Animate with custom FPS and loops
agent.AnimateState(os.Stdout, "plan", 8, 3) // 8 FPS, 3 loops

// Show specific state once
agent.ShowState(os.Stdout, "think")
```

### Introspection

```go
// Get character details
char := agent.GetCharacter()

// List all available states
states := agent.ListStates()
fmt.Printf("Available states: %v\n", states)

// Access state properties
if state, exists := char.States["plan"]; exists {
    fmt.Printf("Plan state has %d frames at %d FPS\n", 
        len(state.Frames), state.AnimationFPS)
}
```

## State Design Examples

### Minimalist (2x2)

**Base:**
```
██
██
```

**Plan (3 frames):**
```
Frame 1:  Frame 2:  Frame 3:
▀█        ██        ▐█
██        ▀█        ██
```

### Detailed (11x5)

**Base:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
    ███    
   █   █   
   █   █   
```

**Think (3 frames):**
```
Frame 1:      Frame 2:      Frame 3:
  ▐▜███▜▌      ▐▜███▜▌      ▐▜███▜▌  
 ▀▜█████▛▀    ▀▜█████▛▀    ▀▜█████▛▀ 
  . ███  .      ███          ███    
   █   █       █   █      . █   █ .
   █   █       █   █        █   █   
```

## Validation Requirements

When creating or contributing characters:

- ✓ Base character required
- ✓ Minimum 3 states (plan, think, execute)
- ✓ Minimum 3 frames per state
- ✓ Valid Unicode block patterns
- ✓ Consistent dimensions across all frames
- ✓ FPS between 1-30
- ✓ State names are valid identifiers

The `tangent admin validate` command checks all these requirements before registration.

## Best Practices

1. **Start Simple:** Begin with 3 frames per state, expand if needed
2. **Test Early:** Preview animations frequently during creation
3. **Be Consistent:** Maintain visual identity across states
4. **Think Context:** Consider how states will be used (CLI agent status)
5. **Follow Patterns:** Study existing characters for inspiration

## See Also

- [API Reference](API.md) - Using states in your code
- [Pattern Guide](PATTERNS.md) - Unicode block patterns
- [Contributing Characters](CONTRIBUTING_CHARACTERS.md) - Submission workflow

---

**Tangent** - Terminal Agent Designer

