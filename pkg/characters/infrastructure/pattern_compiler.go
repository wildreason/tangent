package infrastructure

import (
	"fmt"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/patterns"
)

// SimplePatternCompiler implements the PatternCompiler interface
type SimplePatternCompiler struct {
	patterns   map[rune]rune
	validators []PatternValidator
}

// PatternValidator defines the interface for pattern validation
type PatternValidator interface {
	Validate(pattern string) error
}

// NewPatternCompiler creates a new pattern compiler
func NewPatternCompiler() domain.PatternCompiler {
	codes := patterns.DefaultPatternCodes()
	return &SimplePatternCompiler{
		patterns: map[rune]rune{
			// Basic blocks (uppercase)
			'F': codes.FullBlock,  // Full Block
			'T': codes.TopHalf,    // Top Half Block
			'B': codes.BottomHalf, // Bottom Half Block
			'L': codes.LeftHalf,   // Left Half Block
			'R': codes.RightHalf,  // Right Half Block

			// Basic blocks (lowercase)
			'f': codes.FullBlock,  // Full Block
			't': codes.TopHalf,    // Top Half Block
			'b': codes.BottomHalf, // Bottom Half Block
			'l': codes.LeftHalf,   // Left Half Block
			'r': codes.RightHalf,  // Right Half Block

			// Shading blocks
			'.': codes.LightShade,  // Light Shade
			':': codes.MediumShade, // Medium Shade
			'#': codes.DarkShade,   // Dark Shade

			// Single quadrants (1-4)
			'1': codes.Quad1, // Quadrant Upper Left
			'2': codes.Quad2, // Quadrant Upper Right
			'3': codes.Quad3, // Quadrant Lower Left
			'4': codes.Quad4, // Quadrant Lower Right

			// Three-quadrant composites (5-8, reverse of 1-4)
			'5': codes.Quad5, // Three Quadrants: UL+UR+LL
			'6': codes.Quad6, // Three Quadrants: UL+UR+LR
			'7': codes.Quad7, // Three Quadrants: UL+LL+LR
			'8': codes.Quad8, // Three Quadrants: UR+LL+LR

			// Diagonals
			'\\': codes.DiagonalBackward, // Diagonal Backward
			'/':  codes.DiagonalForward,  // Diagonal Forward

			// Special
			'_': codes.Space,  // Space
			'X': codes.Mirror, // Mirror
		},
		validators: []PatternValidator{
			&LengthValidator{},
			&CharacterValidator{},
		},
	}
}

// Compile compiles a pattern string to Unicode block elements
func (c *SimplePatternCompiler) Compile(pattern string) string {
	result := make([]rune, 0, len(pattern))
	for _, char := range pattern {
		if unicode, exists := c.patterns[char]; exists {
			result = append(result, unicode)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

// Validate validates a pattern string with detailed error reporting
func (c *SimplePatternCompiler) Validate(pattern string) error {
	if pattern == "" {
		return domain.ErrEmptyPattern
	}

	// Check pattern length
	if len(pattern) > 100 { // Reasonable limit
		return domain.NewValidationError("pattern", pattern, "pattern exceeds maximum length of 100 characters")
	}

	// Check for invalid characters
	for i, char := range pattern {
		if _, exists := c.patterns[char]; !exists && char != ' ' && char != '\t' {
			return domain.NewPatternCompilationError(pattern, i, fmt.Sprintf("invalid character '%c'", char), nil)
		}
	}

	// Run additional validators
	for _, validator := range c.validators {
		if err := validator.Validate(pattern); err != nil {
			return err
		}
	}
	return nil
}

// LengthValidator validates pattern length
type LengthValidator struct{}

func (v *LengthValidator) Validate(pattern string) error {
	if len(pattern) == 0 {
		return domain.ErrEmptyPattern
	}
	return nil
}

// CharacterValidator validates pattern characters
type CharacterValidator struct{}

func (v *CharacterValidator) Validate(pattern string) error {
	// For now, accept any characters
	// Could add specific character validation here
	return nil
}
