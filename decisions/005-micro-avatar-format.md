---
id: ADR-005
title: Micro Avatar Format Specification
version: v0.3.0
status: accepted
date: 2025-12-11
tags: [format, cross-product, wave-integration]
---

# Micro Avatar Format Specification

## Context

Wave v1 compact overlay mode needed smaller avatars than standard 11x5 format. Cross-team coordination required defining a compact variant.

## Decision

8x2 micro avatar format:

- Width: 8 characters (down from 11)
- Height: 2 lines (down from 5)
- Padding removed for density
- Same state support as full avatars

Evolution: 11x2 → 10x2 → 8x2 (iterative reduction based on Wave feedback)

API:
```go
micro := character.GetMicroFrames("think")
```

## Alternatives

1. **Single size only** - Rejected: Wave overlay too cramped
2. **Arbitrary scaling** - Rejected: ASCII art doesn't scale well
3. **Icon-based** - Rejected: loses character personality

## Consequences

**Positive:**
- Fits Wave compact overlay
- Maintains character recognition
- Same animation states as full

**Negative:**
- Two frame sets to maintain per state
- Less detail than full avatars
- Cross-product dependency on format spec
