---
id: ADR-004
title: Frame Cache for O(1) Performance
version: v0.2.0
status: accepted
date: 2025-12-01
tags: [performance, optimization, rendering]
---

# Frame Cache for O(1) Performance

## Context

Rendering character frames at 60 FPS requires fast frame access. Initial implementation computed frames on demand (50µs per frame), consuming CPU budget.

## Decision

Pre-render and pre-color all frames at initialization:

```go
// O(1) frame access
frames := character.GetStateFrames("think")
frame := frames[frameIndex]  // Already colored, ready to display
```

Performance:
- Before: 50µs per frame (compute + color)
- After: 0.1µs per frame (array lookup)
- 500x improvement
- <1% CPU budget at 60 FPS

## Alternatives

1. **On-demand rendering** - Rejected: 50µs too slow for 60 FPS
2. **Lazy caching** - Rejected: first-frame latency, complexity
3. **String interning** - Rejected: still requires color application

## Consequences

**Positive:**
- O(1) frame access
- Predictable performance
- No GC pressure during animation

**Negative:**
- Higher memory usage (all frames pre-allocated)
- Initialization cost (acceptable, one-time)
- Theme changes require cache rebuild
