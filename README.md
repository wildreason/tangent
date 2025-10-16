# Tangent - Terminal Avatars for AI Agents

Give your AI agent a face. Expressive, state-based avatars that bring personality to AI agents in terminal applications.

Status: Alpha.6 | Go 1.21+

---

## Why AI Agents Need Avatars

When users interact with AI agents, they need more than text logs. They need to *see* the agent working - planning, thinking, executing. Tangent provides terminal-native avatars with semantic states that map directly to AI agent workflows.

**The Vision**: By v1.0, Tangent aims to be the single source for AI agent terminal avatars - the standard library every AI-native CLI application uses.

---

## Quick Start

### Browse Available Avatars

#### Install CLI (copy paste to your terminal)
```bash
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash
```

#### Window Shop
```bash
# Browse avatars
tangent browse              # List all available avatars
tangent browse fire         # Preview fire avatar with states
```

### Integrate with Your AI Agent

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("fire")

// Show agent states during workflow
agent.Plan(os.Stdout)       // Agent is planning
agent.Think(os.Stdout)      // Agent is thinking
agent.Success(os.Stdout)    // Agent was successful
agent.Searching(os.Stdout)  // Agent is searching files
```

One line to load, and methods to give your agent presence.

---

## Available Avatar

**fire** (11x3) - 14 states including:
- Core: plan, think, execute, wait, error, success
- Extended: bash, read, write, search, build, communicate, block, blocked

Additional avatars planned for Beta release.

---

## Core Concepts

### State-Based Personality (Lifecycles)

AGENT LIFECYCLE STATES:
  - waiting      → Agent idle, no task assigned
  - thinking     → Agent analyzing/reasoning (LLM inference)
  - planning     → Agent breaking down task into steps
  - reading      → Agent examining files/docs
  - searching    → Agent searching codebase (Grep/Glob)
  - writing      → Agent creating/editing files
  - executing    → Agent running bash commands
  - building     → Agent compiling/testing
  - communicating→ Agent messaging other agents (MCP)
  - blocked      → Agent waiting on dependency
  - error        → Agent encountered failure
  - success      → Agent completed task

### Terminal-First Design

- Built for terminal applications, not web
- Unicode-based, no external dependencies
- Lightweight, zero-config integration
- Works with any Go CLI framework

### AI-Native Philosophy

Tangent is purpose-built for AI agents. Every design decision optimizes for:
- Giving agents visible presence
- Mapping states to AI workflows
- Building trust through personality
- Terminal-native simplicity

---

## Example: AI Agent with Avatar

```go
package main

import (
    "fmt"
    "os"
    "time"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    agent, _ := characters.LibraryAgent("fire")

    fmt.Println("Starting AI analysis...")
    agent.Plan(os.Stdout)
    time.Sleep(2 * time.Second)

    fmt.Println("\nProcessing data...")
    agent.Think(os.Stdout)
    time.Sleep(2 * time.Second)

    fmt.Println("\nExecuting task...")
    agent.Execute(os.Stdout)

    fmt.Println("\nTask complete!")
}
```

---

## Documentation

- **[API Reference](docs/API.md)** - Complete API documentation
- **[State Guide](docs/STATES.md)** - Understanding agent states
- **[Pattern Reference](docs/PATTERNS.md)** - Unicode pattern system

---

## Roadmap

**Alpha** (current): API-first positioning, curated avatar library
**Beta**: Expanded avatar library (8 avatars), richer state vocabulary
**v1.0**: The standard for AI agent terminal avatars

---

## For Advanced Users

Want to create custom avatars? Tangent includes a creation tool for contributors. See [docs/CREATORS.md](docs/CREATORS.md) for the full workflow.

Character contributions welcome via GitHub PR. We curate all avatars to maintain quality and AI-agent focus.

---

## Installation

```bash
# Via installer (recommended)
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash

# Via Go
go install github.com/wildreason/tangent/cmd/tangent@latest

# As Go package
go get github.com/wildreason/tangent/pkg/characters
```

---

## License

MIT © 2025 Wildreason, Inc

**Tangent** - Terminal Avatars for AI Agents
