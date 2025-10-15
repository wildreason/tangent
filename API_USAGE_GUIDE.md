# Tangent Alpha.4 - API Usage Guide

## Overview

Tangent Alpha.4 provides a **simple, state-based API** for using agent characters in your CLI applications. Characters are created using the interactive TUI builder and compiled into the library for instant API access.

## Quick Start

### 1. Import the Package

```go
import "github.com/wildreason/tangent/pkg/characters"
```

### 2. Load an Agent

```go
agent, err := characters.LibraryAgent("mercury")
if err != nil {
    log.Fatal(err)
}
```

### 3. Use Agent States

```go
// Show planning state
agent.Plan(os.Stdout)

// Show thinking state  
agent.Think(os.Stdout)

// Show executing state
agent.Execute(os.Stdout)
```

## Complete API Reference

### Core Functions

#### `ListLibrary() []string`
Lists all available agent characters in the library.

```go
agents := characters.ListLibrary()
fmt.Printf("Available: %v\n", agents)
// Output: Available: [demo4 mercury water water5]
```

#### `LibraryAgent(name string) (*AgentCharacter, error)`
Loads an agent character from the library.

```go
agent, err := characters.LibraryAgent("mercury")
if err != nil {
    return fmt.Errorf("failed to load agent: %w", err)
}
```

### AgentCharacter Methods

#### Required States (Guaranteed)

Every registered agent has these three states:

```go
// Planning state - Agent analyzing and strategizing
agent.Plan(os.Stdout)

// Thinking state - Agent processing and reasoning
agent.Think(os.Stdout)

// Executing state - Agent taking action
agent.Execute(os.Stdout)
```

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
    agent, _ := characters.LibraryAgent("mercury")
    
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
    agent, _ := characters.LibraryAgent("water5")
    
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
    // Load multiple agents
    agents := []Agent{
        {name: "Planner", character: mustLoadAgent("mercury")},
        {name: "Analyzer", character: mustLoadAgent("water")},
        {name: "Executor", character: mustLoadAgent("water5")},
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
    agent, err := loadAgentSafely("mercury")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("✓ Agent loaded and validated")
    agent.Plan(os.Stdout)
}
```

## API Design Principles

### ✓ Simplicity First
- **One function to load**: `LibraryAgent(name)`
- **Three core methods**: `Plan()`, `Think()`, `Execute()`
- **Zero configuration**: Characters work out of the box

### ✓ Type Safety
- Compile-time errors for missing agents
- All agents guaranteed to have required states
- Clear error messages for debugging

### ✓ Consistent Interface
- All agents implement the same `AgentCharacter` interface
- Predictable method signatures
- Standard `io.Writer` for flexible output

### ✓ Discoverability
- `ListLibrary()` shows what's available
- `ListStates()` shows what each agent can do
- Clear method names match agent behaviors

## Character Requirements

All registered agents **must have**:
- ✓ Base character (idle state)
- ✓ Plan state (3+ frames)
- ✓ Think state (3+ frames)
- ✓ Execute state (3+ frames)

This guarantees API consumers can always use the core methods.

## Testing Your Integration

```bash
# 1. List available agents
./tangent browse

# 2. Test an agent interactively
./tangent browse mercury

# 3. Test specific state
./tangent browse mercury --state plan --fps 8 --loops 2

# 4. Run your application
go run your-app.go
```

## Workflow Summary

```
Designer's Workflow:
1. tangent create              → Interactive TUI builder
2. Create base + 3 states      → plan, think, execute
3. Export for contribution     → name.json + name-README.md
4. tangent admin register      → Compiles into library
5. make build                  → Binary ready

Developer's Workflow:
1. Import package              → import "github.com/wildreason/tangent/pkg/characters"
2. List agents                 → characters.ListLibrary()
3. Load agent                  → characters.LibraryAgent("name")
4. Use states                  → agent.Plan(), agent.Think(), agent.Execute()
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

**Alpha.4 (new):**
```go
// One line to load, instant use
agent, _ := characters.LibraryAgent("mercury")
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
err := agent.Wait(os.Stdout)
// Error: state "wait" not found for character mercury
```

**Solution**: Use `agent.ListStates()` to see available states, or stick to required states (plan, think, execute)

### Character not animating correctly
- Ensure agent was registered with `tangent admin register`
- Rebuild with `make build` after registration
- Verify states have 3+ frames: `tangent browse <name>`

## Support

- **Documentation**: `/Users/btsznh/wild/characters/README.md`
- **Examples**: `/Users/btsznh/wild/characters/examples/`
- **Repository**: https://github.com/wildreason/tangent

---

**Tangent v0.1.0-alpha.4** - Terminal Agent Designer

