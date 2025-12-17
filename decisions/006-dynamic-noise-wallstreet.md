---
id: ADR-006
title: Dynamic Noise for Extreme Motion (Wall Street Rush)
version: v0.4.0
status: accepted
date: 2025-12-17
tags: [animation, ux, design, micro-avatars, perception]
---

# Dynamic Noise for Extreme Motion (Wall Street Rush)

## Context

Characters were too calming. Smooth animations didn't reflect agent intensity - an agent reading 50 files or running complex bash commands looked identical to one idling. Agent runtime spans 30s to 2min on normal operations; users couldn't perceive urgency or progress.

Problem: Animation speed ≠ perceived activity. Even fast FPS felt serene. Multiple states (16) didn't create differentiation - they all felt the same.

## Decision

Two key changes:

**1. State Reduction**
- 16 states → 7 states → 3 dominant states
- Current: `read`, `resting`, `approval` handle all operations
- Simplification created clearer visual language

**2. Dynamic Noise (Wall Street Rush)**
- Random character flicker mimicking high-intensity activity
- Protected frame slots (character identity preserved)
- Shifting gradient with logarithmic perception exploit

Result: 0% → 70% of vision achieved in <1 day through rapid iteration.

## The Emergent Perceptual Effect

The breakthrough was unplanned. Linear brightness stepping:
```
0.25 → 0.40 → 0.55 → 0.70 → 0.85 → 1.00 → 1.15 → 1.30
 +15%   +15%   +15%   +15%   +15%   +15%   +15%
```

But human vision perceives brightness **logarithmically**:
- Dark → less dark (0.25 → 0.40): **+60% relative change** - very noticeable
- Bright → brighter (1.15 → 1.30): **+13% relative change** - subtle

Effect:
- Dark-to-mid transitions feel "fast" (big visual jumps)
- Mid-to-bright transitions feel "slow" (subtle changes)
- Wrap-around from 1.30 → 0.25 creates sharp visual "reset"

**Why it works for Wall Street rush:**
- Fast chaotic movement on the dark side (urgency)
- Smooth glide on the bright side (resolution)
- Sharp reset (new cycle begins)

This mimics how stock tickers feel: rapid bursts of activity followed by brief moments of clarity.

## Alternatives Explored

1. **More states** - Rejected: 16 states didn't differentiate, felt same
2. **Glitch effects** - Rejected: felt broken, not energetic
3. **Color cycling** - Rejected: distracting, lost character identity
4. **Fibonacci randomness** - Rejected: too predictable, mathematical feel
5. **Faster FPS** - Rejected: smooth ≠ intense, just faster calm

**What worked:** Fewer states + chaos within structure + logarithmic perception.

## Consequences

**Positive:**
- Users perceive agent activity intensity
- 3 states simpler than 16 (clearer visual language)
- Emergent perceptual effect creates natural rhythm
- Simple implementation with profound visual impact

**Negative:**
- 70% achieved, 30% refinement remaining
- Less state granularity (by design)

## Future Direction (→ 90%)

Introduce variations within `read` state to support nuanced operations:
- `read` base → file reading
- `read` + write hints → editing
- `read` + search hints → grep/glob
- `read` + bash hints → command execution

Same 3-state foundation, richer expression through variations.

## Design Journey

```
16 states, smooth animation → too calm, no differentiation
7 states, faster FPS → still calm
3 states (read/resting/approval) → clearer but still missing urgency
+ Random flicker + protected slots → promising
+ Logarithmic brightness perception → Wall Street Rush ✓ (70%)
+ Read state variations → future (90%)
```

Key insight: Less is more. 3 states with dynamic noise > 16 states with smooth animation.
