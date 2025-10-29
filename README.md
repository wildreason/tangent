# Tangent

Terminal avatars for AI agents. Go library.

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

## Usage

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sa")
agent.Plan(os.Stdout)
agent.Think(os.Stdout)
agent.Success(os.Stdout)
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
agent.Success(writer)

// Custom states
agent.ShowState(writer, "arise")
agent.ShowState(writer, "approval")

// Introspection
states := agent.ListStates()
hasState := agent.HasState("think")
```

## Avatars

7 characters × 17 states

- **sa** - Red (#FF0000)
- **ri** - Orange (#FF8800)
- **ga** - Gold (#FFD700)
- **ma** - Green (#00FF00)
- **pa** - Blue (#0088FF)
- **dha** - Purple (#8800FF)
- **ni** - Pink (#FF0088)

## States

arise, wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked, resting, approval

## License

MIT © 2025 Wildreason, Inc
