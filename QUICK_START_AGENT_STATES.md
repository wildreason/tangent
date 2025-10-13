# Quick Start: Agent States

## For Users (Developers)

### Install
```bash
go get github.com/wildreason/tangent/pkg/characters
```

### Use Agent States
```go
import "github.com/wildreason/tangent/pkg/characters"

// Load character
agent, _ := characters.LibraryAgent("rocket")

// Use states
agent.Plan(os.Stdout)      // Planning
agent.Think(os.Stdout)     // Thinking
agent.Execute(os.Stdout)   // Executing
agent.Success(os.Stdout)   // Success
```

### Complete Workflow
```go
// AI Agent workflow
agent, _ := characters.LibraryAgent("robot")

agent.Wait(os.Stdout)      // Waiting for task
agent.Plan(os.Stdout)      // Analyzing task
agent.Think(os.Stdout)     // Processing solution
agent.Execute(os.Stdout)   // Performing action
agent.Success(os.Stdout)   // Task complete!
```

### Inspect States
```go
// List all states
states := agent.ListStates()
// Output: [error, execute, plan, success, think, wait]

// Check if state exists
if agent.HasState("custom") {
    agent.ShowState(os.Stdout, "custom")
}

// Get state description
desc, _ := agent.GetStateDescription("plan")
```

## For Contributors (Character Designers)

### Create Character
```bash
# Run interactive CLI
tangent

# Choose: Create new character
# Enter name, dimensions, personality
# Add states: plan, think, execute (minimum)
# Add optional states: wait, error, success
# Add custom states as needed
```

### Export for Contribution
```bash
# In character builder menu:
# Choose: 6. Export for contribution (JSON)

# This creates:
# - your-character.json
# - your-character-README.md
```

### Submit to Library
```bash
# 1. Fork repository
git clone https://github.com/YOUR_USERNAME/tangent

# 2. Create branch
git checkout -b add-your-character

# 3. Add JSON file
cp your-character.json tangent/characters/

# 4. Commit and push
git add characters/your-character.json
git commit -m "Add your-character"
git push origin add-your-character

# 5. Create Pull Request on GitHub
```

## Agent States Reference

### Required States (Minimum 3)
- **plan** - Agent analyzing and planning
- **think** - Agent processing information
- **execute** - Agent performing actions

### Optional States
- **wait** - Agent waiting for input
- **error** - Agent handling errors
- **success** - Agent celebrating success

### Custom States
Add any custom states for unique behaviors!

## Pattern Codes

```
Basic Blocks:
  F = █  T = ▀  B = ▄  L = ▌  R = ▐

Quadrants:
  1 = ▘  2 = ▝  3 = ▖  4 = ▗
  5 = ▛  6 = ▜  7 = ▙  8 = ▟

Shading:
  . = ░  : = ▒  # = ▓

Special:
  _ = Space  X = Mirror
```

## Examples

### Example 1: Simple Agent
```go
agent, _ := characters.LibraryAgent("robot")
agent.Plan(os.Stdout)
agent.Execute(os.Stdout)
agent.Success(os.Stdout)
```

### Example 2: Error Handling
```go
agent, _ := characters.LibraryAgent("robot")
agent.Execute(os.Stdout)

if err != nil {
    agent.Error(os.Stdout)
    agent.Think(os.Stdout)
    agent.Execute(os.Stdout)  // Retry
}

agent.Success(os.Stdout)
```

### Example 3: Custom States
```go
agent, _ := characters.LibraryAgent("robot")

// Use standard states
agent.Plan(os.Stdout)

// Use custom state
agent.ShowState(os.Stdout, "celebrate")
```

## Documentation

- **Full Guide**: [docs/AGENT_STATES.md](docs/AGENT_STATES.md)
- **Contributing**: [.github/CONTRIBUTING_CHARACTERS.md](.github/CONTRIBUTING_CHARACTERS.md)
- **Examples**: [examples/agent_states.go](examples/agent_states.go)
- **API Reference**: [docs/AGENT_STATES.md#api-reference](docs/AGENT_STATES.md#api-reference)

## Help

- **CLI Help**: `tangent help`
- **Gallery**: `tangent gallery`
- **Issues**: https://github.com/wildreason/tangent/issues
- **Discussions**: https://github.com/wildreason/tangent/discussions

