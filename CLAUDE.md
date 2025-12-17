# CLAUDE.md

This file provides guidance to Claude Code when working with this repository.

## Project Overview

**characters** (aka Tangent) is a Go library providing animated ASCII avatars for AI agents. Used by Jump TUI and other WildReason tools.

Core capabilities:
- 7 character set (sam, rio, ga, ma, pa, da, ni) with unique visual identities
- 16 animation states (think, read, write, bash, resting, etc.)
- Thread-safe TangentClient API for TUI integration
- Pre-rendered frame cache for O(1) performance
- Micro avatar format (8x2) for compact displays

## Build & Run

```bash
make generate        # Regenerate character files from constants.go
make verify-generate # CI check: ensure generated files are committed
make test            # Run tests
make build-cli       # Build tangent-cli (internal tool)
make release         # Create tag + push + GitHub release
```

## Architecture Decisions

Key decisions in `decisions/`:
- Code generation (v0.2.0): `decisions/001-codegen-single-source.md`
- State registry (v0.2.0): `decisions/002-state-registry.md`
- TangentClient API (v0.3.0): `decisions/003-tangent-client-api.md`
- Frame cache (v0.2.0): `decisions/004-frame-cache.md`
- Micro avatar format (v0.3.0): `decisions/005-micro-avatar-format.md`

- Wall Street rush (v0.4.0): `decisions/006-dynamic-noise-wallstreet.md`

Agent queries:
```bash
grep -r "Alternatives" decisions/     # Find tradeoffs
grep -r "tags:.*performance" decisions/  # Find performance decisions
```

## Key Patterns

**Code Generation:**
- Edit `pkg/characters/library/constants.go` (single source of truth)
- Run `make generate` to regenerate character files
- Never edit sam.go, rio.go, etc. directly

**Adding a State:**
1. Create JSON in `pkg/characters/stateregistry/states/`
2. Run `make generate`
3. All characters inherit the new state

**TUI Integration:**
```go
client := characters.NewTangentClient(characters.Sam)
client.SetState("think")
frame := client.GetFrame()  // Thread-safe, auto-ticking
```

## Directory Structure

```
pkg/characters/
  library/           # Generated character files + constants.go
  stateregistry/     # State definitions (JSON)
  tangentclient.go   # Main API
docs/                # API reference, ecosystem guide
dev/                 # Internal development guides
decisions/           # Architecture decision records
```

## Common Tasks

**Add a character:**
1. Add to `constants.go` (name, symbol, colors)
2. Create frame data in library/
3. Run `make generate`

**Change theme:**
```go
characters.SetTheme("garden")  // latte, bright, garden, cozy
```

**Get micro avatar:**
```go
frames := character.GetMicroFrames("think")
```
