# Tangent API Reference

## Overview

Tangent provides **terminal avatars for AI agents** with a simple, state-based API. Give your AI agent a face with expressive avatars that show what your agent is doing - planning, thinking, executing.

### Why Terminal Avatars?

AI agents need visible presence. When your agent is working, users should see it - not just read logs. Tangent's state-based avatars map directly to AI agent workflows, making your agent more trustworthy and engaging.

### Discovery Workflow

Before integrating, browse available avatars:

```bash
# Install CLI to browse avatars
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash

# Discover avatars
tangent browse              # See all available
tangent browse fire         # Preview specific avatar
tangent browse fire --state think --fps 8
```

Pick the avatar that matches your agent's personality, then integrate with three simple API calls.

## Quick Start

### 1. Import the Package

```go
import "github.com/wildreason/tangent/pkg/characters"
```

### 2. Load an Avatar

```go
// Load an avatar by name
agent, err := characters.LibraryAgent("fire")
if err != nil {
    log.Fatal(err)
}
```

### 3. Show Agent States

```go
// During AI agent workflow, show what the agent is doing

agent.Plan(os.Stdout)       // Agent is planning/strategizing
agent.Think(os.Stdout)      // Agent is processing/reasoning
agent.Execute(os.Stdout)    // Agent is acting/performing
```

That's it. Your AI agent now has a visible presence during its workflow.

## Complete API Reference

### Core Functions

#### `ListLibrary() []string`
Lists all available avatars in the library.

```go
avatars := characters.ListLibrary()
fmt.Printf("Available avatars: %v\n", avatars)
// Output: Available avatars: [fire]
```

**Pro tip**: Use `tangent browse` in terminal for visual preview before integrating.

#### `LibraryAgent(name string) (*AgentCharacter, error)`
Loads an avatar from the library.

```go
agent, err := characters.LibraryAgent("fire")
if err != nil {
    return fmt.Errorf("failed to load agent: %w", err)
}
```

### AgentCharacter Methods

#### Required States (Guaranteed)

Every avatar has these three core states that map to AI agent workflows:

```go
// Planning state - AI agent analyzing options and strategizing
agent.Plan(os.Stdout)

// Thinking state - AI agent processing information and reasoning
agent.Think(os.Stdout)

// Executing state - AI agent taking action and performing tasks
agent.Execute(os.Stdout)
```

These states are **guaranteed** - your code will never break because a state is missing.

#### Optional States

Some agents may include these additional states:

```go
// Waiting state - Agent idle/pending
agent.Wait(os.Stdout)

// Error state - Agent encountered error
agent.Error(os.Stdout)

// Success state - Agent completed successfully
agent.Success(os.Stdout)
```

#### Utility Methods

```go
// Show base (idle) character
agent.ShowBase(os.Stdout)

// List all available states for this agent
states := agent.ListStates()
fmt.Printf("States: %v\n", states)

// Get underlying character data
char := agent.GetCharacter()
fmt.Printf("Name: %s, Size: %dx%d\n", char.Name, char.Width, char.Height)
```

#### Advanced: Custom State Animation

```go
// Animate specific state with custom settings
err := agent.ShowState(os.Stdout, "plan")
err := agent.AnimateState(os.Stdout, "think", 10, 3) // 10 FPS, 3 loops
```

## Practical Examples

### Example 1: Simple CLI Agent Indicator

```go
package main

import (
    "fmt"
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    agent, _ := characters.LibraryAgent("fire")
    
    fmt.Println("Starting analysis...")
    agent.Plan(os.Stdout)
    
    fmt.Println("\nProcessing data...")
    agent.Think(os.Stdout)
    
    fmt.Println("\nExecuting task...")
    agent.Execute(os.Stdout)
    
    fmt.Println("\n✓ Complete!")
}
```

### Example 2: AI Agent Status Display

```go
package main

import (
    "fmt"
    "os"
    "time"
    "github.com/wildreason/tangent/pkg/characters"
)

type AgentStatus int

const (
    StatusPlanning AgentStatus = iota
    StatusThinking
    StatusExecuting
)

func showAgentStatus(agent *characters.AgentCharacter, status AgentStatus) {
    switch status {
    case StatusPlanning:
        fmt.Print("\r[PLANNING] ")
        agent.Plan(os.Stdout)
    case StatusThinking:
        fmt.Print("\r[THINKING] ")
        agent.Think(os.Stdout)
    case StatusExecuting:
        fmt.Print("\r[EXECUTING] ")
        agent.Execute(os.Stdout)
    }
}

func main() {
    agent, _ := characters.LibraryAgent("fire")
    
    // Simulate agent workflow
    showAgentStatus(agent, StatusPlanning)
    time.Sleep(2 * time.Second)
    
    showAgentStatus(agent, StatusThinking)
    time.Sleep(2 * time.Second)
    
    showAgentStatus(agent, StatusExecuting)
    fmt.Println("\n✓ Task completed!")
}
```

### Example 3: Multi-Agent System

```go
package main

import (
    "fmt"
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

type Agent struct {
    name      string
    character *characters.AgentCharacter
}

func main() {
    // Load multiple agents (all using fire with different roles)
    agents := []Agent{
        {name: "Planner", character: mustLoadAgent("fire")},
        {name: "Analyzer", character: mustLoadAgent("fire")},
        {name: "Executor", character: mustLoadAgent("fire")},
    }
    
    // Show agent states
    for _, agent := range agents {
        fmt.Printf("\n%s:\n", agent.name)
        agent.character.ShowBase(os.Stdout)
    }
    
    // Planner thinks
    fmt.Println("\n→ Planner is analyzing...")
    agents[0].character.Think(os.Stdout)
    
    // Analyzer processes
    fmt.Println("\n→ Analyzer is processing...")
    agents[1].character.Think(os.Stdout)
    
    // Executor acts
    fmt.Println("\n→ Executor is running...")
    agents[2].character.Execute(os.Stdout)
}

func mustLoadAgent(name string) *characters.AgentCharacter {
    agent, err := characters.LibraryAgent(name)
    if err != nil {
        panic(err)
    }
    return agent
}
```

### Example 4: Error Handling & Validation

```go
package main

import (
    "fmt"
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func loadAgentSafely(name string) (*characters.AgentCharacter, error) {
    // Check if agent exists
    available := characters.ListLibrary()
    found := false
    for _, a := range available {
        if a == name {
            found = true
            break
        }
    }
    
    if !found {
        return nil, fmt.Errorf("agent '%s' not found. Available: %v", name, available)
    }
    
    // Load agent
    agent, err := characters.LibraryAgent(name)
    if err != nil {
        return nil, fmt.Errorf("failed to load agent: %w", err)
    }
    
    // Verify required states
    char := agent.GetCharacter()
    requiredStates := []string{"plan", "think", "execute"}
    for _, state := range requiredStates {
        if _, exists := char.States[state]; !exists {
            return nil, fmt.Errorf("agent missing required state: %s", state)
        }
    }
    
    return agent, nil
}

func main() {
    agent, err := loadAgentSafely("fire")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("✓ Agent loaded and validated")
    agent.Plan(os.Stdout)
}
```

## API Design Principles

### Simplicity First
- **One function to load**: `LibraryAgent(name)`
- **Three core methods**: `Plan()`, `Think()`, `Execute()`
- **Zero configuration**: Characters work out of the box

### Type Safety
- Compile-time errors for missing agents
- All agents guaranteed to have required states
- Clear error messages for debugging

### Consistent Interface
- All agents implement the same `AgentCharacter` interface
- Predictable method signatures
- Standard `io.Writer` for flexible output

### Discoverability
- `ListLibrary()` shows what's available
- `ListStates()` shows what each agent can do
- Clear method names match agent behaviors

## Character Requirements

All registered agents **must have**:
- Base character (idle state)
- Plan state (3+ frames)
- Think state (3+ frames)
- Execute state (3+ frames)

This guarantees API consumers can always use the core methods.

## Testing Your Integration

```bash
# 1. List available agents
./tangent browse

# 2. Test an agent interactively
./tangent browse fire

# 3. Test specific state
./tangent browse fire --state plan --fps 8 --loops 2

# 4. Run your application
go run your-app.go
```

## Migration from Alpha.3

If you used Alpha.3, the API has **simplified significantly**:

**Alpha.3 (old):**
```go
// Complex setup required
repo := infrastructure.NewFileCharacterRepository("./characters")
engine := infrastructure.NewAnimationEngine()
service := service.NewCharacterService(compiler, repo, engine)
char, _ := service.LoadCharacter("rocket")
```

**Alpha.4+ (new):**
```go
// One line to load, instant use
agent, _ := characters.LibraryAgent("fire")
agent.Plan(os.Stdout)
```

## Troubleshooting

### Agent not found
```go
agent, err := characters.LibraryAgent("myagent")
// Error: agent 'myagent' not found in library
```

**Solution**: Check available agents with `tangent browse` or `characters.ListLibrary()`

### State missing
```go
err := agent.CustomState(os.Stdout)
// Error: state "custom_state" not found for character fire
```

**Solution**: Use `agent.ListStates()` to see available states, or stick to required states (plan, think, execute)

### Character not animating correctly
- Ensure agent was registered with `tangent admin register`
- Rebuild with `make build` after registration
- Verify states have 3+ frames: `tangent browse <name>`

## Support

- **Documentation**: `docs/`
- **Examples**: `examples/`
- **Repository**: https://github.com/wildreason/tangent

---

**Tangent** - Terminal Avatars for AI Agents

