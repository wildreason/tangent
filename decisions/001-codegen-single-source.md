---
id: ADR-001
title: Code Generation as Single Source of Truth
version: v0.2.0
status: accepted
date: 2025-12-01
tags: [codegen, architecture, characters]
---

# Code Generation as Single Source of Truth

## Context

Character definitions (sam.go, rio.go, etc.) contain repetitive boilerplate: name, symbol, colors, frame data. Manual maintenance led to inconsistencies and tedious updates when renaming or adding characters.

## Decision

Use `constants.go` as single source of truth. Code generator (`generator_codegen.go`) auto-generates all character files.

```
constants.go (edit this)
    ↓ make generate
sam.go, rio.go, etc. (generated, don't edit)
```

- Character files are NOT manually editable
- Renaming = change 1 line in constants.go + regenerate
- CI enforces via `make verify-generate`

## Alternatives

1. **Manual character files** - Rejected: error-prone, tedious updates
2. **YAML/JSON config** - Rejected: loses Go type safety
3. **Runtime generation** - Rejected: startup cost, no static analysis

## Consequences

**Positive:**
- Single source of truth eliminates drift
- Trivial to add/rename characters
- CI catches uncommitted generation

**Negative:**
- Must run `make generate` after constants.go changes
- Generated files in git (intentional for go get compatibility)
- Learning curve for contributors
