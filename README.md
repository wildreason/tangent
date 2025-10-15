# Tangent — Terminal Agent Designer

Design and use state-animated Unicode characters for CLI agents.

Status: Alpha • Go 1.21+

---

## Quick Start

### For Creators (TUI)
```bash
# 1) Install
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash

# 2) Create → Export → PR
tangent create          
# Build character with live preview
# In TUI: Export JSON + README
# Submit GitHub PR with the exported JSON
```

### For Developers (API)
```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("mercury")
agent.Plan(os.Stdout)
agent.Think(os.Stdout)
agent.Execute(os.Stdout)
```

---

## What Works (Alpha)
- Bubbletea TUI with split-pane live preview
- Agent states: plan, think, execute (+ optional wait/error/success)
- Minimum 3 frames per state; override fps/loops when browsing
- Simple CLI: create, browse, view (admin hidden; demo removed)
- Library workflow: export JSON → PR → admin register → compiled in
- Ready-to-use API: `LibraryAgent()` with state methods

---

## Core Concepts
- Base frame (idle) + state animations
- Required states: plan, think, execute
- 3+ frames per state for animation

---

## Minimal Examples

### Developers
```go
agent, _ := characters.LibraryAgent("mercury")
agent.Plan(os.Stdout)
```

### Creators
```bash
tangent create
tangent browse              # list agents
tangent browse mercury      # animate agent
```

---

## Documentation
- `docs/API.md` – API reference
- `docs/STATES.md` – Agent states guide
- `docs/PATTERNS.md` – Unicode pattern reference
- `docs/CLI.md` – CLI commands and options

---

## Contributing
Character PRs welcome (JSON export from TUI)

## License
MIT © 2025 Wildreason, Inc
