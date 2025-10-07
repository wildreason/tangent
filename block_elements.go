package palette

import (
	"errors"
	"fmt"
	"strings"
)

// Block Elements (U+2580–U+259F)
// Only block elements are defined here. No geometric shapes.
const (
	// Solid and half blocks
	FullBlock      rune = '█' // U+2588
	UpperHalfBlock rune = '▀' // U+2580
	LowerHalfBlock rune = '▄' // U+2584
	LeftHalfBlock  rune = '▌' // U+258C
	RightHalfBlock rune = '▐' // U+2590

	// Shading blocks
	LightShade  rune = '░' // U+2591
	MediumShade rune = '▒' // U+2592
	DarkShade   rune = '▓' // U+2593

	// Quadrants
	QuadrantUpperLeft  rune = '▘' // U+2598
	QuadrantUpperRight rune = '▝' // U+259D
	QuadrantLowerLeft  rune = '▖' // U+2596
	QuadrantLowerRight rune = '▗' // U+2597

	// Diagonals
	QuadrantUL_LR rune = '▚' // U+259A upper-left + lower-right
	QuadrantUR_LL rune = '▞' // U+259E upper-right + lower-left

	// Three-quadrant composites
	QuadrantUL_UR_LL rune = '▛' // U+259B
	QuadrantUL_UR_LR rune = '▜' // U+259C
	QuadrantUL_LL_LR rune = '▙' // U+2599
	QuadrantUR_LL_LR rune = '▟' // U+259F
)

// Short, easy-to-type variable names
var (
	// Solid and half blocks
	fb = FullBlock      // █
	uh = UpperHalfBlock // ▀
	lh = LowerHalfBlock // ▄
	lf = LeftHalfBlock  // ▌
	rf = RightHalfBlock // ▐

	// Shading blocks
	ls = LightShade  // ░
	ms = MediumShade // ▒
	ds = DarkShade   // ▓

	// Quadrants
	ul = QuadrantUpperLeft  // ▘
	ur = QuadrantUpperRight // ▝
	ll = QuadrantLowerLeft  // ▖
	lr = QuadrantLowerRight // ▗

	// Diagonals
	diag1 = QuadrantUL_LR // ▚
	diag2 = QuadrantUR_LL // ▞

	// Three-quadrant composites
	comp1 = QuadrantUL_UR_LL // ▛
	comp2 = QuadrantUL_UR_LR // ▜
	comp3 = QuadrantUL_LL_LR // ▙
	comp4 = QuadrantUR_LL_LR // ▟
)

// ByName exposes a stable name → rune mapping for easy lookup.
var ByName = map[string]rune{
	"FULL_BLOCK":       FullBlock,
	"UPPER_HALF_BLOCK": UpperHalfBlock,
	"LOWER_HALF_BLOCK": LowerHalfBlock,
	"LEFT_HALF_BLOCK":  LeftHalfBlock,
	"RIGHT_HALF_BLOCK": RightHalfBlock,

	"LIGHT_SHADE":  LightShade,
	"MEDIUM_SHADE": MediumShade,
	"DARK_SHADE":   DarkShade,

	"QUADRANT_UPPER_LEFT":  QuadrantUpperLeft,
	"QUADRANT_UPPER_RIGHT": QuadrantUpperRight,
	"QUADRANT_LOWER_LEFT":  QuadrantLowerLeft,
	"QUADRANT_LOWER_RIGHT": QuadrantLowerRight,

	"QUADRANT_UL_LR": QuadrantUL_LR,
	"QUADRANT_UR_LL": QuadrantUR_LL,

	"QUADRANT_UL_UR_LL": QuadrantUL_UR_LL,
	"QUADRANT_UL_UR_LR": QuadrantUL_UR_LR,
	"QUADRANT_UL_LL_LR": QuadrantUL_LL_LR,
	"QUADRANT_UR_LL_LR": QuadrantUR_LL_LR,
}

// Allowed returns the default allow-list of runes for sprite authoring.
// Typically you will restrict to FullBlock and space for strict pixel art,
// but the complete set is provided for advanced shading.
func Allowed() map[rune]bool {
	m := map[rune]bool{
		' ': true, // space is always allowed
	}
	for _, r := range ByName {
		m[r] = true
	}
	return m
}

// B returns a string made of n full blocks (█).
func B(n int) string { return strings.Repeat(string(FullBlock), n) }

// S returns a string made of n spaces.
func S(n int) string { return strings.Repeat(" ", n) }

// R builds a row by concatenating parts; handy for explicit row composition.
func R(parts ...string) string { return strings.Join(parts, "") }

// ValidateFrame checks width, height, and allowed runes.
func ValidateFrame(frame []string, width, height int, allowed map[rune]bool) error {
	if len(frame) != height {
		return fmt.Errorf("frame height=%d expected=%d", len(frame), height)
	}
	for i, line := range frame {
		if l := len([]rune(line)); l != width {
			return fmt.Errorf("line %d width=%d expected=%d", i, l, width)
		}
		for _, r := range line {
			if !allowed[r] {
				return fmt.Errorf("line %d contains disallowed rune: %q", i, r)
			}
		}
	}
	return nil
}

// MustValidate is a helper that panics on validation error (useful in init()).
func MustValidate(frame []string, width, height int, allowed map[rune]bool) {
	if err := ValidateFrame(frame, width, height, allowed); err != nil {
		panic(err)
	}
}

// EnsureSpaceOnly is a guard for strict two-glyph sprites (█ and space).
func EnsureSpaceOnly(allowed map[rune]bool) map[rune]bool {
	if allowed == nil {
		allowed = map[rune]bool{}
	}
	// reset to only full block and space
	for k := range allowed {
		delete(allowed, k)
	}
	allowed[' '] = true
	allowed[FullBlock] = true
	return allowed
}

// Err helper for callers that prefer an exported error variable.
var ErrInvalidFrame = errors.New("invalid sprite frame")
