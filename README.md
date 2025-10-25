# Tangent - Terminal Avatars for AI Agents

Give your AI agent a face. 
Tangent provides expressive, state-based avatars that bring personality to AI agents in terminal environments

Status: Alpha.6 | Go 1.21+

---

## Why AI Agents Need Avatars

AI agents do more than print logs - they plan, think, execute and wait. Tangent gives that process a visible rhythm through avatars that reflect agent state in real time.

**The Vision**:
By v1, Tangent aims to be the standard library for terminal avatars - the shared visual language for every AI-native CLI and development agent.

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
tangent browse fire --state success --fps 8 --loops 4 # Preview each live states
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

## Available Avatars

All avatars are 11x3 dimensions with 14 states each:

| Name | Theme | Color Palette | Agent Mapping | Description |
|------|-------|---------------|---------------|-------------|
| **fire** | Flames | Orange/red | `sa` (default) | High-energy, active agents |
| **mercury** | Liquid metal | Silver/white | `ri` | Fast, fluid agents |
| **neptune** | Ocean waves | Cyan/blue | `ga` | Deep-thinking, analytical agents |
| **mars** | War energy | Red/crimson | `ma` | Aggressive, action-oriented agents |
| **jupiter** | Storm power | Gold/yellow | `pa` | Powerful, commanding agents |
| **saturn** | Orbital rings | Purple/violet | `da` | Organized, systematic agents |
| **uranus** | Ice crystals | Teal/aqua | `ni` | Cool, methodical agents |

**States:** wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked

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

### Build By AI Agents (Understory v0.1)

Tangent was built entirely by AI coding agents through Wildreason's Understory autonomous coding platform. It's not a design experiment - its' proof that autonomous agents can design, build, and ship usabel tools for developers.

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
