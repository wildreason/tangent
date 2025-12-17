---
id: ADR-002
title: Global State Registry Pattern
version: v0.2.0
status: accepted
date: 2025-12-01
tags: [states, architecture, animation]
---

# Global State Registry Pattern

## Context

Characters need animation states (think, read, write, bash, etc.). Early design had character-specific state definitions, causing duplication and inconsistent state support.

## Decision

Global state definitions in `pkg/characters/stateregistry/states/*.json` via `go:embed`:

- States are universal, not character-specific
- Adding a state → all 7 characters inherit it automatically
- 16 states: arise, wait, think, plan, execute, error, read, search, write, bash, build, communicate, block, blocked, resting, approval

State JSON defines:
- Frame sequences
- FPS per state
- Transition rules

## Alternatives

1. **Per-character states** - Rejected: duplication, inconsistent support
2. **Hardcoded in Go** - Rejected: less flexible, harder to preview
3. **External state files** - Rejected: deployment complexity

## Consequences

**Positive:**
- Single place to add/modify states
- All characters stay in sync
- JSON enables tooling (preview, validation)

**Negative:**
- Characters can't have unique states (by design)
- State changes affect all characters at once
- Must understand registry pattern to add states
