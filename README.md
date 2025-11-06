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
agent.Execute(os.Stdout)
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

// Custom states
agent.ShowState(writer, "arise")
agent.ShowState(writer, "approval")

// Introspection
states := agent.ListStates()
hasState := agent.HasState("think")

// Themes - set color theme for all agents
characters.SetTheme("latte")        // Switch to latte theme
themes := characters.ListThemes()   // ["bright", "cozy", "garden", "latte"]
current := characters.GetCurrentTheme()  // "latte"
```

## Avatars

7 characters × 16 states × 4 themes

## States

arise, wait, think, plan, execute, error, read, search, write, bash, build, communicate, block, blocked, resting, approval

## License

MIT © 2025 Wildreason, Inc
