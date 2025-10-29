# Tangent Development TODO

**Project:** Terminal Character Library
**Current Version:** v1.0.0
**Next Release:** v1.1.0

---

## High Priority - v1.1.0 Release

### Character Library Expansion

**Goal:** Create 7 character variants for murmur's agent system

**Status:** 🔴 Not Started

**Requirements:**
- [ ] All characters 11x3 dimensions
- [ ] All characters support 14 states (wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked)
- [ ] Color variants of fire base design (same animation patterns, different colors)
- [ ] Each character must loop smoothly (first/last frame compatibility)

---

### Characters to Create

#### 1. fire (EXISTING - v1.0.0)
**Status:** ✅ Complete
- Color: Orange/red (#FF4500, #FF6347, #FFD700)
- Theme: Flames
- Use case: Default avatar, `sa` agent

---

#### 2. mercury
**Status:** 🔴 TODO
- Color: Silver/white (#C0C0C0, #E8E8E8, #FFFFFF)
- Theme: Liquid metal, flowing
- Use case: `ri` agent
- Pattern codes: F (bright white), 6 (light silver), 4 (medium silver), 2 (dark silver)

**Implementation:**
```
File: pkg/characters/library/mercury.go
Copy: fire.go structure
Replace: Color codes (R→F, L→6, 5→4, 1→2)
Test: Visual render of all 14 states
```

**Example frame:**
```
__F6FFF6F__
_46FFFFF64_
___66_66___
```

---

#### 3. neptune
**Status:** 🔴 TODO
- Color: Cyan/blue (#00CED1, #1E90FF, #87CEEB)
- Theme: Ocean waves, water
- Use case: `ga` agent
- Pattern codes: L (blue), 6 (light cyan), 4 (medium cyan), 2 (dark blue)

**Implementation:**
```
File: pkg/characters/library/neptune.go
Copy: fire.go structure
Replace: Color codes (R→L, maintain L, adjust shades for cyan)
Test: Visual render of all 14 states
```

**Example frame:**
```
__L6LLL6L__
_46LLLLL64_
___66_66___
```

---

#### 4. mars
**Status:** 🔴 TODO
- Color: Red/crimson (#DC143C, #8B0000, #FF0000)
- Theme: War, intensity, energy
- Use case: `ma` agent
- Pattern codes: R (red), 6 (light red), 4 (medium red), 1 (dark red)

**Implementation:**
```
File: pkg/characters/library/mars.go
Copy: fire.go structure
Replace: Emphasize R codes, darker reds
Test: Visual render of all 14 states
```

**Example frame:**
```
__R6RRR6R__
_16RRRRR61_
___11_11___
```

---

#### 5. jupiter
**Status:** 🔴 TODO
- Color: Gold/yellow (#FFD700, #FFA500, #FFFF00)
- Theme: Storm clouds, power, majesty
- Use case: `pa` agent
- Pattern codes: F (bright gold), 6 (light yellow), 5 (medium gold), 3 (dark gold)

**Implementation:**
```
File: pkg/characters/library/jupiter.go
Copy: fire.go structure
Replace: Color codes for gold theme
Test: Visual render of all 14 states
```

**Example frame:**
```
__F6FFF6F__
_56FFFFF65_
___55_55___
```

---

#### 6. saturn
**Status:** 🔴 TODO
- Color: Purple/violet (#9370DB, #8A2BE2, #9400D3)
- Theme: Rings, organization, systems
- Use case: `da` agent
- Pattern codes: New code P or reuse R→purple, 6 (light purple), 4 (medium purple), 2 (dark purple)

**Implementation:**
```
File: pkg/characters/library/saturn.go
Copy: fire.go structure
Replace: Color codes for purple theme (may need new pattern code)
Test: Visual render of all 14 states
```

**Example frame:**
```
__P6PPP6P__
_46PPPPP64_
___66_66___
```

**Note:** May require extending pattern code system to support purple. Alternative: reuse existing codes with color palette mapping.

---

#### 7. uranus
**Status:** 🔴 TODO
- Color: Teal/aqua (#20B2AA, #48D1CC, #00CED1)
- Theme: Ice crystals, cold, methodical
- Use case: `ni` agent
- Pattern codes: L (teal), 6 (light aqua), 4 (medium teal), 2 (dark teal)

**Implementation:**
```
File: pkg/characters/library/uranus.go
Copy: fire.go structure
Replace: Color codes for teal theme
Test: Visual render of all 14 states
```

**Example frame:**
```
__L6LLL6L__
_46LLLLL64_
___66_66___
```

---

## Implementation Checklist

### Per Character

For each character (mercury, neptune, mars, jupiter, saturn, uranus):

**Setup:**
- [ ] Create `pkg/characters/library/<name>.go`
- [ ] Copy fire.go structure as template
- [ ] Update metadata (name, description, author)
- [ ] Define color palette for pattern codes

**State Implementation (14 states):**
- [ ] `wait` - Gentle pulse/breathing (3 frames)
- [ ] `think` - Focused intensity (3 frames)
- [ ] `plan` - Organized movement (3 frames)
- [ ] `execute` - Dynamic action (3 frames)
- [ ] `error` - Alert/warning indication (2 frames)
- [ ] `success` - Celebration/confirmation (3 frames)
- [ ] `read` - Scanning motion (3 frames)
- [ ] `search` - Sweeping/seeking (3 frames)
- [ ] `write` - Creating/building (3 frames)
- [ ] `bash` - Executing/processing (3 frames)
- [ ] `build` - Constructing (3 frames)
- [ ] `communicate` - Transmitting/sending (3 frames)
- [ ] `block` - Waiting/synchronizing (2 frames)
- [ ] `blocked` - Paused/holding (2 frames)

**Testing:**
- [ ] Visual render test (print all states)
- [ ] Animation loop test (smooth transitions)
- [ ] Dimension verification (11x3 all frames)
- [ ] Pattern code validation (only valid codes)

**Registration:**
- [ ] Add to library registry (`library.go`)
- [ ] Update README.md character catalog
- [ ] Add to CHANGELOG.md

---

## Timeline Estimate

**Total:** 6 characters × ~2 hours each = ~12 hours work

**Per Character Breakdown:**
- Setup & template: 15 minutes
- Implement 14 states: 60 minutes (copy/modify from fire)
- Testing & validation: 30 minutes
- Documentation: 15 minutes

**Schedule:**
- Week 1: mercury, neptune, mars (3 characters)
- Week 2: jupiter, saturn, uranus (3 characters)
- Week 3: Integration testing, documentation, release

---

## Testing Strategy

### Visual Render Tests

Create test script to render all characters:

```bash
# test_characters.sh
for char in fire mercury neptune mars jupiter saturn uranus; do
    echo "=== $char ==="
    go run test/render.go $char wait
    echo ""
done
```

### Animation Tests

```bash
# test_animation.sh
go run test/animate.go mercury wait
# Should show smooth 3-frame loop
```

### Integration Tests

```bash
# Test handwave integration
cd /path/to/handwave
go mod edit -replace github.com/wildreason/tangent=/path/to/tangent
go test ./internal/characters/...
```

---

## Release Checklist - v1.1.0

**Pre-release:**
- [ ] All 7 characters implemented
- [ ] All characters tested visually
- [ ] Animation loops verified
- [ ] README.md updated with character catalog
- [ ] CONTRACT.md reviewed (ensure compliance)
- [ ] CHANGELOG.md updated

**Release:**
- [ ] Version bump in code
- [ ] Git tag: `git tag v1.1.0`
- [ ] Push tag: `git push origin v1.1.0`
- [ ] GitHub release notes

**Post-release:**
- [ ] Notify handwave team
- [ ] Update handwave dependency: `go get github.com/wildreason/tangent@v1.1.0`
- [ ] Verify handwave integration works
- [ ] Update murmur documentation with avatar mappings

---

## Future Enhancements (v1.2.0+)

**Character Variants:**
- [ ] Large size (20x5 dimensions)
- [ ] Small size (7x2 dimensions)
- [ ] Alternate styles (geometric, pixel-art, minimalist)

**Color System:**
- [ ] Extend pattern codes (P=purple, G=green, Y=yellow)
- [ ] Custom color palettes (user-defined)
- [ ] Dynamic color shifting

**Animation Features:**
- [ ] Variable frame rates per state
- [ ] Transition frames between states
- [ ] Particle effects

**Tooling:**
- [ ] Character editor/designer tool
- [ ] Animation preview tool
- [ ] Character validator script

---

## Notes

**Design Philosophy:**
- Keep it simple: Start with color variants of fire
- Maintain consistency: All 11x3, all 14 states
- Quality over quantity: Better to have 7 solid characters than 20 mediocre ones

**Pattern Code Strategy:**
For v1.1.0, reuse existing codes (R, L, F, 1-6) with character-specific color palettes. This avoids extending the pattern system prematurely.

**Character Naming:**
Stick to planetary theme for consistency and memorability. Future expansions could explore other themes (elements, seasons, myths).

---

**Last Updated:** 2025-10-24
**Owner:** Tangent Team
**Review:** Weekly progress check
