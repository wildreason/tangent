---
id: ADR-003
title: TangentClient API Design
version: v0.3.0
status: accepted
date: 2025-12-11
tags: [api, animation, framework-agnostic]
---

# TangentClient API Design

## Context

TUI frameworks (Bubble Tea, tview) needed 300+ lines of adapter code each for character animation. Thread safety, state management, and tick timing were duplicated across integrations.

## Decision

Centralized `TangentClient` animation controller:

```go
client := characters.NewTangentClient(sam)
client.SetState("think")
frame := client.GetFrame()  // Thread-safe, auto-ticking
```

Features:
- Thread-safe state management
- Per-state FPS (fast for edits, calm for idle)
- State aliases (deprecated states map to current)
- Auto-tick option
- State queuing

## Alternatives

1. **Per-framework adapters** - Rejected: 300+ lines each, maintenance burden
2. **Raw character API** - Rejected: unsafe, manual tick management
3. **Event-based API** - Rejected: complexity for simple use case

## Consequences

**Positive:**
- 98% reduction in integration code (310 → 5 lines)
- Consistent behavior across frameworks
- Thread-safe by default

**Negative:**
- Opinionated API (not customizable tick)
- Global state per client instance
- Must use client, not raw character
