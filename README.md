# Tangent - Terminal Avatars for AI Agents

Give your AI agent a face. Expressive, state-based avatars that bring personality to AI agents in terminal applications.

Status: Alpha.5 | Go 1.21+

---

## Why AI Agents Need Avatars

When users interact with AI agents, they need more than text logs. They need to *see* the agent working - planning, thinking, executing. Tangent provides terminal-native avatars with semantic states that map directly to AI agent workflows.

**The Vision**: By v1.0, Tangent aims to be the single source for AI agent terminal avatars - the standard library every AI-native CLI application uses.

---

## Quick Start

### Browse Available Avatars

```bash
# Install CLI
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash

# Browse avatars
tangent browse              # List all available avatars
tangent browse mercury      # Preview mercury avatar with states
```

### Integrate with Your AI Agent

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("mercury")

// Show agent states during workflow
agent.Plan(os.Stdout)       // Agent is planning
agent.Think(os.Stdout)      // Agent is thinking
agent.Execute(os.Stdout)    // Agent is executing
```

That's it. One line to load, three methods to give your agent presence.

---

## Available Avatars

Each avatar has a distinct personality designed for AI agent workflows:

- **mercury** - Fast, analytical avatar
- **water** - Flowing, adaptive avatar
- **water5** - Refined aquatic personality
- **demo4** - Classic demonstration avatar

More avatars coming in Beta. See `tangent browse` for the latest.

---

## Core Concepts

### State-Based Personality

Avatars aren't decorations - they're semantic representations of what your AI agent is doing:

- **plan** - Agent analyzing options and strategizing
- **think** - Agent processing information and reasoning
- **execute** - Agent taking action and performing tasks

Optional states: **wait**, **error**, **success**

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
    agent, _ := characters.LibraryAgent("mercury")

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

## Philosophy

**Terminal-first.** Not a web component library ported to terminals. Built for terminal applications from the ground up.

**AI-native.** Not generic animations. Purpose-built for AI agents that need visible presence during planning, reasoning, and execution.

**Single-source vision.** By v1.0, Tangent aims to be the standard - the library every AI CLI developer reaches for when their agent needs a face.

---

## Roadmap

**Alpha.5** (current): API-first positioning, curated avatar library
**Beta**: Expanded avatar library (10-20 avatars), richer state vocabulary
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
