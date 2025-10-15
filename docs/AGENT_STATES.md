# Agent States

Tangent characters support agent states, allowing AI agents to express their current behavioral state through visual representation.

## Overview

Agent states represent what an AI agent is doing at any given moment:

- **plan** - Agent is analyzing and planning
- **think** - Agent is processing information
- **execute** - Agent is performing actions
- **wait** - Agent is waiting for input
- **error** - Agent is handling errors
- **success** - Agent is celebrating success

## Standard States

All characters must have a minimum of 3 standard states:

### Required States

#### plan
**Purpose:** Show the agent analyzing a problem and formulating a plan

**Visual Guidelines:**
- Use question marks (?) for inquiry
- Use dots (.) for analysis
- Show contemplative posture

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   ?  ?    
```

#### think
**Purpose:** Show the agent processing information and reasoning

**Visual Guidelines:**
- Use dots (...) for processing
- Use brain/thought symbols
- Show concentration

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   . . .   
```

#### execute
**Purpose:** Show the agent performing actions

**Visual Guidelines:**
- Use arrows (>) for movement
- Use action symbols
- Show dynamic motion

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   > > >   
```

### Optional Standard States

#### wait
**Purpose:** Show the agent in idle/waiting state

**Visual Guidelines:**
- Use calm, neutral pose
- Minimal animation
- Patient appearance

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   ▘▘  ▝▝  
```

#### error
**Purpose:** Show the agent encountering an issue

**Visual Guidelines:**
- Use X marks for errors
- Use confused symbols
- Show distress or confusion

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   X  X    
```

#### success
**Purpose:** Show the agent completing a task successfully

**Visual Guidelines:**
- Use checkmarks (✓) for success
- Use celebration symbols
- Show happiness or achievement

**Example:**
```
  ▐▜███▜▌  
 ▀▜█████▛▀ 
   ✓  ✓    
```

## Custom States

Characters can have custom states for unique behaviors. Custom states should:

- Have descriptive names
- Serve a specific purpose
- Be visually distinct from standard states
- Be documented in character description

**Examples of custom states:**
- `celebrate` - Special celebration animation
- `sleep` - Resting/inactive state
- `alert` - High-priority attention state
- `confused` - Uncertainty state

## Using Agent States

### Basic Usage

```go
import "github.com/wildreason/tangent/pkg/characters"

// Load character with agent states
agent, _ := characters.LibraryAgent("robot")

// Use standard states
agent.Plan(os.Stdout)      // Planning state
agent.Think(os.Stdout)     // Thinking state
agent.Execute(os.Stdout)   // Executing state
agent.Success(os.Stdout)   // Success state
```

### Custom State Access

```go
// Access any state by name
agent.ShowState(os.Stdout, "custom-state")

// Check if state exists
if agent.HasState("celebrate") {
    agent.ShowState(os.Stdout, "celebrate")
}

// List all available states
states := agent.ListStates()
fmt.Println("Available states:", states)
```

### State Inspection

```go
// Get state description
desc, _ := agent.GetStateDescription("plan")
fmt.Println(desc) // "Agent analyzing and planning"

// Check character info
name := agent.Name()
personality := agent.Personality()
```

## Agent State Workflow

Typical AI agent workflow using states:

```go
agent, _ := characters.LibraryAgent("robot")

// 1. Receive task
agent.Wait(os.Stdout)

// 2. Analyze task
agent.Plan(os.Stdout)

// 3. Process solution
agent.Think(os.Stdout)

// 4. Execute solution
agent.Execute(os.Stdout)

// 5. Handle result
if success {
    agent.Success(os.Stdout)
} else {
    agent.Error(os.Stdout)
    // Retry logic...
}
```

## Creating Characters with Agent States

### Using the CLI

1. Run `tangent` and create a new character
2. Choose personality (efficient, friendly, analytical, creative)
3. Add agent states when prompted:
   - Start with required states: plan, think, execute
   - Add optional states: wait, error, success
   - Add custom states as needed

### Programmatic Creation

```go
import "github.com/wildreason/tangent/pkg/characters/domain"

character := &domain.Character{
    Name:        "my-agent",
    Personality: "efficient",
    Width:       9,
    Height:      4,
    States: map[string]domain.State{
        "plan": {
            Name:        "Planning",
            Description: "Agent analyzing task",
            StateType:   "standard",
            Frames: []domain.Frame{
                {
                    Name: "plan_frame",
                    Lines: []string{
                        "_L5FFF5R_",
                        "_6FFFFF6_",
                        "__?F_F?__",
                        "__FF_FF__",
                    },
                },
            },
        },
        // ... more states
    },
}
```

## State Design Guidelines

### Visual Consistency

- Maintain character shape across states
- Use consistent visual language
- Keep personality consistent
- Ensure states are distinguishable

### State Transitions

Consider how states flow together:
- `wait` → `plan` → `think` → `execute` → `success`
- `execute` → `error` → `think` → `execute` (retry)
- `success` → `wait` (ready for next task)

### Animation Within States

States can have multiple frames for animation:

```go
"execute": {
    Frames: []domain.Frame{
        {Lines: []string{"...", "...", "..."}},  // Frame 1
        {Lines: []string{"...", "...", "..."}},  // Frame 2
        {Lines: []string{"...", "...", "..."}},  // Frame 3
    },
}
```

## Best Practices

### For Character Designers

1. **Start with required states** - Ensure plan, think, execute are well-designed
2. **Test state transitions** - Preview how states flow together
3. **Match personality** - Design should reflect character personality
4. **Be creative** - Make states visually interesting and expressive
5. **Document custom states** - Explain purpose of custom states

### For Developers

1. **Use appropriate states** - Match agent behavior to state
2. **Handle missing states** - Check state existence before use
3. **Provide feedback** - Show state changes to users
4. **Consider timing** - Allow time for state visibility
5. **Error handling** - Gracefully handle state errors

## Examples

See `examples/agent_states.go` for a complete demonstration of:
- Creating characters with states
- Using standard states
- Custom state access
- State inspection
- Practical agent workflows

## Contributing Characters

When contributing characters to the library:

1. Include minimum 3 required states (plan, think, execute)
2. Design states that match character personality
3. Test all states in the CLI
4. Document any custom states
5. Follow the contribution guidelines in `.github/CONTRIBUTING_CHARACTERS.md`

## API Reference

### AgentCharacter Methods

- `Plan(writer io.Writer) error` - Show planning state
- `Think(writer io.Writer) error` - Show thinking state
- `Execute(writer io.Writer) error` - Show executing state
- `Wait(writer io.Writer) error` - Show waiting state
- `Error(writer io.Writer) error` - Show error state
- `Success(writer io.Writer) error` - Show success state
- `ShowState(writer io.Writer, stateName string) error` - Show any state by name
- `ListStates() []string` - Get all available state names
- `HasState(stateName string) bool` - Check if state exists
- `GetStateDescription(stateName string) (string, error)` - Get state description
- `Name() string` - Get character name
- `Personality() string` - Get character personality

## Further Reading

- [Character Library](LIBRARY.md) - Available characters
- [Pattern Guide](PATTERN_GUIDE.md) - Pattern character reference
- [Contributing Characters](../.github/CONTRIBUTING_CHARACTERS.md) - Contribution guidelines



