package characters

import (
	"fmt"
	"strings"
)

// PatternCompiler converts single-character hex-style codes to Unicode block elements
// Pattern format: "00R9FFF9L00" (no commas, like hex color codes)
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
			'F': '█', // Full block
			'T': '▀', // Top half (Upper)
			'B': '▄', // Bottom half (Lower)
			'L': '▌', // Left half
			'R': '▐', // Right half

			// Shading blocks
			'.': '░', // Light shade (like a light dot)
			':': '▒', // Medium shade (like colon density)
			'#': '▓', // Dark shade (like hash density)

			// Single quadrants
			'1': '▘', // Quadrant 1 (Upper Left)
			'2': '▝', // Quadrant 2 (Upper Right)
			'3': '▖', // Quadrant 3 (Lower Left)
			'4': '▗', // Quadrant 4 (Lower Right)

			// Diagonals
			'/':  '▚', // Diagonal forward slash pattern
			'\\': '▞', // Diagonal backslash pattern

			// Three-quadrant composites (look like the numbers)
			'7': '▛', // Three-quad: UL_UR_LL (looks like 7)
			'9': '▜', // Three-quad: UL_UR_LR (looks like 9)
			'6': '▙', // Three-quad: UL_LL_LR (looks like 6)
			'8': '▟', // Three-quad: UR_LL_LR (looks like 8)

			// Space (support both 0 and _)
			'0': ' ', // Space as zero (like hex colors)
			'_': ' ', // Space as underscore (for readability)
		},
	}
}

// CompilePattern converts a compact hex-style pattern to a Unicode string
// Example: "00R9FFF9L00" → "  ▐▜███▜▌  "
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
	legend.WriteString("Quadrants (Single):\n")
	legend.WriteString("  1 = ▘ (Upper Left)\n")
	legend.WriteString("  2 = ▝ (Upper Right)\n")
	legend.WriteString("  3 = ▖ (Lower Left)\n")
	legend.WriteString("  4 = ▗ (Lower Right)\n\n")
	legend.WriteString("Quadrants (Three):\n")
	legend.WriteString("  7 = ▛ (UL+UR+LL)\n")
	legend.WriteString("  9 = ▜ (UL+UR+LR)\n")
	legend.WriteString("  6 = ▙ (UL+LL+LR)\n")
	legend.WriteString("  8 = ▟ (UR+LL+LR)\n\n")
	legend.WriteString("Diagonals:\n")
	legend.WriteString("  / = ▚ (Forward slash)\n")
	legend.WriteString("  \\ = ▞ (Backslash)\n\n")
	legend.WriteString("Space:\n")
	legend.WriteString("  0 = Space (like hex colors)\n")
	legend.WriteString("  _ = Space (for readability)\n")
	return legend.String()
}
