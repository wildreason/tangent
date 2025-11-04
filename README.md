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

## Complete Example

Here's a complete working example demonstrating an AI agent workflow with state transitions:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	// Load the "sa" (Red) agent from the library
	agent, err := characters.LibraryAgent("sa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🤖 AI Agent Workflow Demo\n")

	// Agent starts planning
	fmt.Println("📋 Planning task...")
	agent.Plan(os.Stdout)
	time.Sleep(500 * time.Millisecond)

	// Agent processes information
	fmt.Println("\n💭 Thinking...")
	agent.Think(os.Stdout)
	time.Sleep(500 * time.Millisecond)

	// Agent executes the task
	fmt.Println("\n⚡ Executing task...")
	agent.Execute(os.Stdout)
	time.Sleep(500 * time.Millisecond)

	// Agent waits for results
	fmt.Println("\n⏳ Waiting for completion...")
	agent.Wait(os.Stdout)
	time.Sleep(500 * time.Millisecond)

	// Agent completes successfully
	fmt.Println("\n✅ Success!")
	agent.Success(os.Stdout)

	fmt.Println("\n\n🎉 Agent workflow complete!")
}
```

This example demonstrates a typical AI agent lifecycle:
1. **Plan** - Agent analyzes and plans the task
2. **Think** - Agent processes and contemplates the approach
3. **Execute** - Agent actively performs the task
4. **Wait** - Agent waits for results or feedback
5. **Success** - Agent celebrates successful completion

You can also handle errors in your workflow:

```go
// If something goes wrong
if err != nil {
	fmt.Println("\n❌ Error encountered!")
	agent.Error(os.Stdout)
	return
}
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
