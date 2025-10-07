package characters

import (
	"fmt"
	"strings"
)

// Block Elements - Centralized character definitions
// All block element constants used throughout the package
const (
	FullBlock       = '█' // U+2588 Full Block
	TopHalfBlock    = '▀' // U+2580 Top Half
	BottomHalfBlock = '▄' // U+2584 Bottom Half
	LeftHalfBlock   = '▌' // U+258C Left Half
	RightHalfBlock  = '▐' // U+2590 Right Half

	LightShade  = '░' // U+2591 Light Shade
	MediumShade = '▒' // U+2592 Medium Shade
	DarkShade   = '▓' // U+2593 Dark Shade

	QuadUpperLeft  = '▘' // U+2598 Quadrant Upper Left
	QuadUpperRight = '▝' // U+259D Quadrant Upper Right
	QuadLowerLeft  = '▖' // U+2596 Quadrant Lower Left
	QuadLowerRight = '▗' // U+2597 Quadrant Lower Right

	ThreeQuad5 = '▛' // U+259B Three Quadrants: UL+UR+LL (reverse of 4)
	ThreeQuad6 = '▜' // U+259C Three Quadrants: UL+UR+LR (reverse of 3)
	ThreeQuad7 = '▙' // U+2599 Three Quadrants: UL+LL+LR (reverse of 2)
	ThreeQuad8 = '▟' // U+259F Three Quadrants: UR+LL+LR (reverse of 1)

	DiagonalBackward = '▚' // U+259A Diagonal Backward (\)
	DiagonalForward  = '▞' // U+259E Diagonal Forward (/)
)

// PatternCompiler converts single-character hex-style codes to Unicode block elements
// Pattern format: "__R5FFF6L__" (no commas, like hex color codes)
type PatternCompiler struct {
	// Pattern mapping: single char -> Unicode block element
	patterns map[rune]rune
}

// NewPatternCompiler creates a new pattern compiler with hex-style mappings
// Variant B: Compact single-character codes
func NewPatternCompiler() *PatternCompiler {
	return &PatternCompiler{
		patterns: map[rune]rune{
			// Basic blocks
			'F': FullBlock,
			'T': TopHalfBlock,
			'B': BottomHalfBlock,
			'L': LeftHalfBlock,
			'R': RightHalfBlock,

			// Shading blocks
			'.': LightShade,
			':': MediumShade,
			'#': DarkShade,

			// Single quadrants (1-4)
			'1': QuadUpperLeft,
			'2': QuadUpperRight,
			'3': QuadLowerLeft,
			'4': QuadLowerRight,

			// Three-quadrant composites (5-8, reverse of 1-4)
			'5': ThreeQuad5, // ▛ (reverse of 4): UL+UR+LL
			'6': ThreeQuad6, // ▜ (reverse of 3): UL+UR+LR
			'7': ThreeQuad7, // ▙ (reverse of 2): UL+LL+LR
			'8': ThreeQuad8, // ▟ (reverse of 1): UR+LL+LR

			// Diagonals
			'\\': DiagonalBackward,
			'/':  DiagonalForward,

			// Space
			'_': ' ',
		},
	}
}

// CompilePattern converts a compact hex-style pattern to a Unicode string
// Example: "__R6FFF6L__" → "  ▐▜███▜▌  "
func (pc *PatternCompiler) CompilePattern(pattern string) string {
	var result strings.Builder

	// Process each character in the pattern
	for _, char := range pattern {
		// Check if it's a pattern we know
		if unicode, exists := pc.patterns[char]; exists {
			result.WriteRune(unicode)
		} else {
			// Unknown pattern - treat as space
			result.WriteRune(' ')
		}
	}

	return result.String()
}

// ValidatePattern checks if a pattern string is valid
func (pc *PatternCompiler) ValidatePattern(pattern string) error {
	for i, char := range pattern {
		if _, exists := pc.patterns[char]; !exists {
			return fmt.Errorf("invalid pattern character '%c' at position %d", char, i)
		}
	}
	return nil
}

// CompileFrame compiles a complete frame from a slice of patterns
func (pc *PatternCompiler) CompileFrame(patterns []string) string {
	var lines []string
	for _, pattern := range patterns {
		line := pc.CompilePattern(pattern)
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

// CompileCharacter compiles a complete character from frame patterns
func (pc *PatternCompiler) CompileCharacter(name string, width, height int, framePatterns [][]string) (*Character, error) {
	builder := NewBuilder(name, width, height)

	for i, framePattern := range framePatterns {
		// Compile the frame
		frame := pc.CompileFrame(framePattern)

		// Split into lines and add as patterns
		lines := strings.Split(frame, "\n")
		for _, line := range lines {
			if line != "" {
				builder.Pattern(line)
			}
		}

		// Start new frame if there are more frames (but not after the last one)
		if i < len(framePatterns)-1 {
			builder.NewFrame()
		}
	}

	return builder.Build()
}

// AddPattern adds a custom pattern mapping
func (pc *PatternCompiler) AddPattern(key rune, unicode rune) {
	pc.patterns[key] = unicode
}

// GetPatterns returns all available patterns
func (pc *PatternCompiler) GetPatterns() map[rune]rune {
	return pc.patterns
}

// GetPatternLegend returns a formatted legend of all patterns
func (pc *PatternCompiler) GetPatternLegend() string {
	var legend strings.Builder
	legend.WriteString("Pattern Legend (Hex-Style Codes):\n")
	legend.WriteString("==================================\n\n")
	legend.WriteString("Basic Blocks:\n")
	legend.WriteString("  F = █ (Full block)\n")
	legend.WriteString("  T = ▀ (Top half)\n")
	legend.WriteString("  B = ▄ (Bottom half)\n")
	legend.WriteString("  L = ▌ (Left half)\n")
	legend.WriteString("  R = ▐ (Right half)\n\n")
	legend.WriteString("Shading:\n")
	legend.WriteString("  . = ░ (Light shade)\n")
	legend.WriteString("  : = ▒ (Medium shade)\n")
	legend.WriteString("  # = ▓ (Dark shade)\n\n")
	legend.WriteString("Quadrants (Single, 1-4):\n")
	legend.WriteString("  1 = ▘ (Upper Left)\n")
	legend.WriteString("  2 = ▝ (Upper Right)\n")
	legend.WriteString("  3 = ▖ (Lower Left)\n")
	legend.WriteString("  4 = ▗ (Lower Right)\n\n")
	legend.WriteString("Quadrants (Three, 5-8 = reverse of 1-4):\n")
	legend.WriteString("  5 = ▛ (reverse of 4)\n")
	legend.WriteString("  6 = ▜ (reverse of 3)\n")
	legend.WriteString("  7 = ▙ (reverse of 2)\n")
	legend.WriteString("  8 = ▟ (reverse of 1)\n\n")
	legend.WriteString("Diagonals:\n")
	legend.WriteString("  \\ = ▚ (Backward)\n")
	legend.WriteString("  / = ▞ (Forward)\n\n")
	legend.WriteString("Space:\n")
	legend.WriteString("  _ = Space\n")
	return legend.String()
}
