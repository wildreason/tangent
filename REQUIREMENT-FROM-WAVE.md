# Micro Avatar Requirement

**To:** Sam@Wave
**From:** Rio@Tangent
**Date:** 2024-12-09
**Subject:** RE: 11x2 "Face-Only" Avatar Variants for Compact Mode

---

## Implementation Complete

Sam,

Micro avatars are done and merged to main. Here's what we shipped:

### API

```go
// Get micro avatar (10x2)
agent, _ := characters.LibraryAgentMicro("sam")

// Use states - same as standard avatars
agent.ShowBase(writer)
agent.ShowState(writer, "resting")
agent.AnimateState(writer, "read", 5, 2)

// List available micro characters
names := characters.ListMicroLibrary()
// Returns: [da-micro, ga-micro, ma-micro, ni-micro, pa-micro, rio-micro, sam-micro]
```

### Specs

| Attribute | Value |
|-----------|-------|
| Width | 10 chars |
| Height | 2 lines |
| Characters | All 7 (sam, rio, ga, ma, pa, da, ni) |
| States | All 7 (resting, arise, wait, read, write, search, approval) |
| Frames | 2-5 per state |

Note: Final width is 10 chars, not 11. Your JSON designs used 10-char patterns and you confirmed that works for Wave's layout.

### CLI Testing

```bash
./tangent-cli browse sam --micro
./tangent-cli browse sam --micro --state resting
./tangent-cli browse rio --micro --state arise --fps 3 --loops 5
```

### Answers to Your Questions

1. **Can expressions remain distinct at 2-line height?** Yes - the patterns use quadrant blocks (corners, halves) which give good expression variety even at 2 lines.

2. **Standardize on specific ASCII chars?** We use the existing Tangent pattern language (F, L, R, T, B, 1-8, etc.) which compiles to Unicode block characters. All agents share the same micro patterns, just different colors.

3. **Timeline?** Done. Same day.

### Files Added

```
pkg/characters/microstateregistry/
  types.go           - MicroDefinition, MicroState, MicroFrame
  loader.go          - Embedded JSON loader
  states/micro.json  - Your 7-state definitions

pkg/characters/library/
  micro_characters.go - Registers all 7 micro variants

pkg/characters/
  characters.go      - Added LibraryAgentMicro(), ListMicroLibrary()
```

Let me know if Wave needs any adjustments.

- Rio

---

**To:** Rio@Tangent
**From:** Sam@Wave
**Date:** 2024-12-08
**Subject:** 11x2 "Face-Only" Avatar Variants for Compact Mode

---

## Context

Wave is adding a compact overlay mode (`--mini`) that fits in **3 terminal lines**. The current 11x4 avatars don't fit this constraint. We need a cropped "micro" variant.

## Current State

```
Tangent provides 11x4 avatars:

   .---.      <- line 1
  ( o o )     <- line 2 (eyes)
  |  >  |     <- line 3 (mouth)
   '---'      <- line 4
```

## Requirement

**New variant: 11 wide x 2 tall** - just the face (eyes + mouth expression).

```
Micro avatar (11x2):

 ( o o )      <- line 1: eyes
 |  >  |      <- line 2: mouth/expression
```

## Layout in Wave Compact Mode

```
 ◕ ◕   sam │ reading src/components/Button.tsx
  ‿        │ analyzing props and dependencies...
───────────────────────────────────────────────────
```

| Column | Width | Content |
|--------|-------|---------|
| Avatar | 11 chars | Micro face (2 lines) |
| Name | variable | Agent name (colored) |
| Task | remaining | 2-line status text |

## Animation Requirements

The micro avatars should support the same **7 states** as full avatars:

| State | Expression Idea |
|-------|-----------------|
| `resting` | Neutral, maybe subtle blink |
| `arise` | Eyes widen, alert |
| `wait` | Relaxed, patient |
| `read` | Focused, scanning |
| `write` | Concentrated, active |
| `search` | Looking around, curious |
| `approval` | Raised eyebrow, questioning |

Animation frames can be simpler than full avatars - even 2-3 frames per state is fine for this size.

## Technical Constraints

- **Width:** Exactly 11 characters (matches current)
- **Height:** Exactly 2 lines (hard constraint)
- **Characters:** ASCII printable (same as current)
- **States:** Same 7 states as full avatars
- **Frames:** Minimum 2 frames per state for animation

## Deliverable

New character set or variant mode in Tangent:

```go
// Option A: Separate micro character files
characters.LibraryAgent("sam-micro")

// Option B: Size parameter on existing agents
characters.LibraryAgent("sam", characters.SizeMicro)
```

Either approach works for Wave. Let us know which fits Tangent's architecture better.

## Questions for Tangent

1. Can expressions remain distinct at 2-line height?
2. Should we standardize on specific ASCII chars for eyes/mouth across all agents?
3. Timeline estimate for this variant set?

---

**Priority:** Medium-High (blocks Wave compact mode feature)

Let me know if you need clarification or want to discuss design options.

- Sam
