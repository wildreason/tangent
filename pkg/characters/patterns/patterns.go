package patterns

import (
	"fmt"
	"math/rand"
)

// PatternCodes provides centralized pattern code definitions
type PatternCodes struct {
	// Basic blocks
	FullBlock  rune // F
	TopHalf    rune // T
	BottomHalf rune // B
	LeftHalf   rune // L
	RightHalf  rune // R

	// Shading blocks
	LightShade  rune // .
	MediumShade rune // :
	DarkShade   rune // #

	// Quadrants (1-4)
	Quad1 rune // 1
	Quad2 rune // 2
	Quad3 rune // 3
	Quad4 rune // 4

	// Three-quadrant composites (5-8)
	Quad5 rune // 5
	Quad6 rune // 6
	Quad7 rune // 7
	Quad8 rune // 8

	// Diagonals
	DiagonalBackward rune // \
	DiagonalForward  rune // /

	// Special
	Space rune // _

	// Dynamic noise - replaced at render time with random block
	Noise rune // $
}

// NoisePlaceholder is the Unicode character used as a placeholder for noise
// This gets replaced with a random block character at render time
const NoisePlaceholder = '◌' // U+25CC DOTTED CIRCLE

// NoisePool contains all block characters that can be randomly selected
// for the dynamic noise effect (Wall Street "rush" visual)
var NoisePool = []rune{
	// Quadrants - high visual activity
	'▘', '▝', '▖', '▗', // Single quadrants
	'▛', '▜', '▙', '▟', // Three-quadrant composites
	// Shading - medium visual activity
	'░', '▒', '▓',
	// Diagonals - directional energy
	'▚', '▞',
}

// DefaultPatternCodes returns the standard pattern code mapping
func DefaultPatternCodes() PatternCodes {
	return PatternCodes{
		// Basic blocks
		FullBlock:  '█', // F
		TopHalf:    '▀', // T
		BottomHalf: '▄', // B
		LeftHalf:   '▌', // L
		RightHalf:  '▐', // R

		// Shading blocks
		LightShade:  '░', // .
		MediumShade: '▒', // :
		DarkShade:   '▓', // #

		// Quadrants (1-4)
		Quad1: '▘', // 1
		Quad2: '▝', // 2
		Quad3: '▖', // 3
		Quad4: '▗', // 4

		// Three-quadrant composites (5-8)
		Quad5: '▛', // 5
		Quad6: '▜', // 6
		Quad7: '▙', // 7
		Quad8: '▟', // 8

		// Diagonals
		DiagonalBackward: '▚', // \
		DiagonalForward:  '▞', // /

		// Special
		Space: ' ', // _

		// Dynamic noise
		Noise: NoisePlaceholder, // $
	}
}

// GetPatternHelp returns a formatted help string for pattern codes
func GetPatternHelp() string {
	codes := DefaultPatternCodes()
	return fmt.Sprintf("Pattern codes: F=%c T=%c B=%c L=%c R=%c 1-8=quads .=%c :=%c #=%c _=%c",
		codes.FullBlock, codes.TopHalf, codes.BottomHalf, codes.LeftHalf, codes.RightHalf,
		codes.LightShade, codes.MediumShade, codes.DarkShade, codes.Space)
}

// GetPatternDescription returns a detailed description of pattern codes
func GetPatternDescription() string {
	codes := DefaultPatternCodes()
	return fmt.Sprintf(`Pattern codes:
  F=%c  T=%c  B=%c  L=%c  R=%c  (basic blocks)
  1-8=quads: %c%c%c%c %c%c%c%c
  .=%c :=%c #=%c (shades) _=%c $=%c (special)`,
		codes.FullBlock, codes.TopHalf, codes.BottomHalf, codes.LeftHalf, codes.RightHalf,
		codes.Quad1, codes.Quad2, codes.Quad3, codes.Quad4,
		codes.Quad5, codes.Quad6, codes.Quad7, codes.Quad8,
		codes.LightShade, codes.MediumShade, codes.DarkShade, codes.Space, codes.Noise)
}

// RandomNoise returns a random block character from the noise pool
func RandomNoise() rune {
	return NoisePool[rand.Intn(len(NoisePool))]
}

// ReplaceNoise replaces all noise placeholders in a string with random blocks
func ReplaceNoise(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r == NoisePlaceholder {
			runes[i] = RandomNoise()
		}
	}
	return string(runes)
}
